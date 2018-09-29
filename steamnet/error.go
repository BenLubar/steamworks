package steamnet

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// RegisterErrorCallback registers a function to be called when packets can't get through to the specified user.
//
// All queued packets unsent at this point will be dropped, further attempts to send will retry making the connection (but will be dropped if we fail again).
func RegisterErrorCallback(f func(steamworks.SteamID, error)) steamworks.Registration {
	return internal.RegisterCallback_P2PSessionConnectFail(func(data *internal.P2PSessionConnectFail, _ bool) {
		f(steamworks.SteamID(data.SteamIDRemote.Get()), toError(internal.EP2PSessionError(data.EP2PSessionError)))
	}, 0)
}

// Error represents a connection error in the Steam P2P API.
type Error internal.EP2PSessionError

var (
	// ErrNotRunningApp is returned if the target user is not running the same game.
	ErrNotRunningApp error = Error(internal.EP2PSessionError_NotRunningApp)
	// ErrNoRightsToApp is returned if the local user doesn't own the app that is running.
	ErrNoRightsToApp error = Error(internal.EP2PSessionError_NoRightsToApp)
	// ErrDestinationNotLoggedIn is returned if the target user isn't connected to Steam.
	ErrDestinationNotLoggedIn error = Error(internal.EP2PSessionError_DestinationNotLoggedIn)
	// ErrTimeout is returned if the connection timed out because the target user didn't respond.
	ErrTimeout error = Error(internal.EP2PSessionError_Timeout)
)

func (err Error) Error() string {
	if err == 0 || int(err) >= len(p2pErrors) {
		return "steamnet: unknown error"
	}

	return p2pErrors[err]
}

// Timeout returns true iff this error was caused by a timeout.
func (err Error) Timeout() bool {
	return err == ErrTimeout
}

// Temporary returns true iff this error might go away after a retry with no other local actions.
func (err Error) Temporary() bool {
	return err == ErrNotRunningApp || err == ErrDestinationNotLoggedIn || err == ErrTimeout
}

var p2pErrors = [...]string{
	internal.EP2PSessionError_None:                   "",
	internal.EP2PSessionError_NotRunningApp:          "steamnet: the target user is not running the same game",
	internal.EP2PSessionError_NoRightsToApp:          "steamnet: the local user doesn't own the app that is running",
	internal.EP2PSessionError_DestinationNotLoggedIn: "steamnet: the target user isn't connected to Steam",
	internal.EP2PSessionError_Timeout:                "steamnet: the connection timed out because the target user didn't respond",
}

func toError(code internal.EP2PSessionError) error {
	if code == internal.EP2PSessionError_None {
		return nil
	}

	return Error(code)
}
