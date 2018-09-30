package steamworks

import (
	"strconv"

	"github.com/BenLubar/steamworks/internal"
)

const (
	// SteamIDNil is a generic invalid SteamID.
	SteamIDNil SteamID = 0

	// SteamIDOutofDateGS is a SteamID that comes from a user game connection to an out-of-date game server that hasn't implemented the protocol to provide its SteamID.
	SteamIDOutofDateGS SteamID = SteamID(0) | SteamID(AccountTypeInvalid)<<52 | SteamID(UniverseInvalid)<<56
	// SteamIDLanModeGS is a SteamID that comes from a user game connection to an sv_lan game server.
	SteamIDLanModeGS SteamID = SteamID(0) | SteamID(AccountTypeInvalid)<<52 | SteamID(UniversePublic)<<56
	// SteamIDNotInitYetGS is a SteamID that can come from a user game connection to a game server that has just booted but hasn't yet even initialized its Steam3 component and started logging on.
	SteamIDNotInitYetGS SteamID = SteamID(1) | SteamID(AccountTypeInvalid)<<52 | SteamID(UniverseInvalid)<<56
	// SteamIDNonSteamGS is a SteamID that can come from a user game connection to a GS that isn't using the steam authentication system but still wants to support the "Join Game" option in the friends list.
	SteamIDNonSteamGS SteamID = SteamID(2) | SteamID(AccountTypeInvalid)<<52 | SteamID(UniverseInvalid)<<56
)

// AccountType is a Steam account type.
type AccountType = internal.EAccountType

// Constants for AccountType
const (
	AccountTypeInvalid        AccountType = internal.EAccountType_Invalid
	AccountTypeIndividual     AccountType = internal.EAccountType_Individual
	AccountTypeMultiseat      AccountType = internal.EAccountType_Multiseat
	AccountTypeGameServer     AccountType = internal.EAccountType_GameServer
	AccountTypeAnonGameServer AccountType = internal.EAccountType_AnonGameServer
	AccountTypePending        AccountType = internal.EAccountType_Pending
	AccountTypeContentServer  AccountType = internal.EAccountType_ContentServer
	AccountTypeClan           AccountType = internal.EAccountType_Clan
	AccountTypeChat           AccountType = internal.EAccountType_Chat
	AccountTypeConsoleUser    AccountType = internal.EAccountType_ConsoleUser
	AccountTypeAnonUser       AccountType = internal.EAccountType_AnonUser
)

// AccountUniverse is a Steam universe.
type AccountUniverse = internal.EUniverse

// Constants for AccountUniverse
const (
	UniverseInvalid  AccountUniverse = internal.EUniverse_Invalid
	UniversePublic   AccountUniverse = internal.EUniverse_Public
	UniverseBeta     AccountUniverse = internal.EUniverse_Beta
	UniverseInternal AccountUniverse = internal.EUniverse_Internal
	UniverseDev      AccountUniverse = internal.EUniverse_Dev
)

// AccountInstance represents an instance of a Steam account.
type AccountInstance uint32

// Constants for AccountInstance
const (
	InstanceDesktop AccountInstance = 1 << 0
	InstanceConsole AccountInstance = 1 << 1
	InstanceWeb     AccountInstance = 1 << 2

	// Special flags for Chat accounts - they go in the top 8 bits
	// of the steam ID's "instance", leaving 12 for the actual instances
	InstanceFlagClan     AccountInstance = 1 << 19
	InstanceFlagLobby    AccountInstance = 1 << 18
	InstanceFlagMMSLobby AccountInstance = 1 << 17
)

// SteamID is a 64-bit ID representing an object within the Steam "multiverse".
type SteamID uint64

func (id SteamID) String() string {
	return id.Steam2String()
}

// Steam2String returns the Steam2 string representation of this ID.
//
// If the account type is AccountTypeInvalid or AccountTypeIndividual
func (id SteamID) Steam2String() string {
	if id.Type() == AccountTypeInvalid || id.Type() == AccountTypeIndividual {
		universe := uint64(id.Universe())
		accountID := uint64(id.AccountID())
		return "STEAM_" + strconv.FormatUint(universe, 10) + ":" + strconv.FormatUint(accountID&1, 10) + ":" + strconv.FormatUint(accountID>>1, 10)
	}
	return strconv.FormatUint(uint64(id), 10)
}

// AccountID extracts the account ID from this SteamID.
func (id SteamID) AccountID() uint32 {
	return uint32(id)
}

// Instance extracts the account instance from this SteamID.
//
// The meaning of the account instance is different for each account type.
func (id SteamID) Instance() AccountInstance {
	return AccountInstance((id >> 32) & 0xFFFFF)
}

// Type extracts the account type from this SteamID.
//
// The most common account type you will interact with is AccountTypeIndividual.
func (id SteamID) Type() AccountType {
	return AccountType((id >> 52) & 0xF)
}

// Universe extracts the universe from this SteamID.
//
// It is safe to assume the universe is UniversePublic in any valid SteamID.
// Other universes are only used within Valve's internal network.
func (id SteamID) Universe() AccountUniverse {
	return AccountUniverse((id >> 56) & 0xFF)
}

// IsValid returns true if the SteamID has a valid format. It does not check
// whether the target of the ID exists.
func (id SteamID) IsValid() bool {
	return internal.SteamID_IsValid(internal.SteamID(id))
}

// GetSteamID returns the Steam ID associated with the current user or game
// server.
func GetSteamID() SteamID {
	if internal.IsGameServer {
		return SteamID(internal.SteamAPI_ISteamGameServer_GetSteamID())
	}

	return SteamID(internal.SteamAPI_ISteamUser_GetSteamID())
}
