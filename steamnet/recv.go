package steamnet

import (
	"sync"
	"unsafe"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

var packetLock sync.Mutex

// ReadPacket checks if a P2P packet is available and returns the packet if
// there is one.
//
// This should be called in a loop for each channel that you use.
//
// This call is non-blocking. It will return (nil, 0) if no data is available.
func ReadPacket(channel int32) ([]byte, steamworks.SteamID) {
	defer internal.Cleanup()()

	// Although the call is non-blocking, we need to call two functions, and
	// we don't want the state of the connection to be changed by another
	// caller to ReadPacket in-between.
	packetLock.Lock()
	defer packetLock.Unlock()

	var size uint32
	if !internal.SteamAPI_ISteamNetworking_IsP2PPacketAvailable(&size, channel) {
		return nil, 0
	}

	buffer := make([]byte, size)
	var steamID internal.SteamID
	if !internal.SteamAPI_ISteamNetworking_ReadP2PPacket(unsafe.Pointer(&buffer[0]), size, &size, &steamID, channel) {
		panic("steamnet: packet was not actually available")
	}
	if int(size) != len(buffer) {
		panic("steamnet: packet size mismatch")
	}
	return buffer, steamworks.SteamID(steamID)
}
