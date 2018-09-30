// Package steamauth wraps Steam's user authentication API.
//
// This package is available on both clients and servers.
//
// See the Steam User Authentication documentation for more details.
// <https://partner.steamgames.com/doc/features/auth>
package steamauth

import (
	"errors"
	"os"
	"runtime"
	"sync"
	"unsafe"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// All session mutations are protected by this global lock.
var sessionLock sync.Mutex
var sessions = make(map[steamworks.SteamID]*sessionData)

// Session represents a Steam authentication session. All methods on Session
// are safe to call concurrently.
//
// Sessions must be closed when they are no longer in use. Failing to do so
// will result in a message being written to the standard error stream.
type Session struct {
	claimedID steamworks.SteamID
	data      *sessionData
	closed    bool
}

// sessionData is separate from Session to allow Session to be garbage
// collected. The reference counter is used to know when a session has
// been leaked.
type sessionData struct {
	ownerID steamworks.SteamID
	status  SessionStatus
	change  chan SessionStatus
	refs    uintptr
}

// ClaimedID returns the SteamID of the remote user for this session.
//
// This value should only be trusted if Status is StatusOK.
func (s *Session) ClaimedID() steamworks.SteamID {
	// immutable; no need to lock
	return s.claimedID
}

// OwnerID returns the SteamID of the user who owns this game.
//
// This may be different than ClaimedID if the game is shared. This function's
// return value may be zero or invalid if Status was never StatusOK.
func (s *Session) OwnerID() steamworks.SteamID {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	return s.data.ownerID
}

// Status returns the current status of this session.
//
// See the documentation for SessionStatus for information on what the various
// possible values mean.
func (s *Session) Status() SessionStatus {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	return s.data.status
}

// Change returns a channel which receives status changes. The channel is closed
// by Close, so closing a session will cause this channel to repeatedly receive
// StatusClosed.
func (s *Session) Change() <-chan SessionStatus {
	// immutable; no need to lock
	return s.data.change
}

// OwnsDLC returns true if the user owns the specified DLC, or false if the user
// does not own the DLC or if the session is not authenticated.
func (s *Session) OwnsDLC(dlc steamworks.AppID) bool {
	defer internal.Cleanup()()

	var result internal.EUserHasLicenseForAppResult

	if internal.IsGameServer {
		result = internal.SteamAPI_ISteamGameServer_UserHasLicenseForApp(internal.SteamID(s.claimedID), internal.AppId(dlc))
	} else {
		result = internal.SteamAPI_ISteamUser_UserHasLicenseForApp(internal.SteamID(s.claimedID), internal.AppId(dlc))
	}

	return result == internal.EUserHasLicenseForAppResult_EUserHasLicenseResultHasLicense
}

// Errors that can be returned from BeginSession.
var (
	ErrInvalidTicket    = errors.New("steamworks/steamauth: invalid ticket")
	ErrDuplicateRequest = errors.New("steamworks/steamauth: there is already an active session for this SteamID")
	ErrInvalidVersion   = errors.New("steamworks/steamauth: ticket is from an incompatible version of the Steam API")
	ErrGameMismatch     = errors.New("steamworks/steamauth: ticket is for a different game")
	ErrExpired          = errors.New("steamworks/steamauth: ticket is expired")
	ErrUnknown          = errors.New("steamworks/steamauth: unknown error")
)

// BeginSession verifies a ticket returned by CreateTicket from another Steam
// user. The returned session must be closed when it is no longer in use.
//
// Sessions start with StatusUnknown. In this state, the session should continue
// but the user's claimed identity should not be trusted. If a session
// transitions to the StatusOK state, the claimed identity can be trusted.
// If the state is any other value, the session is invalid.
//
// When a session is closed, it may transition to the StatusCanceled state
// before transitioning to the StatusClosed state if the remote user closes the
// session first. If the user is already disconnecting, this can be ignored
// as both sides of the connection have already agreed that the session is
// to be closed.
//
// If the error returned by this function is ErrDuplicateRequest, the session
// may or may not be nil. If the error is nil, the session will not be nil.
// In any other case, the session is nil.
func BeginSession(ticket []byte, claimedID steamworks.SteamID) (*Session, error) {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	defer internal.Cleanup()()

	initOnce.Do(doInit)

	var result internal.EBeginAuthSessionResult

	if internal.IsGameServer {
		result = internal.SteamAPI_ISteamGameServer_BeginAuthSession(unsafe.Pointer(&ticket[0]), int32(len(ticket)), internal.SteamID(claimedID))
	} else {
		result = internal.SteamAPI_ISteamGameServer_BeginAuthSession(unsafe.Pointer(&ticket[0]), int32(len(ticket)), internal.SteamID(claimedID))
	}

	var sdata *sessionData
	var err error

	switch result {
	case internal.EBeginAuthSessionResult_OK:
		sdata = &sessionData{
			status: StatusUnknown,
			change: make(chan SessionStatus, 1),
		}
		sessions[claimedID] = sdata
	case internal.EBeginAuthSessionResult_InvalidTicket:
		err = ErrInvalidTicket
	case internal.EBeginAuthSessionResult_DuplicateRequest:
		sdata = sessions[claimedID]
		if sdata != nil {
			sdata.refs++
		}
		err = ErrDuplicateRequest
	case internal.EBeginAuthSessionResult_InvalidVersion:
		err = ErrInvalidVersion
	case internal.EBeginAuthSessionResult_GameMismatch:
		err = ErrGameMismatch
	case internal.EBeginAuthSessionResult_ExpiredTicket:
		err = ErrExpired
	default:
		err = ErrUnknown
	}

	var sess *Session
	if sdata != nil {
		sess = &Session{
			claimedID: claimedID,
			data:      sdata,
		}

		runtime.SetFinalizer(sess, (*Session).complain)
	}

	return sess, err
}

// ErrSessionAlreadyClosed is returned by Session.Close if the session is
// closed multiple times.
var ErrSessionAlreadyClosed = errors.New("steamworks/steamauth: session was already closed")

// Close closes a Session. If the session has already been closed,
// ErrSessionAlreadyClosed is returned and no other action is taken.
//
// Close must be called when the session ends.
func (s *Session) Close() error {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	if s.data != sessions[s.claimedID] || s.closed {
		// This session was already closed.
		return ErrSessionAlreadyClosed
	}

	s.close()
	return nil
}

func (s *Session) close() {
	defer internal.Cleanup()()

	if internal.IsGameServer {
		internal.SteamAPI_ISteamGameServer_EndAuthSession(internal.SteamID(s.claimedID))
	} else {
		internal.SteamAPI_ISteamUser_EndAuthSession(internal.SteamID(s.claimedID))
	}

	delete(sessions, s.claimedID)
	runtime.SetFinalizer(s, nil)

	s.data.status = StatusClosed
	close(s.data.change)

	s.closed = true
}

func (s *Session) complain() {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	if s.data != sessions[s.claimedID] || s.closed {
		// This session was already closed.
		return
	}

	// There is at least one other reference to this session still alive.
	if s.data.refs != 0 {
		s.data.refs--
		s.closed = true
		return
	}

	s.close()
	// Don't handle an error writing to Stderr because there's nothing we
	// can do about it.

	// nolint: gosec
	_, _ = os.Stderr.WriteString("[DEVELOPER ERROR] steamworks/steamauth: Sessions must be closed when they are no longer in use!\n")
}
