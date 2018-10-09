package steamparentalsettings

import (
	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

func OnChanged(f func()) steamworks.Registration {
	return internal.RegisterCallback_SteamParentalSettingsChanged(func(*internal.SteamParentalSettingsChanged, bool) {
		f()
	}, 0)
}

type Feature internal.EParentalFeature

const (
	Store         Feature = Feature(internal.EParentalFeature_Store)
	Community     Feature = Feature(internal.EParentalFeature_Community)
	Profile       Feature = Feature(internal.EParentalFeature_Profile)
	Friends       Feature = Feature(internal.EParentalFeature_Friends)
	News          Feature = Feature(internal.EParentalFeature_News)
	Trading       Feature = Feature(internal.EParentalFeature_Trading)
	Settings      Feature = Feature(internal.EParentalFeature_Settings)
	Console       Feature = Feature(internal.EParentalFeature_Console)
	Browser       Feature = Feature(internal.EParentalFeature_Browser)
	ParentalSetup Feature = Feature(internal.EParentalFeature_ParentalSetup)
	Library       Feature = Feature(internal.EParentalFeature_Library)
	Test          Feature = Feature(internal.EParentalFeature_Test)
)

func (f Feature) String() string {
	return internal.EParentalFeature(f).String()
}

func (f Feature) IsBlocked() bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsFeatureBlocked(internal.EParentalFeature(f))
}

func (f Feature) IsInBlockList() bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsFeatureInBlockList(internal.EParentalFeature(f))
}

func IsParentalLockEnabled() bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsParentalLockEnabled()
}

func IsParentalLockLocked() bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsParentalLockLocked()
}

func IsAppBlocked(appID steamworks.AppID) bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsAppBlocked(internal.AppId(appID))
}

func IsAppInBlockList(appID steamworks.AppID) bool {
	return internal.SteamAPI_ISteamParentalSettings_BIsAppInBlockList(internal.AppId(appID))
}
