package steamworks

import (
	"hash/crc32"
	"path/filepath"

	"github.com/BenLubar/steamworks/internal"
)

// AppID is a numeric identifier for a Steam app.
type AppID uint32

// GameID returns the base GameID for this AppID.
func (id AppID) GameID() GameID {
	return GameID(id)
}

// GameIDType is the type of a GameID.
type GameIDType = internal.EGameIDType

const (
	// TypeApp is a Steam app.
	TypeApp GameIDType = internal.EGameIDType_App
	// TypeGameMod is a modification of a Steam app.
	TypeGameMod GameIDType = internal.EGameIDType_GameMod
	// TypeShortcut is a shortcut to a non-Steam game.
	TypeShortcut GameIDType = internal.EGameIDType_Shortcut
	_            GameIDType = internal.EGameIDType_P2P
)

// GameID identifies a game in Steam.
type GameID uint64

func crc(strs ...string) GameID {
	h := crc32.NewIEEE()
	for _, s := range strs {
		_, _ = h.Write([]byte(s)) // nolint: gosec
	}
	return GameID(h.Sum32()|0x80000000) << 32
}

// NewModID returns a GameID for a modification of a Steam app with the given
// base directory path.
func NewModID(app AppID, path string) GameID {
	base := filepath.Base(path)
	if ext := filepath.Ext(base); ext != "" {
		base = base[:len(base)-len(ext)]
	}

	return crc(base) | GameID(TypeGameMod)<<24 | GameID(app)
}

// NewShortcutID returns a GameID for a shortcut.
func NewShortcutID(exePath, appName string) GameID {
	return crc(exePath, appName) | GameID(TypeShortcut)<<24
}

// AppID returns the AppID for this GameID.
func (id GameID) AppID() AppID {
	return AppID(id & 0xffffff)
}

// Type returns the type of this GameID.
func (id GameID) Type() GameIDType {
	return GameIDType((id >> 24) & 0xff)
}

// GetAppID returns the App ID of the current process.
func GetAppID() AppID {
	return AppID(internal.SteamAPI_ISteamUtils_GetAppID())
}
