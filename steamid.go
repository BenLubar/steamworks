//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type AccountType -trimprefix AccountType
//go:generate stringer -type AccountUniverse -trimprefix Universe
//go:generate stringer -type AccountInstance -trimprefix Instance

package steamworks

import "strconv"

// Most of this code is cannibalized from SteamID.php.
// <https://github.com/xPaw/SteamID.php>

type AccountType uint8

const (
	AccountTypeInvalid        AccountType = 0
	AccountTypeIndividual     AccountType = 1
	AccountTypeMultiseat      AccountType = 2
	AccountTypeGameServer     AccountType = 3
	AccountTypeAnonGameServer AccountType = 4
	AccountTypePending        AccountType = 5
	AccountTypeContentServer  AccountType = 6
	AccountTypeClan           AccountType = 7
	AccountTypeChat           AccountType = 8
	AccountTypeP2PSuperSeeder AccountType = 9
	AccountTypeAnonUser       AccountType = 10
	accountTypeMax            AccountType = 11
)

type AccountUniverse uint8

const (
	UniverseInvalid  AccountUniverse = 0
	UniversePublic   AccountUniverse = 1
	UniverseBeta     AccountUniverse = 2
	UniverseInternal AccountUniverse = 3
	UniverseDev      AccountUniverse = 4
	universeMax      AccountUniverse = 5
)

type AccountInstance uint32

const (
	InstanceDesktop AccountInstance = 1<<0
	InstanceConsole AccountInstance = 1<<1
	InstanceWeb     AccountInstance = 1<<2

	// Special flags for Chat accounts - they go in the top 8 bits
	// of the steam ID's "instance", leaving 12 for the actual instances
	InstanceFlagClan     AccountInstance = 1<<19
	InstanceFlagLobby    AccountInstance = 1<<18
	InstanceFlagMMSLobby AccountInstance = 1<<17
)

type SteamID uint64

func (id SteamID) String() string {
	return id.Steam2String()
}

func (id SteamID) Steam2String() string {
	if id.Type() == AccountTypeInvalid || id.Type() == AccountTypeIndividual {
		universe := uint64(id.Universe())
		accountID := uint64(id.AccountID())
		return "STEAM_" + strconv.FormatUint(universe, 10) + ":" + strconv.FormatUint(accountID&1, 10) + ":" + strconv.FormatUint(accountID>>1, 10)
	}
	return strconv.FormatUint(uint64(id), 10)
}

func (id SteamID) AccountID() uint32 {
	return uint32(id)
}

func (id SteamID) Instance() AccountInstance {
	return AccountInstance((id >> 32) & 0xFFFFF)
}

func (id SteamID) Type() AccountType {
	return AccountType((id >> 52) & 0xF)
}

func (id SteamID) Universe() AccountUniverse {
	return AccountUniverse((id >> 56) & 0xFF)
}

func (id SteamID) IsValid() bool {
	if t := id.Type(); t <= AccountTypeInvalid || t >= accountTypeMax {
		return false
	}

	if u := id.Universe(); u <= UniverseInvalid || u >= universeMax {
		return false
	}

	switch id.Type() {
	case AccountTypeIndividual:
		if id.AccountID() == 0 || id.Instance() > InstanceWeb {
			return false
		}
	case AccountTypeClan:
		if id.AccountID() == 0 || id.Instance() != 0 {
			return false
		}
	case AccountTypeGameServer:
		if id.AccountID() == 0 {
			return false
		}
	}

	return true
}
