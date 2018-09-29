//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type SessionStatus -trimprefix Status

package steamauth

import (
	"sync"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

type SessionStatus internal.EAuthSessionResponse

const (
	StatusUnknown                 SessionStatus = SessionStatus(0)
	StatusOK                      SessionStatus = SessionStatus(internal.EAuthSessionResponse_OK + 1)
	StatusUserNotConnectedtoSteam SessionStatus = SessionStatus(internal.EAuthSessionResponse_UserNotConnectedToSteam + 1)
	StatusNoLicenseOrExpired      SessionStatus = SessionStatus(internal.EAuthSessionResponse_NoLicenseOrExpired + 1)
	StatusVACBanned               SessionStatus = SessionStatus(internal.EAuthSessionResponse_VACBanned + 1)
	StatusLoggedInElsewhere       SessionStatus = SessionStatus(internal.EAuthSessionResponse_LoggedInElseWhere + 1)
	StatusVACCheckTimedOut        SessionStatus = SessionStatus(internal.EAuthSessionResponse_VACCheckTimedOut + 1)
	StatusCanceled                SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketCanceled + 1)
	StatusInvalidAlreadyUsed      SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketInvalidAlreadyUsed + 1)
	StatusInvalid                 SessionStatus = SessionStatus(internal.EAuthSessionResponse_AuthTicketInvalid + 1)
	StatusPublisherIssuedBan      SessionStatus = SessionStatus(internal.EAuthSessionResponse_PublisherIssuedBan + 1)
)

var initOnce sync.Once

func doInit() {
	internal.RegisterCallback_ValidateAuthTicketResponse(func(data *internal.ValidateAuthTicketResponse, _ bool) {
		claimedID := steamworks.SteamID(data.SteamID)
		ownerID := steamworks.SteamID(data.OwnerSteamID)
		status := SessionStatus(data.EAuthSessionResponse + 1)

		sessionLock.Lock()

		sess, ok := sessions[claimedID]
		if !ok {
			sessionLock.Unlock()
			return
		}

		sess.ownerID = ownerID
		sess.status = status
		select {
		case <-sess.change:
			sess.change <- status
		case sess.change <- status:
		}

		sessionLock.Unlock()
	}, 0)
}
