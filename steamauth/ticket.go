package steamauth

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// CreateTicket generates a sequence of bytes that verifies your identity and
// ownership of a game to another Steam user or server.
//
// The ticket can only be used once, and cancel should be called when the ticket
// is no longer in use - that is, when the session ends.
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
