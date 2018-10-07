// Package steamcontroller wraps Steam's controller input API.
//
// This package is only available on clients.
//
// Unlike the C++ Steam API, this package automatically calls Init
// and Shutdown as needed.
//
// See the Steam Input documentation for more details.
// <https://partner.steamgames.com/doc/features/steam_controller>
package steamcontroller

import "github.com/BenLubar/steamworks/internal"

// STEAM_CONTROLLER_MAX_COUNT
const maxControllers = 16

var initOnce internal.Once

func cleanup() func() {
	c := internal.Cleanup()

	initOnce.Do(func() {
		internal.SteamAPI_ISteamController_Init()

		internal.OnShutdown(func() {
			internal.SteamAPI_ISteamController_Shutdown()
		})
	})

	return c
}

// RunFrame synchronizes API state with the latest Steam Controller inputs
// available. This is performed automatically by steamworks.RunCallbacks, but
// for the absolute lowest possible latency, you can call this directly before
// reading controller state.
func RunFrame() {
	defer cleanup()()

	internal.SteamAPI_ISteamController_RunFrame()
}

// Handle consistently identifies a controller, even if it is disconnected and
// re-connected.
type Handle = internal.ControllerHandle

// AllControllers is a special value that can be used in place of a specific
// controller handle to send the option to all controllers instead.
const AllControllers Handle = Handle(^uint64(0))

// GetConnectedControllers enumerates currently connected controllers.
func GetConnectedControllers() []Handle {
	defer cleanup()()

	var handlesOut [maxControllers]Handle

	count := internal.SteamAPI_ISteamController_GetConnectedControllers(&handlesOut[0])

	return handlesOut[:count]
}

// GetControllerForGamepadIndex returns the associated controller handle for the
// specified emulated gamepad.
func GetControllerForGamepadIndex(index int) Handle {
	defer cleanup()()

	return internal.SteamAPI_ISteamController_GetControllerForGamepadIndex(int32(index))
}

// GetGamepadIndexForController returns the associated gamepad index for the
// specified controller, if emulating a gamepad.
func GetGamepadIndexForController(controller Handle) int {
	defer cleanup()()

	return int(internal.SteamAPI_ISteamController_GetGamepadIndexForController(controller))
}

// ControllerMotionData represents the current state of a device's motion
// sensor(s).
type ControllerMotionData struct {
	// RotQuat is the sensor-fused absolute rotation.
	//
	// NOTE: The inertial measurement unit on the controller will create a
	// quaternion based on fusing the gyro and the accelerometer. This value is
	// the absolute orientation of the controller, but it will drift on the yaw
	// axis.
	RotQuat [4]float32

	// PosAccel is the positional acceleration.
	PosAccel [3]float32

	// RotVel is the angular velocity
	RotVel [3]float32
}

// GetMotionData returns raw motion data for the specified controller.
func GetMotionData(controller Handle) ControllerMotionData {
	defer cleanup()()

	data := internal.SteamAPI_ISteamController_GetMotionData(controller)

	return ControllerMotionData{
		RotQuat: [4]float32{
			float32(data.RotQuatX),
			float32(data.RotQuatY),
			float32(data.RotQuatZ),
			float32(data.RotQuatW),
		},
		PosAccel: [3]float32{
			float32(data.PosAccelX),
			float32(data.PosAccelY),
			float32(data.PosAccelZ),
		},
		RotVel: [3]float32{
			float32(data.RotVelX),
			float32(data.RotVelY),
			float32(data.RotVelZ),
		},
	}
}

// SetLEDColor sets the controller LED color on supported controllers.
//
// NOTE: The VSC does not support any color but white, and will interpret the
// RGB values as a greyscale value affecting the brightness of the Steam button
// LED. The DS4 responds to full color information and uses the values to set
// the color and brightness of the lightbar.
func SetLEDColor(controller Handle, r, g, b uint8) {
	defer cleanup()()

	internal.SteamAPI_ISteamController_SetLEDColor(controller, r, g, b, uint32(internal.ESteamControllerLEDFlag_SetColor))
}

// ResetLEDColor restores the out-of-game default color for the specified
// controller.
func ResetLEDColor(controller Handle) {
	defer cleanup()()

	internal.SteamAPI_ISteamController_SetLEDColor(controller, 0, 0, 0, uint32(internal.ESteamControllerLEDFlag_RestoreUserDefault))
}

// ShowBindingPanel invokes the Steam overlay and brings up the binding screen.
//
// Returns true for success; false if the overlay is disabled or unavailable,
// or if the user is not in Big Picture mode.
func ShowBindingPanel(controller Handle) bool {
	defer cleanup()()

	return internal.SteamAPI_ISteamController_ShowBindingPanel(controller)
}
