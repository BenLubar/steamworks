package internal

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -isystem ${SRCDIR}/include
// #cgo windows LDFLAGS: -L ${SRCDIR}/lib/windows
// #cgo linux,386 LDFLAGS: -L ${SRCDIR}/lib/linux32
// #cgo linux,amd64 LDFLAGS: -L ${SRCDIR}/lib/linux64
// #cgo linux windows,386 darwin LDFLAGS: -lsteam_api -lsdkencryptedappticket
// #cgo windows,amd64 LDFLAGS: -lsteam_api64 -lsdkencryptedappticket64
//
// #include "steamworks_wrap.h"
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

type SteamID = C.CSteamID

var callbackLock sync.Mutex
var registeredCallbacks = make(map[RegisteredCallback]func(unsafe.Pointer, uintptr, bool, uint64))

//export onCallback
func onCallback(callback_id C.int, data unsafe.Pointer, data_length C.size_t, ioFailure C.bool, apiCallID C.uint64_t) {
	callbackLock.Lock()
	cb, ok := registeredCallbacks[RegisteredCallback(callback_id)]
	callbackLock.Unlock()

	if ok {
		cb(data, uintptr(data_length), bool(ioFailure), uint64(apiCallID))
	}
}

//export discardCallback
func discardCallback(callback_id C.int) {
	callbackLock.Lock()
	delete(registeredCallbacks, RegisteredCallback(callback_id))
	callbackLock.Unlock()
}

func threadCleanup() func() {
	runtime.LockOSThread()

	return func() {
		SteamAPI_ReleaseCurrentThreadMemory()
		runtime.UnlockOSThread()
	}
}

func registerCallback(payload_size uintptr, callback_type_id C.int, f func(unsafe.Pointer)) RegisteredCallback {
	wrap := func(data unsafe.Pointer, data_length uintptr, _ bool, _ uint64) {
		if data_length != payload_size {
			panic("steamworks/internal: payload size mismatch")
		}

		f(data)
	}

	defer threadCleanup()()

	callback_id := RegisteredCallback(C.SteamAPIWrap_RegisterCallback(C.size_t(payload_size), callback_type_id))
	callbackLock.Lock()
	registeredCallbacks[callback_id] = wrap
	callbackLock.Unlock()
	return callback_id
}

type RegisteredCallback C.int

func (rc RegisteredCallback) Unregister() {
	defer threadCleanup()()

	C.SteamAPIWrap_UnregisterCallback(C.int(rc))
}

// SteamAPI_Init initializes the Steamworks API.
//
// See Initialization and Shutdown for additional information.
// <https://partner.steamgames.com/doc/sdk/api#initialization_and_shutdown>
//
// Returns true if all required interfaces have been acquired and are accessible.
//
// false indicates one of the following conditions:
// - The Steam client isn't running. A running Steam client is required to provide implementations of the various Steamworks interfaces.
// - The Steam client couldn't determine the App ID of game. If you're running your application from the executable or debugger directly
//   then you must have a steam_appid.txt in your game directory next to the executable, with your app ID in it and nothing else. Steam will
//   look for this file in the current working directory. If you are running your executable from a different directory you may need to relocate
//   the steam_appid.txt file.
// - Your application is not running under the same OS user context as the Steam client, such as a different user or administration access level.
// - Ensure that you own a license for the App ID on the currently active Steam account. Your game must show up in your Steam library.
// - Your App ID is not completely set up, i.e. in Release State: Unavailable, or it's missing default packages.
func SteamAPI_Init() bool {
	return bool(C.SteamAPI_Init())
}

// SteamAPI_Shutdown shuts down the Steamworks API, releases pointers and frees memory.
//
// You should call this during process shutdown if possible.
//
// This will not unhook the Steam Overlay from your game as there's no guarantee that your rendering API is done using it.
func SteamAPI_Shutdown() {
	C.SteamAPI_Shutdown()
}

// SteamAPI_RestartAppIfNecessary checks if your executable was launched through Steam and relaunches it through Steam if it wasn't.
//
// See Initialization and Shutdown for additional information.
// <https://partner.steamgames.com/doc/sdk/api#initialization_and_shutdown>
//
// If this returns true then it starts the Steam client if required and launches your game again through it, and you should quit your
// process as soon as possible. This effectively runs `steam://run/<AppId>`` so it may not relaunch the exact executable that called it,
// as it will always relaunch from the version installed in your Steam library folder.
//
// If it returns false, then your game was launched by the Steam client and no action needs to be taken. One exception is if a steam_appid.txt
// file is present then this will return false regardless. This allows you to develop and test without launching your game through the Steam client.
// Make sure to remove the steam_appid.txt file when uploading the game to your Steam depot!
func SteamAPI_RestartAppIfNecessary(ownAppID uint32) bool {
	return bool(C.SteamAPI_RestartAppIfNecessary(C.uint32_t(ownAppID)))
}

