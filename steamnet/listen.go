// Package steamnet wraps Steam's peer-to-peer networking API.
//
// See the Steam Networking documentation for more details.
// <https://partner.steamgames.com/doc/features/multiplayer/networking>
package steamnet

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// Listen registers a function to handle connection requests.
//
// The function will be called from the callback thread, and will cause the connection to be accepted if it returns true.
//
// Multiple listeners may be registered simultaneously, and connections will be accepted if any listener returns true.
func Listen(accept func(steamworks.SteamID) bool) steamworks.Registration {
	return internal.RegisterCallback_P2PSessionRequest(func(data *internal.P2PSessionRequest) {
		id := data.SteamIDRemote
		if accept(steamworks.SteamID(id)) {
			internal.SteamAPI_ISteamNetworking_AcceptP2PSessionWithUser(id)
		}
	})
}

// SetAllowPacketRelay allows or disallows P2P connections to fall back to being relayed through the Steam servers if a direct connection or NAT-traversal cannot be established.
//
// This only applies to connections created after setting this value, or to existing connections that need to automatically reconnect after this value is set.
//
// P2P packet relay is allowed by default.
func SetAllowPacketRelay(allow bool) {
	internal.SteamAPI_ISteamNetworking_AllowP2PPacketRelay(allow)
}
