package steamnet

import (
	"encoding/binary"
	"net"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// SessionState is the current connection state to a specified user, returned by GetSessionState.
// This is the under-the-hood info about what's going on with a previous call to SendP2PPacket.
// This typically shouldn't be needed except for debugging purposes.
type SessionState struct {
	ConnectionActive     bool   // Do we have an active open connection with the user (true) or not (false)?
	Connecting           bool   // Are we currently trying to establish a connection with the user (true) or not (false)?
	LastError            error  // Last error recorded on the socket.
	UsingRelay           bool   // Is this connection going through a Steam relay server (true) or not (false)?
	BytesQueuedForSend   int    // The number of bytes queued up to be sent to the user.
	PacketsQueuedForSend int    // The number of packets queued up to be sent to the user.
	RemoteIP             net.IP // The IP of remote host if set. Could be a Steam relay server. This only exists for compatibility with older authentication api's.
	RemotePort           int    // The Port of remote host if set. Could be a Steam relay server. This only exists for compatibility with older authentication api's.
}

// GetSessionState returns a structure with details about the session like whether or not there is an active connection,
// the number of bytes queued on the connection, the last error code (if any), whether or not a relay server is being used,
// and the IP and Port of the remote user (if known).
//
// This should only be needed for debugging purposes.
//
// Returns nil if there was no open session with the specified user.
func GetSessionState(user steamworks.SteamID) *SessionState {
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
		ConnectionActive:     state.BConnectionActive != 0,
		Connecting:           state.BConnecting != 0,
		LastError:            toError(uint8(state.EP2PSessionError)),
		UsingRelay:           state.BUsingRelay != 0,
		BytesQueuedForSend:   int(state.NBytesQueuedForSend),
		PacketsQueuedForSend: int(state.NPacketsQueuedForSend),
		RemoteIP:             remoteIP,
		RemotePort:           int(state.NRemotePort),
	}
}
