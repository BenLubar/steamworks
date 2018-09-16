#pragma once

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C"
{
#endif

// Go exports
extern void onCallback(int callback_id, void *data, size_t data_length, bool ioFailure, uint64_t apiCallID);
extern void discardCallback(int callback_id);

#ifndef STEAMAPI_WRAPPER

// Go imports

bool SteamAPI_Init();
void SteamAPI_Shutdown();
bool SteamAPI_RestartAppIfNecessary(uint32_t unOwnAppID);
void SteamAPI_ReleaseCurrentThreadMemory();

typedef uint64_t CSteamID;

//-----------------------------------------------------------------------------
// Purpose: Base values for callback identifiers, each callback must
//			have a unique ID.
//-----------------------------------------------------------------------------
enum { k_iSteamUserCallbacks = 100 };
enum { k_iSteamGameServerCallbacks = 200 };
enum { k_iSteamFriendsCallbacks = 300 };
enum { k_iSteamBillingCallbacks = 400 };
enum { k_iSteamMatchmakingCallbacks = 500 };
enum { k_iSteamContentServerCallbacks = 600 };
enum { k_iSteamUtilsCallbacks = 700 };
enum { k_iClientFriendsCallbacks = 800 };
enum { k_iClientUserCallbacks = 900 };
enum { k_iSteamAppsCallbacks = 1000 };
enum { k_iSteamUserStatsCallbacks = 1100 };
enum { k_iSteamNetworkingCallbacks = 1200 };
enum { k_iClientRemoteStorageCallbacks = 1300 };
enum { k_iClientDepotBuilderCallbacks = 1400 };
enum { k_iSteamGameServerItemsCallbacks = 1500 };
enum { k_iClientUtilsCallbacks = 1600 };
enum { k_iSteamGameCoordinatorCallbacks = 1700 };
enum { k_iSteamGameServerStatsCallbacks = 1800 };
enum { k_iSteam2AsyncCallbacks = 1900 };
enum { k_iSteamGameStatsCallbacks = 2000 };
enum { k_iClientHTTPCallbacks = 2100 };
enum { k_iClientScreenshotsCallbacks = 2200 };
enum { k_iSteamScreenshotsCallbacks = 2300 };
enum { k_iClientAudioCallbacks = 2400 };
enum { k_iClientUnifiedMessagesCallbacks = 2500 };
enum { k_iSteamStreamLauncherCallbacks = 2600 };
enum { k_iClientControllerCallbacks = 2700 };
enum { k_iSteamControllerCallbacks = 2800 };
enum { k_iClientParentalSettingsCallbacks = 2900 };
enum { k_iClientDeviceAuthCallbacks = 3000 };
enum { k_iClientNetworkDeviceManagerCallbacks = 3100 };
enum { k_iClientMusicCallbacks = 3200 };
enum { k_iClientRemoteClientManagerCallbacks = 3300 };
enum { k_iClientUGCCallbacks = 3400 };
enum { k_iSteamStreamClientCallbacks = 3500 };
enum { k_IClientProductBuilderCallbacks = 3600 };
enum { k_iClientShortcutsCallbacks = 3700 };
enum { k_iClientRemoteControlManagerCallbacks = 3800 };
enum { k_iSteamAppListCallbacks = 3900 };
enum { k_iSteamMusicCallbacks = 4000 };
enum { k_iSteamMusicRemoteCallbacks = 4100 };
enum { k_iClientVRCallbacks = 4200 };
enum { k_iClientGameNotificationCallbacks = 4300 }; 
enum { k_iSteamGameNotificationCallbacks = 4400 }; 
enum { k_iSteamHTMLSurfaceCallbacks = 4500 };
enum { k_iClientVideoCallbacks = 4600 };
enum { k_iClientInventoryCallbacks = 4700 };
enum { k_iClientBluetoothManagerCallbacks = 4800 };
enum { k_iClientSharedConnectionCallbacks = 4900 };
enum { k_ISteamParentalSettingsCallbacks = 5000 };
enum { k_iClientShaderCallbacks = 5100 };

#if defined(__linux__) || defined(__APPLE__)
#pragma pack( push, 4 )
#else
#pragma pack( push, 8 )
#endif
// callback notification - a user wants to talk to us over the P2P channel via the SendP2PPacket() API
// in response, a call to AcceptP2PPacketsFromUser() needs to be made, if you want to talk with them
typedef struct P2PSessionRequest_t
{ 
	CSteamID m_steamIDRemote;			// user who wants to talk to us
} P2PSessionRequest_t;
enum { P2PSessionRequest_iCallback = k_iSteamNetworkingCallbacks + 2 };

typedef struct P2PSessionState_t
{
	uint8_t m_bConnectionActive;		// true if we've got an active open connection
	uint8_t m_bConnecting;			// true if we're currently trying to establish a connection
	uint8_t m_eP2PSessionError;		// last error recorded (see enum above)
	uint8_t m_bUsingRelay;			// true if it's going through a relay server (TURN)
	int32_t m_nBytesQueuedForSend;
	int32_t m_nPacketsQueuedForSend;
	uint32_t m_nRemoteIP;				// potential IP:Port of remote host. Could be TURN server. 
	uint16_t m_nRemotePort;			// Only exists for compatibility with older authentication api's
} P2PSessionState_t;

// callback notification - packets can't get through to the specified user via the SendP2PPacket() API
// all packets queued packets unsent at this point will be dropped
// further attempts to send will retry making the connection (but will be dropped if we fail again)
typedef struct P2PSessionConnectFail_t
{ 
	CSteamID m_steamIDRemote;			// user we were sending packets to
	uint8_t m_eP2PSessionError;			// EP2PSessionError indicating why we're having trouble
} P2PSessionConnectFail_t;
enum { P2PSessionConnectFail_iCallback = k_iSteamNetworkingCallbacks + 3 };
#pragma pack( pop )

// SendP2PPacket() send types
// Typically k_EP2PSendUnreliable is what you want for UDP-like packets, k_EP2PSendReliable for TCP-like packets
typedef enum EP2PSend
{
	// Basic UDP send. Packets can't be bigger than 1200 bytes (your typical MTU size). Can be lost, or arrive out of order (rare).
	// The sending API does have some knowledge of the underlying connection, so if there is no NAT-traversal accomplished or
	// there is a recognized adjustment happening on the connection, the packet will be batched until the connection is open again.
	k_EP2PSendUnreliable = 0,

	// As above, but if the underlying p2p connection isn't yet established the packet will just be thrown away. Using this on the first
	// packet sent to a remote host almost guarantees the packet will be dropped.
	// This is only really useful for kinds of data that should never buffer up, i.e. voice payload packets
	k_EP2PSendUnreliableNoDelay = 1,

	// Reliable message send. Can send up to 1MB of data in a single message. 
	// Does fragmentation/re-assembly of messages under the hood, as well as a sliding window for efficient sends of large chunks of data. 
	k_EP2PSendReliable = 2,

	// As above, but applies the Nagle algorithm to the send - sends will accumulate 
	// until the current MTU size (typically ~1200 bytes, but can change) or ~200ms has passed (Nagle algorithm). 
	// Useful if you want to send a set of smaller messages but have the coalesced into a single packet
	// Since the reliable stream is all ordered, you can do several small message sends with k_EP2PSendReliableWithBuffering and then
	// do a normal k_EP2PSendReliable to force all the buffered data to be sent.
	k_EP2PSendReliableWithBuffering = 3,
} EP2PSend;

bool SteamAPI_ISteamNetworking_AcceptP2PSessionWithUser(intptr_t instancePtr, CSteamID remoteSteamID);
bool SteamAPI_ISteamNetworking_AllowP2PPacketRelay(intptr_t instancePtr, bool bAllow);
bool SteamAPI_ISteamNetworking_GetP2PSessionState(intptr_t instancePtr, CSteamID steamIDRemote, P2PSessionState_t * pConnectionState);
bool SteamAPI_ISteamNetworking_SendP2PPacket(intptr_t instancePtr, CSteamID steamIDRemote, const void * pubData, uint32_t cubData, EP2PSend eP2PSendType, int nChannel);
bool SteamAPI_ISteamNetworking_IsP2PPacketAvailable(intptr_t instancePtr, uint32_t * pcubMsgSize, int nChannel);
bool SteamAPI_ISteamNetworking_ReadP2PPacket(intptr_t instancePtr, void * pubDest, uint32_t cubDest, uint32_t * pcubMsgSize, CSteamID * psteamIDRemote, int nChannel);
bool SteamAPI_ISteamNetworking_CloseP2PSessionWithUser(intptr_t instancePtr, CSteamID steamIDRemote);
bool SteamAPI_ISteamNetworking_CloseP2PChannelWithUser(intptr_t instancePtr, CSteamID steamIDRemote, int nChannel);
CSteamID SteamAPI_ISteamUser_GetSteamID(intptr_t instancePtr);

#endif

// Wrapper exports

intptr_t SteamAPIWrap_SteamNetworking();
intptr_t SteamAPIWrap_SteamUser();
void SteamAPIWrap_RunCallbacks();
int SteamAPIWrap_RegisterCallback(size_t payload_size, int callback_id);
int SteamAPIWrap_UnregisterCallback(int callback_id);
int SteamAPIWrap_RegisterCallResult(size_t payload_size, int callback_id, uint64_t api_call_id);

#ifdef __cplusplus
}
#endif
