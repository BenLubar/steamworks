package steamnet

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// CloseChannel closes a P2P channel when you're done talking to a user on the
// specific channel.
//
// Once all channels to a user have been closed, the open session to the user
// will be closed and new data from this user will trigger a new Listen
// callback.
//
// Returns true if the channel was successfully closed; otherwise, false if
// there was no active session or channel with the user.
func CloseChannel(user steamworks.SteamID, channel int32) bool {
	defer internal.Cleanup()()

	return internal.SteamAPI_ISteamNetworking_CloseP2PChannelWithUser(internal.SteamID(user), channel)
}

// CloseAllChannels should be called when you're done communicating with a user,
// as this will free up all of the resources allocated for the connection
// under-the-hood.
//
// If the remote user tries to send data to you again, a new Listen callback
// will be posted.
//
// Returns true if the session was successfully closed; otherwise, false if no
// connection was open with the user.
func CloseAllChannels(user steamworks.SteamID) bool {
	defer internal.Cleanup()()

	return internal.SteamAPI_ISteamNetworking_CloseP2PSessionWithUser(internal.SteamID(user))
}
