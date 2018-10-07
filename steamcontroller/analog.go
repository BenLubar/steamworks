package steamcontroller

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// AnalogActionHandle is a handle to an analog action. This can be obtained
// from GetAnalogActionHandle.
type AnalogActionHandle = internal.ControllerAnalogActionHandle

// GetAnalogActionData returns the current state of the specified analog game
// action for the specified controller.
//
// NOTE: The exact values, range, etc, depend on the configuration, but
// (broadly speaking) traditional analog actions will provide normalized float
// values in the ballpark of -1.0 to 1.0, whereas mouse-like actions will
// provide delta updates which indicate the number of "pixels" moved since the
// last frame. The upshot of this is that mouse-like actions will provide much
// larger absolute x and y values, and are relative to the last recorded input
// position, whereas traditional analog actions are smaller and relative to a
// central physical anchor point.
//
// While the delta provided by mouse-like actions is very similar to pixel
// deltas as provided by an OS, the SC deltas are floats, not ints. This means
// less potential quantization and loss of precision when mapping this data to
// a camera rotation.
//
// In the case of single-axis analog inputs (such as analog triggers), only the
// x axis will contain data; the y axis will always be zero.
func GetAnalogActionData(controller Handle, analogAction AnalogActionHandle) (x, y float32, mode SourceMode, active bool) {
	defer cleanup()()

	data := internal.SteamAPI_ISteamController_GetAnalogActionData(controller, analogAction)

	return float32(data.X), float32(data.Y), SourceMode(data.EMode), bool(data.BActive)
}

// GetAnalogActionHandle gets the handle of the specified analog action.
//
// NOTE: This function does not take an action set handle parameter. That means
// that each action in your VDF file must have a unique string identifier. In
// other words, if you use an action called "up" in two different action sets,
// this function will only ever return one of them and the other will be
// ignored.
//
// The name refers to an identifier in the game's VDF file.
func GetAnalogActionHandle(name string) AnalogActionHandle {
	defer cleanup()()

	cname := internal.CString(name)
	defer internal.Free(unsafe.Pointer(cname))

	return internal.SteamAPI_ISteamController_GetAnalogActionHandle(cname)
}

// GetAnalogActionOrigins returns a slice containing the origin(s) for an
// analog action within an action set. Use this to display the appropriate
// on-screen prompt for the action.
func GetAnalogActionOrigins(controller Handle, actionSet ActionSetHandle, analogAction AnalogActionHandle) []ActionOrigin {
	defer cleanup()()

	var originsOut [maxOrigins]ActionOrigin

	count := internal.SteamAPI_ISteamController_GetAnalogActionOrigins(controller, actionSet, analogAction, &originsOut[0])

	return originsOut[:count]
}

// StopAnalogActionMomentum stops the momentum of an analog action (where
// applicable, e.g. a touchpad with virtual trackball settings).
//
// NOTE: This will also stop all associated haptics. This is useful for
// situations where you want to indicate to the user that the limit of an
// action has been reached, such as spinning a carousel or scrolling a webpage.
func StopAnalogActionMomentum(controller Handle, analogAction AnalogActionHandle) {
	defer cleanup()()

	internal.SteamAPI_ISteamController_StopAnalogActionMomentum(controller, analogAction)
}