// SteamAPI_ReleaseCurrentThreadMemory free the internal Steamworks API memory associated with the calling thread.
//
// Most Steamworks API functions allocate a small amount of thread-local memory for parameter storage, calling this
// will manually free such memory. This function is called automatically by SteamAPI_RunCallbacks, so a program that
// only ever accesses the Steamworks API from a single thread never needs to explicitly call this function.
func SteamAPI_ReleaseCurrentThreadMemory() {
	C.SteamAPI_ReleaseCurrentThreadMemory()
}

// SteamAPI_RunCallbacks dispatches callbacks and call results to all of the registered listeners.
//
// It's best to call this at >10Hz, the more time between calls, the more potential latency between receiving events
// or results from the Steamworks API. Most games call this once per render-frame. All registered listener functions
// will be invoked during this call, in the callers thread context.
//
// SteamAPI_RunCallbacks is safe to call from multiple threads simultaneously, but if you choose to do this, callback
// code could be executed on any thread. One alternative is to call SteamAPI_RunCallbacks from the main thread only,
// and call SteamAPI_ReleaseCurrentThreadMemory regularly on other threads.
func SteamAPI_RunCallbacks() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	C.SteamAPIWrap_RunCallbacks()
}

// ISteamNetworking_AcceptP2PSessionWithUser allows the game to specify accept an incoming packet.
// This needs to be called before a real connection is established to a remote host, the game will get a chance to say whether or not the remote user is allowed to talk to them.
//
// When a remote user that you haven't sent a packet to recently, tries to first send you a packet, your game will receive a callback P2PSessionRequest_t.
// This callback contains the Steam ID of the user who wants to send you a packet.
// In response to this callback, you'll want to see if it's someone you want to talk to (for example, if they're in a lobby with you), and if so, accept the connection;
// otherwise if you don't want to talk to the user, just ignore the request. If the user continues to send you packets, another P2PSessionRequest_t will be posted periodically.
// If you've called SendP2PPacket on the other user, this implicitly accepts the session request.
//
// Note that this call should only be made in response to a P2PSessionRequest_t callback!
//
// Returns true upon success; false only if steamIDRemote is invalid.
func ISteamNetworking_AcceptP2PSessionWithUser(steamIDRemote SteamID) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_AcceptP2PSessionWithUser(C.SteamAPIWrap_SteamNetworking(), steamIDRemote))
}

// Allow or disallow P2P connections to fall back to being relayed through the Steam servers if a direct connection or NAT-traversal cannot be established.
//
// This only applies to connections created after setting this value, or to existing connections that need to automatically reconnect after this value is set.
//
// P2P packet relay is allowed by default.
//
// This function always returns true.
func ISteamNetworking_AllowP2PPacketRelay(allow bool) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_AllowP2PPacketRelay(C.SteamAPIWrap_SteamNetworking(), C.bool(allow)))
}

// Connection state to a specified user, returned by GetP2PSessionState.
// This is the under-the-hood info about what's going on with a previous call to SendP2PPacket.
// This typically shouldn't be needed except for debugging purposes.
type P2PSessionState struct {
	ConnectionActive     uint8  //Do we have an active open connection with the user (true) or not (false)?
	Connecting           uint8  //Are we currently trying to establish a connection with the user (true) or not (false)?
	EP2PSessionError     uint8  //Last error recorded on the socket. This returns a EP2PSessionError.
	UsingRelay           uint8  //Is this connection going through a Steam relay server (true) or not (false)?
	BytesQueuedForSend   int32  //The number of bytes queued up to be sent to the user.
	PacketsQueuedForSend int32  //The number of packets queued up to be sent to the user.
	RemoteIP             uint32 //The IP of remote host if set. Could be a Steam relay server. This only exists for compatibility with older authentication api's.
	RemotePort           uint16 //The Port of remote host if set. Could be a Steam relay server. This only exists for compatibility with older authentication api's.
}

