// Package steamutils wraps miscellaneous Steam utility functions.
//
// This package includes functions to help monitor computer state and interact
// with the Steam overlay and VR.
//
// This package works with both clients and servers, but the Steam overlay and
// VR-related functions do not work on headless servers.
//
// See the ISteamUtils documentation for more details.
// <https://partner.steamgames.com/doc/api/ISteamUtils>
package steamutils

import (
	"time"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

// CurrentBatteryPower returns the current battery power percentage from
// 0 to 100, or 255 if the user is on AC power.
func CurrentBatteryPower() uint8 {
	return internal.SteamAPI_ISteamUtils_GetCurrentBatteryPower()
}

// OnLowBatteryPower registers a function to be called when the computer is
// running out of power.
//
// The function is called when the computer has less than ten minutes of power
// remaining, and again every minute after that.
func OnLowBatteryPower(f func(remaining time.Duration)) steamworks.Registration {
	return internal.RegisterCallback_LowBatteryPower(func(data *internal.LowBatteryPower, _ bool) {
		f(time.Duration(data.NMinutesBatteryLeft) * time.Minute)
	}, 0)
}

// IPCountry returns the 2 digit ISO 3166-1-alpha-2 format country code which
// client is running in. (e.g. "US" or "UK")
//
// This is looked up via an IP-to-location database.
func IPCountry() string {
	return internal.GoString(internal.SteamAPI_ISteamUtils_GetIPCountry())
}

// OnIPCountryChanged registers a function to be called when the user's country
// changes. Call IPCountry to retrieve the new country code.
func OnIPCountryChanged(f func()) steamworks.Registration {
	return internal.RegisterCallback_IPCountry(func(*internal.IPCountry, bool) {
		f()
	}, 0)
}

// SecondsSinceAppActive returns the number of seconds since the application
// was active.
func SecondsSinceAppActive() time.Duration {
	return time.Duration(internal.SteamAPI_ISteamUtils_GetSecondsSinceAppActive()) * time.Second
}

// SecondsSinceComputerActive returns the number of seconds since the user last
// moved the mouse.
func SecondsSinceComputerActive() time.Duration {
	return time.Duration(internal.SteamAPI_ISteamUtils_GetSecondsSinceComputerActive()) * time.Second
}

// ServerRealTime returns the Steam server time to the nearest second.
func ServerRealTime() time.Time {
	return time.Unix(int64(internal.SteamAPI_ISteamUtils_GetServerRealTime()), 0)
}

// OnSteamShutdown registers a function to be called when Steam wants to shut
// down.
func OnSteamShutdown(f func()) steamworks.Registration {
	return internal.RegisterCallback_SteamShutdown(func(*internal.SteamShutdown, bool) { f() }, 0)
}
