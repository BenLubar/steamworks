// nolint: golint

package internal

/*
#cgo CXXFLAGS: -std=c++11
#cgo CPPFLAGS: -isystem ${SRCDIR}/include

#include "api.gen.h"
#include "gameserver.h"
*/
import "C"
import "unsafe"

const MasterServerUpdaterPort_UseGameSocketShare uint16 = ^uint16(0)

// GameServerFlag is a set of flags for classifying Steam game servers.
type GameServerFlag uint32

const (
	GameServerFlag_None GameServerFlag = 0x00
	// server has users playing
	GameServerFlag_Active GameServerFlag = 0x01
	// server wants to be secure
	GameServerFlag_Secure GameServerFlag = 0x02
	// server is dedicated
	GameServerFlag_Dedicated GameServerFlag = 0x04
	// linux build
	GameServerFlag_Linux GameServerFlag = 0x08
	// password protected
	GameServerFlag_Passworded GameServerFlag = 0x10
	// server shouldn't list on master server and
	// won't enforce authentication of users that connect to the server.
	// Useful when you run a server where the clients may not
	// be connected to the internet but you want them to play (i.e LANs)
	GameServerFlag_Private GameServerFlag = 0x20
)

func SteamAPI_SteamGameServer_InitGameServer(unIP uint32, usGamePort, usQueryPort uint16, unFlags uint32, nGameAppId AppId, pchVersionString *C.char) bool {
	return bool(C.SteamAPI_SteamGameServer_InitGameServer(getSteamGameServer(), C.uint32(unIP), C.uint16(usGamePort), C.uint16(usQueryPort), C.uint32(unFlags), C.AppId_t(nGameAppId), pchVersionString))
}
func SteamAPI_SteamGameServer_SetProduct(pszProduct *C.char) {
	C.SteamAPI_SteamGameServer_SetProduct(getSteamGameServer(), pszProduct)
}
func SteamAPI_SteamGameServer_SetGameDescription(pszGameDescription *C.char) {
	C.SteamAPI_SteamGameServer_SetGameDescription(getSteamGameServer(), pszGameDescription)
}
func SteamAPI_SteamGameServer_SetModDir(pszModDir *C.char) {
	C.SteamAPI_SteamGameServer_SetModDir(getSteamGameServer(), pszModDir)
}
func SteamAPI_SteamGameServer_SetDedicatedServer(bDedicated bool) {
	C.SteamAPI_SteamGameServer_SetDedicatedServer(getSteamGameServer(), C.bool(bDedicated))
}
func SteamAPI_SteamGameServer_LogOn(pszToken *C.char) {
	C.SteamAPI_SteamGameServer_LogOn(getSteamGameServer(), pszToken)
}
func SteamAPI_SteamGameServer_LogOnAnonymous() {
	C.SteamAPI_SteamGameServer_LogOnAnonymous(getSteamGameServer())
}
func SteamAPI_SteamGameServer_LogOff() {
	C.SteamAPI_SteamGameServer_LogOff(getSteamGameServer())
}
func SteamAPI_SteamGameServer_BLoggedOn() bool {
	return bool(C.SteamAPI_SteamGameServer_BLoggedOn(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_BSecure() bool {
	return bool(C.SteamAPI_SteamGameServer_BSecure(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_GetSteamID() SteamID {
	return SteamID(C.SteamAPI_SteamGameServer_GetSteamID(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_WasRestartRequested() bool {
	return bool(C.SteamAPI_SteamGameServer_WasRestartRequested(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_SetMaxPlayerCount(cPlayersMax int32) {
	C.SteamAPI_SteamGameServer_SetMaxPlayerCount(getSteamGameServer(), C.int(cPlayersMax))
}
func SteamAPI_SteamGameServer_SetBotPlayerCount(cBotplayers int32) {
	C.SteamAPI_SteamGameServer_SetBotPlayerCount(getSteamGameServer(), C.int(cBotplayers))
}
func SteamAPI_SteamGameServer_SetServerName(pszServerName *C.char) {
	C.SteamAPI_SteamGameServer_SetServerName(getSteamGameServer(), pszServerName)
}
func SteamAPI_SteamGameServer_SetMapName(pszMapName *C.char) {
	C.SteamAPI_SteamGameServer_SetMapName(getSteamGameServer(), pszMapName)
}
func SteamAPI_SteamGameServer_SetPasswordProtected(bPasswordProtected bool) {
	C.SteamAPI_SteamGameServer_SetPasswordProtected(getSteamGameServer(), C.bool(bPasswordProtected))
}
func SteamAPI_SteamGameServer_SetSpectatorPort(unSpectatorPort uint16) {
	C.SteamAPI_SteamGameServer_SetSpectatorPort(getSteamGameServer(), C.uint16(unSpectatorPort))
}
func SteamAPI_SteamGameServer_SetSpectatorServerName(pszSpectatorServerName *C.char) {
	C.SteamAPI_SteamGameServer_SetSpectatorServerName(getSteamGameServer(), pszSpectatorServerName)
}
func SteamAPI_SteamGameServer_ClearAllKeyValues() {
	C.SteamAPI_SteamGameServer_ClearAllKeyValues(getSteamGameServer())
}
func SteamAPI_SteamGameServer_SetKeyValue(pKey, pValue *C.char) {
	C.SteamAPI_SteamGameServer_SetKeyValue(getSteamGameServer(), pKey, pValue)
}
func SteamAPI_SteamGameServer_SetGameTags(pchGameTags *C.char) {
	C.SteamAPI_SteamGameServer_SetGameTags(getSteamGameServer(), pchGameTags)
}
func SteamAPI_SteamGameServer_SetGameData(pchGameData *C.char) {
	C.SteamAPI_SteamGameServer_SetGameData(getSteamGameServer(), pchGameData)
}
func SteamAPI_SteamGameServer_SetRegion(pszRegion *C.char) {
	C.SteamAPI_SteamGameServer_SetRegion(getSteamGameServer(), pszRegion)
}
func SteamAPI_SteamGameServer_SendUserConnectAndAuthenticate(unIPClient uint32, pvAuthBlob unsafe.Pointer, cubAuthBlobSize uint32, pSteamIDUser *SteamID) bool {
	return bool(C.SteamAPI_SteamGameServer_SendUserConnectAndAuthenticate(getSteamGameServer(), C.uint32(unIPClient), pvAuthBlob, C.uint32(cubAuthBlobSize), pSteamIDUser))
}
func SteamAPI_SteamGameServer_CreateUnauthenticatedUserConnection() SteamID {
	return SteamID(C.SteamAPI_SteamGameServer_CreateUnauthenticatedUserConnection(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_SendUserDisconnect(steamIDUser SteamID) {
	C.SteamAPI_SteamGameServer_SendUserDisconnect(getSteamGameServer(), steamIDUser)
}
func SteamAPI_SteamGameServer_BUpdateUserData(steamIDUser SteamID, pchPlayerName *C.char, uScore uint32) bool {
	return bool(C.SteamAPI_SteamGameServer_BUpdateUserData(getSteamGameServer(), steamIDUser, pchPlayerName, C.uint32(uScore)))
}
func SteamAPI_SteamGameServer_GetAuthSessionTicket(pTicket unsafe.Pointer, cbMaxTicket int32, pcbTicket *uint32) HAuthTicket {
	return HAuthTicket(C.SteamAPI_SteamGameServer_GetAuthSessionTicket(getSteamGameServer(), pTicket, C.int32(cbMaxTicket), (*C.uint32)(pcbTicket)))
}
func SteamAPI_SteamGameServer_BeginAuthSession(pAuthTicket unsafe.Pointer, cbAuthTicket int32, steamID SteamID) EBeginAuthSessionResult {
	return EBeginAuthSessionResult(C.SteamAPI_SteamGameServer_BeginAuthSession(getSteamGameServer(), pAuthTicket, C.int32(cbAuthTicket), steamID))
}
func SteamAPI_SteamGameServer_EndAuthSession(steamID SteamID) {
	C.SteamAPI_SteamGameServer_EndAuthSession(getSteamGameServer(), steamID)
}
func SteamAPI_SteamGameServer_CancelAuthTicket(hAuthTicket HAuthTicket) {
	C.SteamAPI_SteamGameServer_CancelAuthTicket(getSteamGameServer(), hAuthTicket)
}
func SteamAPI_SteamGameServer_UserHasLicenseForApp(steamID SteamID, appID AppId) EUserHasLicenseForAppResult {
	return EUserHasLicenseForAppResult(C.SteamAPI_SteamGameServer_UserHasLicenseForApp(getSteamGameServer(), steamID, appID))
}
func SteamAPI_SteamGameServer_RequestUserGroupStatus(steamIDUser, steamIDGroup SteamID) bool {
	return bool(C.SteamAPI_SteamGameServer_RequestUserGroupStatus(getSteamGameServer(), steamIDUser, steamIDGroup))
}
func SteamAPI_SteamGameServer_GetPublicIP() uint32 {
	return uint32(C.SteamAPI_SteamGameServer_GetPublicIP(getSteamGameServer()))
}
func SteamAPI_SteamGameServer_HandleIncomingPacket(pData unsafe.Pointer, cbData int32, srcIP uint32, srcPort uint16) bool {
	return bool(C.SteamAPI_SteamGameServer_HandleIncomingPacket(getSteamGameServer(), pData, C.int32(cbData), C.uint32(srcIP), C.uint16(srcPort)))
}
func SteamAPI_SteamGameServer_GetNextOutgoingPacket(pOut unsafe.Pointer, cbMaxOut int32, pNetAdr *uint32, pPort *uint16) int {
	return int(C.SteamAPI_SteamGameServer_GetNextOutgoingPacket(getSteamGameServer(), pOut, C.int32(cbMaxOut), (*C.uint32)(pNetAdr), (*C.uint16)(pPort)))
}
func SteamAPI_SteamGameServer_EnableHeartbeats(bActive bool) {
	C.SteamAPI_SteamGameServer_EnableHeartbeats(getSteamGameServer(), C.bool(bActive))
}
func SteamAPI_SteamGameServer_SetHeartbeatInterval(iHeartbeatInterval int32) {
	C.SteamAPI_SteamGameServer_SetHeartbeatInterval(getSteamGameServer(), C.int32(iHeartbeatInterval))
}
func SteamAPI_SteamGameServer_ForceHeartbeat() {
	C.SteamAPI_SteamGameServer_ForceHeartbeat(getSteamGameServer())
}
func SteamAPI_SteamGameServer_AssociateWithClan(steamIDClan SteamID) SteamAPICall {
	return SteamAPICall(C.SteamAPI_SteamGameServer_AssociateWithClan(getSteamGameServer(), steamIDClan))
}
func SteamAPI_SteamGameServer_ComputeNewPlayerCompatibility(steamIDNewPlayer SteamID) SteamAPICall {
	return SteamAPICall(C.SteamAPI_SteamGameServer_ComputeNewPlayerCompatibility(getSteamGameServer(), steamIDNewPlayer))
}

func SteamAPI_SteamGameServerStats_RequestUserStats(steamIDUser SteamID) SteamAPICall {
	return SteamAPICall(C.SteamAPI_SteamGameServerStats_RequestUserStats(getSteamGameServerStats(), steamIDUser))
}
func SteamAPI_SteamGameServerStats_GetUserStat(steamIDUser SteamID, pchName *C.char, pData *int32) bool {
	return bool(C.SteamAPI_SteamGameServerStats_GetUserStat(getSteamGameServerStats(), steamIDUser, pchName, (*C.int32)(pData)))
}
func SteamAPI_SteamGameServerStats_GetUserStat0(steamIDUser SteamID, pchName *C.char, pData *float32) bool {
	return bool(C.SteamAPI_SteamGameServerStats_GetUserStat0(getSteamGameServerStats(), steamIDUser, pchName, (*C.float)(pData)))
}
func SteamAPI_SteamGameServerStats_GetUserAchievement(steamIDUser SteamID, pchName *C.char, pbAchieved *bool) bool {
	return bool(C.SteamAPI_SteamGameServerStats_GetUserAchievement(getSteamGameServerStats(), steamIDUser, pchName, (*C.bool)(pbAchieved)))
}
func SteamAPI_SteamGameServerStats_SetUserStat(steamIDUser SteamID, pchName *C.char, nData int32) bool {
	return bool(C.SteamAPI_SteamGameServerStats_SetUserStat(getSteamGameServerStats(), steamIDUser, pchName, C.int32(nData)))
}
func SteamAPI_SteamGameServerStats_SetUserStat0(steamIDUser SteamID, pchName *C.char, fData float32) bool {
	return bool(C.SteamAPI_SteamGameServerStats_SetUserStat0(getSteamGameServerStats(), steamIDUser, pchName, C.float(fData)))
}
func SteamAPI_SteamGameServerStats_UpdateUserAvgRateStat(steamIDUser SteamID, pchName *C.char, flCountThisSession float32, dSessionLength float64) bool {
	return bool(C.SteamAPI_SteamGameServerStats_UpdateUserAvgRateStat(getSteamGameServerStats(), steamIDUser, pchName, C.float(flCountThisSession), C.double(dSessionLength)))
}
func SteamAPI_SteamGameServerStats_SetUserAchievement(steamIDUser SteamID, pchName *C.char) bool {
	return bool(C.SteamAPI_SteamGameServerStats_SetUserAchievement(getSteamGameServerStats(), steamIDUser, pchName))
}
func SteamAPI_SteamGameServerStats_ClearUserAchievement(steamIDUser SteamID, pchName *C.char) bool {
	return bool(C.SteamAPI_SteamGameServerStats_ClearUserAchievement(getSteamGameServerStats(), steamIDUser, pchName))
}
func SteamAPI_SteamGameServerStats_StoreUserStats(steamIDUser SteamID) SteamAPICall {
	return SteamAPICall(C.SteamAPI_SteamGameServerStats_StoreUserStats(getSteamGameServerStats(), steamIDUser))
}