// ISteamNetworking_GetP2PSessionState fills out a P2PSessionState structure with details about the connection like whether or not there is an active connection;
// number of bytes queued on the connection; the last error code, if any; whether or not a relay server is being used; and the IP and Port of the remote user, if known
//
// This should only needed for debugging purposes.
//
// Returns true if *connectionState was filled out; otherwise, false if there was no open session with the specified user.
func ISteamNetworking_GetP2PSessionState(steamIDRemote SteamID, connectionState *P2PSessionState) bool {
	defer threadCleanup()()

	var cConnectionState C.P2PSessionState_t
	if !C.SteamAPI_ISteamNetworking_GetP2PSessionState(C.SteamAPIWrap_SteamNetworking(), steamIDRemote, &cConnectionState) {
		return false
	}

	connectionState.ConnectionActive = uint8(cConnectionState.m_bConnectionActive)
	connectionState.Connecting = uint8(cConnectionState.m_bConnecting)
	connectionState.EP2PSessionError = uint8(cConnectionState.m_eP2PSessionError)
	connectionState.UsingRelay = uint8(cConnectionState.m_bUsingRelay)
	connectionState.BytesQueuedForSend = int32(cConnectionState.m_nBytesQueuedForSend)
	connectionState.PacketsQueuedForSend = int32(cConnectionState.m_nPacketsQueuedForSend)
	connectionState.RemoteIP = uint32(cConnectionState.m_nRemoteIP)
	connectionState.RemotePort = uint16(cConnectionState.m_nRemotePort)
	return true
}

func ISteamNetworking_IsP2PPacketAvailable(size *uint32, channel int) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_IsP2PPacketAvailable(C.SteamAPIWrap_SteamNetworking(), (*C.uint32_t)(size), C.int(channel)))
}

func ISteamNetworking_ReadP2PPacket(buffer []byte, size *uint32, steamID *SteamID, channel int) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_ReadP2PPacket(C.SteamAPIWrap_SteamNetworking(), unsafe.Pointer(&buffer[0]), C.uint32_t(len(buffer)), (*C.uint32_t)(size), steamID, C.int(channel)))
}

func ISteamNetworking_SendP2PPacket(steamIDRemote SteamID, data []byte, sendType, channel int) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_SendP2PPacket(C.SteamAPIWrap_SteamNetworking(), steamIDRemote, (unsafe.Pointer)(&data[0]), C.uint32_t(len(data)), C.EP2PSend(sendType), C.int(channel)))
}

func ISteamNetworking_CloseP2PChannelWithUser(steamIDRemote SteamID, channel int) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_CloseP2PChannelWithUser(C.SteamAPIWrap_SteamNetworking(), steamIDRemote, C.int(channel)))
}

func ISteamNetworking_CloseP2PSessionWithUser(steamIDRemote SteamID) bool {
	defer threadCleanup()()

	return bool(C.SteamAPI_ISteamNetworking_CloseP2PSessionWithUser(C.SteamAPIWrap_SteamNetworking(), steamIDRemote))
}

// A user wants to communicate with us over the P2P channel via the SendP2PPacket.
// In response, a call to AcceptP2PSessionWithUser needs to be made, if you want to open the network channel with them.
type P2PSessionRequest struct {
	SteamIDRemote SteamID // The user who wants to start a P2P session with us.
}

func RegisterCallback_P2PSessionRequest(f func(P2PSessionRequest)) RegisteredCallback {
	return registerCallback(unsafe.Sizeof(C.P2PSessionRequest_t{}), C.P2PSessionRequest_iCallback, func(pdata unsafe.Pointer) {
		data := (*C.P2PSessionRequest_t)(pdata)
		f(P2PSessionRequest{
			SteamIDRemote: SteamID(data.m_steamIDRemote),
		})
	})
}

// Called when packets can't get through to the specified user.
// All queued packets unsent at this point will be dropped, further attempts to send will retry making the connection (but will be dropped if we fail again).
type P2PSessionConnectFail struct {
	SteamIDRemote    SteamID // User we were trying to send the packets to.
	EP2PSessionError uint8   //Indicates the reason why we're having trouble. Actually a EP2PSessionError.
}

func RegisterCallback_P2PSessionConnectFail(f func(P2PSessionConnectFail)) RegisteredCallback {
	return registerCallback(unsafe.Sizeof(C.P2PSessionConnectFail_t{}), C.P2PSessionConnectFail_iCallback, func(pdata unsafe.Pointer) {
		data := (*C.P2PSessionConnectFail_t)(pdata)
		f(P2PSessionConnectFail{
			SteamIDRemote:    SteamID(data.m_steamIDRemote),
			EP2PSessionError: uint8(data.m_eP2PSessionError),
		})
	})
}

// ISteamUser_GetSteamID gets the Steam ID of the account currently logged into the Steam client. This is commonly called the 'current user', or 'local user'.
//
// A Steam ID is a unique identifier for a Steam accounts, Steam groups, Lobbies and Chat rooms, and used to differentiate users in all parts of the Steamworks API.
func ISteamUser_GetSteamID() SteamID {
	defer threadCleanup()()

	return SteamID(C.SteamAPI_ISteamUser_GetSteamID(C.SteamAPIWrap_SteamUser()))
}
