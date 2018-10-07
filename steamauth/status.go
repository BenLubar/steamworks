//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type SessionStatus -trimprefix Status

package steamauth

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// SessionStatus represents the current status of an authentication session.
//
// StatusUnknown means that Steam has not yet responded to the verification
// request. The session should continue, but the user's claimed identity should
// not be trusted.
//
// StatusOK means that Steam has verified that the user is who they say they are
// and that they own the game.
//
// All other statuses are errors and should result in the session being
// terminated in most use cases. Valve recommends that in this case, in P2P
// sessions, the user that failed to authenticate should be kicked, and if
// the other users in the lobby do not also kick that user, the lobby should
// be left.
type SessionStatus internal.EAuthSessionResponse

const (
	// StatusClosed means that the session has been closed. This status can
	// only be received from the Change channel as a result of the channel being
	// closed, which means the channel will always be "ready" once this is the
	// session's status.
	StatusClosed SessionStatus = SessionStatus(0)
	// StatusUnknown means that Steam has not yet responded to the verification
	// request. The session should continue, but the user's claimed identity
	// should not be trusted.
	StatusUnknown SessionStatus = SessionStatus(1)
	// StatusOK means Steam has verified the user is online, the ticket is
	// valid, and the ticket has not been reused.
	StatusOK SessionStatus = SessionStatus(internal.EAuthSessionResponse_OK + 2)
	// StatusUserNotConnectedToSteam means the user in question is not
	// connected to steam.
	StatusUserNotConnectedToSteam SessionStatus = SessionStatus(internal.EAuthSessionResponse_UserNotConnectedToSteam + 2)
	// StatusNoLicenseOrExpired means the user doesn't have a license for this
	// App ID or the ticket has expired.
	StatusNoLicenseOrExpired SessionStatus = SessionStatus(internal.EAuthSessionResponse_NoLicenseOrExpired + 2)
	// StatusVACBanned means the user is VAC banned for this game.
	StatusVACBanned SessionStatus = SessionStatus(internal.EAuthSessionResponse_VACBanned + 2)
	// StatusLoggedInElsewhere means the user account has logged in elsewhere
	// and the session containing the game instance has been disconnected.
	StatusLoggedInElsewhere SessionStatus = SessionStatus(internal.EAuthSessionResponse_LoggedInElseWhere + 2)
	// StatusVACCheckTimedOut means VAC has been unable to perform anti-cheat
	// checks on this user.
	StatusVACCheckTimedOut SessionStatus = SessionStatus(internal.EAuthSessionResponse_VACCheckTimedOut + 2)
	// StatusCanceled means the ticket has been canceled by the issuer.
	StatusCanceled SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketCanceled + 2)
	// StatusInvalidAlreadyUsed means this ticket has already been used. It is
	// not valid.
	StatusInvalidAlreadyUsed SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketInvalidAlreadyUsed + 2)
	// StatusInvalid means this ticket is not from a user instance currently
	// connected to steam.
	StatusInvalid SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketInvalid + 2)
	// StatusPublisherIssuedBan means the user is banned for this game. The ban
	// came via the web api and not VAC.
	StatusPublisherIssuedBan SessionStatus = SessionStatus(internal.EAuthSessionResponse_PublisherIssuedBan + 2)
)

var initOnce internal.Once

func doInit() {
	internal.RegisterCallback_ValidateAuthTicketResponse(func(data *internal.ValidateAuthTicketResponse, _ bool) {
		claimedID := steamworks.SteamID(data.SteamID.Get())
		ownerID := steamworks.SteamID(data.OwnerSteamID.Get())
		status := SessionStatus(data.EAuthSessionResponse + 2)

		sessionLock.Lock()
		defer sessionLock.Unlock()

		sess := sessions[claimedID]
		if sess == nil {
			return
		}

		sess.ownerID = ownerID
		sess.status = status
		select {
		case <-sess.change:
			sess.change <- status
		case sess.change <- status:
		}
	}, 0)
}
