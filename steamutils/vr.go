package steamutils

import "github.com/BenLubar/steamworks/internal"

// IsSteamRunningInVR returns true if Steam itself is running in VR mode.
func IsSteamRunningInVR() bool {
	return internal.SteamAPI_ISteamUtils_IsSteamRunningInVR()
}

// StartVRDashboard asks Steam to create and render the OpenVR dashboard.
func StartVRDashboard() {
	defer internal.Cleanup()()

	internal.SteamAPI_ISteamUtils_StartVRDashboard()
}

// IsVRHeadsetStreamingEnabled checks if the HMD view will be streamed via
// Steam In-Home Streaming.
func IsVRHeadsetStreamingEnabled() bool {
	return internal.SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled()
}

// SetVRHeadsetStreamingEnabled sets whether the HMD content will be streamed
// via Steam In-Home Streaming.
//
// If this is enabled, then the scene in the HMD headset will be streamed, and
// remote input will not be allowed. Otherwise if this is disabled, then the
// application window will be streamed instead, and remote input will be
// allowed. VR games default to enabled unless "VRHeadsetStreaming" "0" is in
// the extended appinfo for a game.
//
// This is useful for games that have asymmetric multiplayer gameplay.
func SetVRHeadsetStreamingEnabled(enabled bool) {
	internal.SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled(enabled)
}
