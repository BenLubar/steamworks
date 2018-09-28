#include "shim.h"
#include <steam/steam_gameserver.h>

extern "C"
{

#include "gameserver.h"

//-----------------------------------------------------------------------------
// Purpose: Functions for authenticating users via Steam to play on a game server
//-----------------------------------------------------------------------------

//
// Basic server data.  These properties, if set, must be set before before calling LogOn.  They
// may not be changed after logged in.
//

// This is called by SteamGameServer_Init, and you will usually not need to call it directly
bool SteamAPI_SteamGameServer_InitGameServer(intp instance, uint32 unIP, uint16 usGamePort, uint16 usQueryPort, uint32 unFlags, AppId_t nGameAppId, const char *pchVersionString)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->InitGameServer(unIP, usGamePort, usQueryPort, unFlags, nGameAppId, pchVersionString);
}

// Game product identifier.  This is currently used by the master server for version checking purposes.
// It's a required field, but will eventually will go away, and the AppID will be used for this purpose.
void SteamAPI_SteamGameServer_SetProduct(intp instance, const char *pszProduct)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetProduct(pszProduct);
}

// Description of the game.  This is a required field and is displayed in the steam server browser....for now.
// This is a required field, but it will go away eventually, as the data should be determined from the AppID.
void SteamAPI_SteamGameServer_SetGameDescription(intp instance, const char *pszGameDescription)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetGameDescription(pszGameDescription);
}

// If your game is a "mod," pass the string that identifies it.  The default is an empty string, meaning
// this application is the original game, not a mod.
//
// @see k_cbMaxGameServerGameDir
void SteamAPI_SteamGameServer_SetModDir(intp instance, const char *pszModDir)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetModDir(pszModDir);
}

// Is this is a dedicated server?  The default value is false.
void SteamAPI_SteamGameServer_SetDedicatedServer(intp instance, bool bDedicated)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetDedicatedServer(bDedicated);
}

//
// Login
//

// Begin process to login to a persistent game server account
//
// You need to register for callbacks to determine the result of this operation.
// @see SteamServersConnected_t
// @see SteamServerConnectFailure_t
// @see SteamServersDisconnected_t
void SteamAPI_SteamGameServer_LogOn(intp instance, const char *pszToken)
{
    reinterpret_cast<ISteamGameServer *>(instance)->LogOn(pszToken);
}

// Login to a generic, anonymous account.
//
// Note: in previous versions of the SDK, this was automatically called within SteamGameServer_Init,
// but this is no longer the case.
void SteamAPI_SteamGameServer_LogOnAnonymous(intp instance)
{
    reinterpret_cast<ISteamGameServer *>(instance)->LogOnAnonymous();
}

// Begin process of logging game server out of steam
void SteamAPI_SteamGameServer_LogOff(intp instance)
{
    reinterpret_cast<ISteamGameServer *>(instance)->LogOff();
}

// status functions
bool SteamAPI_SteamGameServer_BLoggedOn(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->BLoggedOn();
}
bool SteamAPI_SteamGameServer_BSecure(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->BSecure();
}
CSteamID SteamAPI_SteamGameServer_GetSteamID(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->GetSteamID();
}

// Returns true if the master server has requested a restart.
// Only returns true once per request.
bool SteamAPI_SteamGameServer_WasRestartRequested(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->WasRestartRequested();
}

//
// Server state.  These properties may be changed at any time.
//

// Max player count that will be reported to server browser and client queries
void SteamAPI_SteamGameServer_SetMaxPlayerCount(intp instance, int cPlayersMax)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetMaxPlayerCount(cPlayersMax);
}

// Number of bots.  Default value is zero
void SteamAPI_SteamGameServer_SetBotPlayerCount(intp instance, int cBotplayers)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetBotPlayerCount(cBotplayers);
}

// Set the name of server as it will appear in the server browser
//
// @see k_cbMaxGameServerName
void SteamAPI_SteamGameServer_SetServerName(intp instance, const char *pszServerName)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetServerName(pszServerName);
}

// Set name of map to report in the server browser
//
// @see k_cbMaxGameServerName
void SteamAPI_SteamGameServer_SetMapName(intp instance, const char *pszMapName)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetMapName(pszMapName);
}

// Let people know if your server will require a password
void SteamAPI_SteamGameServer_SetPasswordProtected(intp instance, bool bPasswordProtected)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetPasswordProtected(bPasswordProtected);
}

