package steamnet

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// RegisterErrorCallback registers a function to be called when packets can't get through to the specified user.
//
// All queued packets unsent at this point will be dropped, further attempts to send will retry making the connection (but will be dropped if we fail again).
func RegisterErrorCallback(f func(steamworks.SteamID, error)) steamworks.Registration {
	return internal.RegisterCallback_P2PSessionConnectFail(func(data *internal.P2PSessionConnectFail) {
		f(steamworks.SteamID(data.SteamIDRemote), toError(uint8(data.EP2PSessionError)))
	})
}

// Error represents a connection error in the Steam P2P API.
type Error uint8

var (
	ErrNotRunningApp          error = Error(1)
	ErrNoRightsToApp          error = Error(2)
	ErrDestinationNotLoggedIn error = Error(3)
	ErrTimeout                error = Error(4)
)

func (err Error) Error() string {
	if err == 0 || int(err) >= len(p2pErrors) {
		return "steamnet: unknown error"
	}

	return p2pErrors[err]
}

func (err Error) Timeout() bool {
	return err == ErrTimeout
}

func (err Error) Temporary() bool {
	return err == ErrTimeout
}

var p2pErrors = [...]string{
	0: "",                                                                          // k_EP2PSessionErrorNone
	1: "steamnet: the target user is not running the same game",                    // k_EP2PSessionErrorNotRunningApp
	2: "steamnet: the local user doesn't own the app that is running",              // k_EP2PSessionErrorNoRightsToApp
	3: "steamnet: the target user isn't connected to Steam",                        // k_EP2PSessionErrorDestinationNotLoggedIn
	4: "steamnet: the connection timed out because the target user didn't respond", // k_EP2PSessionErrorTimeout
}

func toError(code uint8) error {
	if code == 0 {
		return nil
	}

	return Error(code)
}
