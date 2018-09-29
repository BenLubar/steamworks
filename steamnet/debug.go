package steamnet

import (
	"encoding/binary"
	"net"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// SessionState is the current connection state to a specified user, returned
// by GetSessionState. This is the under-the-hood info about what's going on
// with a previous call to SendP2PPacket. This typically shouldn't be needed
// except for debugging purposes.
type SessionState struct {
	// LastError recorded on the socket.
	LastError error
	// RemoteIP of the other end of the connection (if set). Could be a Steam
	// relay server. This only exists for compatibility with older
	// authentication APIs.
	RemoteIP net.IP
	// RemotePort of the other end of the connection (if set). Could be a Steam
	// relay server. This only exists for compatibility with older
	// authentication APIs.
	RemotePort int
	// BytesQueuedForSend is the number of bytes queued up to be sent to the
	// user.
	BytesQueuedForSend int
	// PacketsQueuedForSend is the number of packets queued up to be sent to
	// the user.
	PacketsQueuedForSend int
	// ConnectionActive is true if there is an open connection with the user.
	ConnectionActive bool
	// Connecting is true if we are currently trying to establish a connection
	// with the user.
	Connecting bool
	// UsingRelay is true if the connection is currently routed through a Steam
	// relay server.
	UsingRelay bool
}

// GetSessionState returns a structure with details about the session like
// whether there is an active connection, the number of bytes queued on the
// connection, the last error code (if any), whether or not a relay server is
// being used, and the IP and Port of the remote user (if known).
//
// This should only be needed for debugging purposes.
//
// Returns nil if there was no open session with the specified user.
func GetSessionState(user steamworks.SteamID) *SessionState {
	defer internal.Cleanup()()

	var state internal.P2PSessionState
	if !internal.SteamAPI_ISteamNetworking_GetP2PSessionState(internal.SteamID(user), &state) {
		return nil
	}

	remoteIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(remoteIP, uint32(state.NRemoteIP))
	if remoteIP.IsUnspecified() {
		remoteIP = nil
	}

	return &SessionState{
		LastError:            toError(internal.EP2PSessionError(state.EP2PSessionError)),
		RemoteIP:             remoteIP,
		RemotePort:           int(state.NRemotePort),
		BytesQueuedForSend:   int(state.NBytesQueuedForSend),
		PacketsQueuedForSend: int(state.NPacketsQueuedForSend),
		ConnectionActive:     state.BConnectionActive != 0,
		Connecting:           state.BConnecting != 0,
		UsingRelay:           state.BUsingRelay != 0,
	}
}