// Spectator server.  The default value is zero, meaning the service
// is not used.
void SteamAPI_SteamGameServer_SetSpectatorPort(intp instance, uint16 unSpectatorPort)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetSpectatorPort(unSpectatorPort);
}

// Name of the spectator server.  (Only used if spectator port is nonzero.)
//
// @see k_cbMaxGameServerMapName
void SteamAPI_SteamGameServer_SetSpectatorServerName(intp instance, const char *pszSpectatorServerName)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetSpectatorServerName(pszSpectatorServerName);
}

// Call this to clear the whole list of key/values that are sent in rules queries.
void SteamAPI_SteamGameServer_ClearAllKeyValues(intp instance)
{
    reinterpret_cast<ISteamGameServer *>(instance)->ClearAllKeyValues();
}

// Call this to add/update a key/value pair.
void SteamAPI_SteamGameServer_SetKeyValue(intp instance, const char *pKey, const char *pValue)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetKeyValue(pKey, pValue);
}

// Sets a string defining the "gametags" for this server, this is optional, but if it is set
// it allows users to filter in the matchmaking/server-browser interfaces based on the value
//
// @see k_cbMaxGameServerTags
void SteamAPI_SteamGameServer_SetGameTags(intp instance, const char *pchGameTags)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetGameTags(pchGameTags);
}

// Sets a string defining the "gamedata" for this server, this is optional, but if it is set
// it allows users to filter in the matchmaking/server-browser interfaces based on the value
// don't set this unless it actually changes, its only uploaded to the master once (when
// acknowledged)
//
// @see k_cbMaxGameServerGameData
void SteamAPI_SteamGameServer_SetGameData(intp instance, const char *pchGameData)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetGameData(pchGameData);
}

// Region identifier.  This is an optional field, the default value is empty, meaning the "world" region
void SteamAPI_SteamGameServer_SetRegion(intp instance, const char *pszRegion)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetRegion(pszRegion);
}

//
// Player list management / authentication
//

// Handles receiving a new connection from a Steam user.  This call will ask the Steam
// servers to validate the users identity, app ownership, and VAC status.  If the Steam servers 
// are off-line, then it will validate the cached ticket itself which will validate app ownership 
// and identity.  The AuthBlob here should be acquired on the game client using SteamUser()->InitiateGameConnection()
// and must then be sent up to the game server for authentication.
//
// Return Value: returns true if the users ticket passes basic checks. pSteamIDUser will contain the Steam ID of this user. pSteamIDUser must NOT be NULL
// If the call succeeds then you should expect a GSClientApprove_t or GSClientDeny_t callback which will tell you whether authentication
// for the user has succeeded or failed (the steamid in the callback will match the one returned by this call)
bool SteamAPI_SteamGameServer_SendUserConnectAndAuthenticate(intp instance, uint32 unIPClient, const void *pvAuthBlob, uint32 cubAuthBlobSize, CSteamID *pSteamIDUser)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->SendUserConnectAndAuthenticate(unIPClient, pvAuthBlob, cubAuthBlobSize, pSteamIDUser);
}

// Creates a fake user (ie, a bot) which will be listed as playing on the server, but skips validation.  
// 
// Return Value: Returns a SteamID for the user to be tracked with, you should call HandleUserDisconnect()
// when this user leaves the server just like you would for a real user.
CSteamID SteamAPI_SteamGameServer_CreateUnauthenticatedUserConnection(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->CreateUnauthenticatedUserConnection();
}

// Should be called whenever a user leaves our game server, this lets Steam internally
// track which users are currently on which servers for the purposes of preventing a single
// account being logged into multiple servers, showing who is currently on a server, etc.
void SteamAPI_SteamGameServer_SendUserDisconnect(intp instance, CSteamID steamIDUser)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SendUserDisconnect(steamIDUser);
}

// Update the data to be displayed in the server browser and matchmaking interfaces for a user
// currently connected to the server.  For regular users you must call this after you receive a
// GSUserValidationSuccess callback.
// 
// Return Value: true if successful, false if failure (ie, steamIDUser wasn't for an active player)
bool SteamAPI_SteamGameServer_BUpdateUserData(intp instance, CSteamID steamIDUser, const char *pchPlayerName, uint32 uScore)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->BUpdateUserData(steamIDUser, pchPlayerName, uScore);
}

// New auth system APIs - do not mix with the old auth system APIs.
// ----------------------------------------------------------------

