package steamauth

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

var sessionLock sync.Mutex
var sessions = make(map[steamworks.SteamID]*Session)

type Session struct {
	claimedID steamworks.SteamID
	ownerID   steamworks.SteamID
	status    SessionStatus
	change    chan SessionStatus
}

func (s *Session) ClaimedID() steamworks.SteamID {
	// immutable; no need to lock
	return s.claimedID
}

func (s *Session) OwnerID() steamworks.SteamID {
	sessionLock.Lock()
	id := s.ownerID
	sessionLock.Unlock()
	return id
}

func (s *Session) Status() SessionStatus {
	sessionLock.Lock()
	status := s.status
	sessionLock.Unlock()
	return status
}

func (s *Session) Change() <-chan SessionStatus {
	// immutable; no need to lock
	return s.change
}

var (
	ErrInvalidTicket    = errors.New("steamworks/steamauth: invalid ticket")
	ErrDuplicateRequest = errors.New("steamworks/steamauth: there is already an active session for this SteamID")
	ErrInvalidVersion   = errors.New("steamworks/steamauth: ticket is from an incompatible version of the Steam API")
	ErrGameMismatch     = errors.New("steamworks/steamauth: ticket is for a different game")
	ErrExpired          = errors.New("steamworks/steamauth: ticket is expired")
	ErrUnknown          = errors.New("steamworks/steamauth: unknown error")
)

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

	switch result {
	case internal.EBeginAuthSessionResult_OK:
		break
	case internal.EBeginAuthSessionResult_InvalidTicket:
		return nil, ErrInvalidTicket
	case internal.EBeginAuthSessionResult_DuplicateRequest:
		return sessions[claimedID], ErrDuplicateRequest
	case internal.EBeginAuthSessionResult_InvalidVersion:
		return nil, ErrInvalidVersion
	case internal.EBeginAuthSessionResult_GameMismatch:
		return nil, ErrGameMismatch
	case internal.EBeginAuthSessionResult_ExpiredTicket:
		return nil, ErrExpired
	default:
		return nil, ErrUnknown
	}

	sess := &Session{
		claimedID: claimedID,
		status:    StatusUnknown,
		change:    make(chan SessionStatus, 1),
	}

	sessions[claimedID] = sess

	return sess, nil
}

func (s *Session) End() {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	defer internal.Cleanup()()

	if internal.IsGameServer {
		internal.SteamAPI_ISteamGameServer_EndAuthSession(internal.SteamID(s.claimedID))
	} else {
		internal.SteamAPI_ISteamUser_EndAuthSession(internal.SteamID(s.claimedID))
	}

	delete(sessions, s.claimedID)

	s.status = StatusUnknown
	close(s.change)
}
