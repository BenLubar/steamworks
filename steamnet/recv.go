package steamnet

import (
	"sync"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

var packetLock sync.Mutex

// ReadPacket checks if a P2P packet is available and returns the packet if there is one.
//
// This should be called in a loop for each channel that you use.
//
// This call is non-blocking. It will return (nil, 0) if no data is available.
func ReadPacket(channel int) ([]byte, steamworks.SteamID) {
	packetLock.Lock()
	defer packetLock.Unlock()

	var size uint32
	if !internal.ISteamNetworking_IsP2PPacketAvailable(&size, channel) {
		return nil, 0
	}

	buffer := make([]byte, size)
	var steamID steamworks.SteamID
	if !internal.ISteamNetworking_ReadP2PPacket(buffer, &size, (*internal.SteamID)(&steamID), channel) {
		panic("steamnet: packet was not actually available")
	}
	if int(size) != len(buffer) {
		panic("steamnet: packet size mismatch")
	}
	return buffer, steamID
}