// Retrieve ticket to be sent to the entity who wishes to authenticate you ( using BeginAuthSession API). 
// pcbTicket retrieves the length of the actual ticket.
HAuthTicket SteamAPI_SteamGameServer_GetAuthSessionTicket(intp instance, void *pTicket, int cbMaxTicket, uint32 *pcbTicket)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->GetAuthSessionTicket(pTicket, cbMaxTicket, pcbTicket);
}

// Authenticate ticket ( from GetAuthSessionTicket) from entity steamID to be sure it is valid and isnt reused
// Registers for callbacks if the entity goes offline or cancels the ticket ( see ValidateAuthTicketResponse_t callback and EAuthSessionResponse)
EBeginAuthSessionResult SteamAPI_SteamGameServer_BeginAuthSession(intp instance, const void *pAuthTicket, int cbAuthTicket, CSteamID steamID)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->BeginAuthSession(pAuthTicket, cbAuthTicket, steamID);
}

// Stop tracking started by BeginAuthSession - called when no longer playing game with this entity
void SteamAPI_SteamGameServer_EndAuthSession(intp instance, CSteamID steamID)
{
    reinterpret_cast<ISteamGameServer *>(instance)->EndAuthSession(steamID);
}

// Cancel auth ticket from GetAuthSessionTicket, called when no longer playing game with the entity you gave the ticket to
void SteamAPI_SteamGameServer_CancelAuthTicket(intp instance, HAuthTicket hAuthTicket)
{
    reinterpret_cast<ISteamGameServer *>(instance)->CancelAuthTicket(hAuthTicket);
}

// After receiving a user's authentication data, and passing it to SendUserConnectAndAuthenticate, use this function
// to determine if the user owns downloadable content specified by the provided AppID.
EUserHasLicenseForAppResult SteamAPI_SteamGameServer_UserHasLicenseForApp(intp instance, CSteamID steamID, AppId_t appID)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->UserHasLicenseForApp(steamID, appID);
}

// Ask if a user in in the specified group, results returns async by GSUserGroupStatus_t
// returns false if we're not connected to the steam servers and thus cannot ask
bool SteamAPI_SteamGameServer_RequestUserGroupStatus(intp instance, CSteamID steamIDUser, CSteamID steamIDGroup)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->RequestUserGroupStatus(steamIDUser, steamIDGroup);
}

// Returns the public IP of the server according to Steam, useful when the server is 
// behind NAT and you want to advertise its IP in a lobby for other clients to directly
// connect to
uint32 SteamAPI_SteamGameServer_GetPublicIP(intp instance)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->GetPublicIP();
}

// These are in GameSocketShare mode, where instead of ISteamGameServer creating its own
// socket to talk to the master server on, it lets the game use its socket to forward messages
// back and forth. This prevents us from requiring server ops to open up yet another port
// in their firewalls.
//
// the IP address and port should be in host order, i.e 127.0.0.1 == 0x7f000001

// These are used when you've elected to multiplex the game server's UDP socket
// rather than having the master server updater use its own sockets.
// 
// Source games use this to simplify the job of the server admins, so they 
// don't have to open up more ports on their firewalls.

// Call this when a packet that starts with 0xFFFFFFFF comes in. That means
// it's for us.
bool SteamAPI_SteamGameServer_HandleIncomingPacket(intp instance, const void *pData, int cbData, uint32 srcIP, uint16 srcPort)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->HandleIncomingPacket(pData, cbData, srcIP, srcPort);
}

// AFTER calling HandleIncomingPacket for any packets that came in that frame, call this.
// This gets a packet that the master server updater needs to send out on UDP.
// It returns the length of the packet it wants to send, or 0 if there are no more packets to send.
// Call this each frame until it returns 0.
int SteamAPI_SteamGameServer_GetNextOutgoingPacket(intp instance, void *pOut, int cbMaxOut, uint32 *pNetAdr, uint16 *pPort)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->GetNextOutgoingPacket(pOut, cbMaxOut, pNetAdr, pPort);
}

//
// Control heartbeats / advertisement with master server
//

// Call this as often as you like to tell the master server updater whether or not
// you want it to be active (default: off).
void SteamAPI_SteamGameServer_EnableHeartbeats(intp instance, bool bActive)
{
    reinterpret_cast<ISteamGameServer *>(instance)->EnableHeartbeats(bActive);
}

