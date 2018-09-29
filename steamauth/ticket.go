// Package steamauth wraps Steam's user authentication API.
//
// This package is available on both clients and servers.
//
// See the Steam User Authentication documentation for more details.
// <https://partner.steamgames.com/doc/features/auth>
package steamauth

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

func CreateTicket() (ticket []byte, cancel func()) {
	var handle internal.HAuthTicket
	var buffer [1024]byte
	var actualLength uint32

	if internal.IsGameServer {
		handle = internal.SteamAPI_ISteamGameServer_GetAuthSessionTicket(unsafe.Pointer(&buffer[0]), int32(len(buffer)), &actualLength)
	} else {
		handle = internal.SteamAPI_ISteamUser_GetAuthSessionTicket(unsafe.Pointer(&buffer[0]), int32(len(buffer)), &actualLength)
	}

	return buffer[:actualLength], func() {
		if internal.IsGameServer {
			internal.SteamAPI_ISteamGameServer_CancelAuthTicket(handle)
		} else {
			internal.SteamAPI_ISteamUser_CancelAuthTicket(handle)
		}
	}
}
