package steamcontroller

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// DigitalActionHandle is a handle to a digital action. This can be obtained
// from GetDigitalActionHandle.
type DigitalActionHandle = internal.ControllerDigitalActionHandle

// GetDigitalActionData returns the current state of the specified digital game
// action.
func GetDigitalActionData(controller Handle, digitalAction DigitalActionHandle) (state, active bool) {
	defer cleanup()()

	data := internal.SteamAPI_ISteamController_GetDigitalActionData(controller, digitalAction)

	return bool(data.BState), bool(data.BActive)
}

// GetDigitalActionHandle gets the handle of the specified digital action.
//
// NOTE: This function does not take an action set handle parameter. That means
// that each action in your VDF file must have a unique string identifier. In
// other words, if you use an action called "up" in two different action sets,
// this function will only ever return one of them and the other will be
// ignored.
//
// The name refers to an identifier in the game's VDF file.
func GetDigitalActionHandle(name string) DigitalActionHandle {
	defer cleanup()()

	cname := internal.CString(name)
	defer internal.Free(unsafe.Pointer(cname))

	return internal.SteamAPI_ISteamController_GetDigitalActionHandle(cname)
}

// GetDigitalActionOrigins returns a slice containing the origin(s) for a
// digital action within an action set. Use this to display the appropriate
// on-screen prompt for the action.
func GetDigitalActionOrigins(controller Handle, actionSet ActionSetHandle, digitalAction DigitalActionHandle) []ActionOrigin {
	defer cleanup()()

	var originsOut [maxOrigins]ActionOrigin

	count := internal.SteamAPI_ISteamController_GetDigitalActionOrigins(controller, actionSet, digitalAction, &originsOut[0])

	return originsOut[:count]
}