// You usually don't need to modify this.
// Pass -1 to use the default value for iHeartbeatInterval.
// Some mods change this.
void SteamAPI_SteamGameServer_SetHeartbeatInterval(intp instance, int iHeartbeatInterval)
{
    reinterpret_cast<ISteamGameServer *>(instance)->SetHeartbeatInterval(iHeartbeatInterval);
}

// Force a heartbeat to steam at the next opportunity
void SteamAPI_SteamGameServer_ForceHeartbeat(intp instance)
{
    reinterpret_cast<ISteamGameServer *>(instance)->ForceHeartbeat();
}

// associate this game server with this clan for the purposes of computing player compat
SteamAPICall_t SteamAPI_SteamGameServer_AssociateWithClan(intp instance, CSteamID steamIDClan)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->AssociateWithClan(steamIDClan);
}

// ask if any of the current players dont want to play with this new player - or vice versa
SteamAPICall_t SteamAPI_SteamGameServer_ComputeNewPlayerCompatibility(intp instance, CSteamID steamIDNewPlayer)
{
    return reinterpret_cast<ISteamGameServer *>(instance)->ComputeNewPlayerCompatibility(steamIDNewPlayer);
}

//-----------------------------------------------------------------------------
// Purpose: Functions for authenticating users via Steam to play on a game server
//-----------------------------------------------------------------------------

// downloads stats for the user
// returns a GSStatsReceived_t callback when completed
// if the user has no stats, GSStatsReceived_t.m_eResult will be set to k_EResultFail
// these stats will only be auto-updated for clients playing on the server. For other
// users you'll need to call RequestUserStats() again to refresh any data
SteamAPICall_t SteamAPI_SteamGameServerStats_RequestUserStats(intp instance, CSteamID steamIDUser)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->RequestUserStats(steamIDUser);
}

// requests stat information for a user, usable after a successful call to RequestUserStats()
bool SteamAPI_SteamGameServerStats_GetUserStat(intp instance, CSteamID steamIDUser, const char *pchName, int32 *pData)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->GetUserStat(steamIDUser, pchName, pData);
}
bool SteamAPI_SteamGameServerStats_GetUserStat0(intp instance, CSteamID steamIDUser, const char *pchName, float *pData)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->GetUserStat(steamIDUser, pchName, pData);
}
bool SteamAPI_SteamGameServerStats_GetUserAchievement(intp instance, CSteamID steamIDUser, const char *pchName, bool *pbAchieved)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->GetUserAchievement(steamIDUser, pchName, pbAchieved);
}

// Set / update stats and achievements. 
// Note: These updates will work only on stats game servers are allowed to edit and only for 
// game servers that have been declared as officially controlled by the game creators. 
// Set the IP range of your official servers on the Steamworks page
bool SteamAPI_SteamGameServerStats_SetUserStat(intp instance, CSteamID steamIDUser, const char *pchName, int32 nData)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->SetUserStat(steamIDUser, pchName, nData);
}
bool SteamAPI_SteamGameServerStats_SetUserStat0(intp instance, CSteamID steamIDUser, const char *pchName, float fData)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->SetUserStat(steamIDUser, pchName, fData);
}
bool SteamAPI_SteamGameServerStats_UpdateUserAvgRateStat(intp instance, CSteamID steamIDUser, const char *pchName, float flCountThisSession, double dSessionLength)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->UpdateUserAvgRateStat(steamIDUser, pchName, flCountThisSession, dSessionLength);
}

bool SteamAPI_SteamGameServerStats_SetUserAchievement(intp instance, CSteamID steamIDUser, const char *pchName)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->SetUserAchievement(steamIDUser, pchName);
}
bool SteamAPI_SteamGameServerStats_ClearUserAchievement(intp instance, CSteamID steamIDUser, const char *pchName)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->ClearUserAchievement(steamIDUser, pchName);
}

// Store the current data on the server, will get a GSStatsStored_t callback when set.
//
// If the callback has a result of k_EResultInvalidParam, one or more stats 
// uploaded has been rejected, either because they broke constraints
// or were out of date. In this case the server sends back updated values.
// The stats should be re-iterated to keep in sync.
SteamAPICall_t SteamAPI_SteamGameServerStats_StoreUserStats(intp instance, CSteamID steamIDUser)
{
    return reinterpret_cast<ISteamGameServerStats *>(instance)->StoreUserStats(steamIDUser);
}

}
