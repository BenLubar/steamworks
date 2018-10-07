package steamcontroller

import (
	"time"

	"github.com/BenLubar/steamworks/internal"
)

// Pad is a touchpad region on a Steam Controller Device.
//
// On the VSC, the values correspond to the left & right haptic touchpads.
//
// On the DS4, the values correspond to the left & right halves of the single,
// central touchpad.
type Pad = internal.ESteamControllerPad

// Pad enum values.
const (
	PadLeft  Pad = internal.ESteamControllerPad_Left
	PadRight Pad = internal.ESteamControllerPad_Right
)

func micros(duration time.Duration) uint16 {
	if d := duration / time.Microsecond; time.Duration(uint16(d)) == d {
		return uint16(d)
	}

	panic("steamcontroller: duration " + duration.String() + " exceeds maximum value of 65535 microseconds")
}

// Triggers a (low-level) haptic pulse on supported controllers.
//
// NOTE: Currently only the VSC supports haptic pulses. This API call will be
// ignored for all other controller models.
//
// The longest haptic pulse you can trigger with this method has a duration of
// 0.065535 seconds (i.e., less than 1/10th of a second). This function should
// be thought of as a low-level primitive meant to be repeatedly used in
// higher-level user functions to generate more sophisticated behavior.
func TriggerHapticPulse(controller Handle, targetPad Pad, duration time.Duration) {
	defer cleanup()()

	internal.SteamAPI_ISteamController_TriggerHapticPulse(controller, targetPad, micros(duration))
}

// TriggerRepeatedHapticPulse triggers a repeated haptic pulse on supported
// controllers.
//
// NOTE: Currently only the VSC supports haptic pulses. This API call will be
// ignored for incompatible controller models.
//
// This is a more user-friendly function to call than TriggerHapticPulse as it
// can generate pulse patterns long enough to be actually noticed by the user.
//
// Changing the duration and off parameters will change the "texture" of the
// haptic pulse. The maximum value for either parameter is 0.065535 seconds.
func TriggerRepeatedHapticPulse(controller Handle, targetPad Pad, duration, off time.Duration, repeats uint16) {
	internal.SteamAPI_ISteamController_TriggerRepeatedHapticPulse(controller, targetPad, micros(duration), micros(off), repeats, 0)
}

// TriggerVibration triggers a vibration event on supported controllers.
//
// NOTE: This API call will be ignored for incompatible controller models. This
// generates the traditional "rumble" vibration effect. The VSC will emulate
// traditional rumble using its haptics.
//
// leftSpeed and rightSpeed are the period of the corresponding rumble motor's
// vibration. The maximum value for either parameter is 0.065535 seconds.
func TriggerVibration(controller Handle, leftSpeed, rightSpeed time.Duration) {
	internal.SteamAPI_ISteamController_TriggerVibration(controller, micros(leftSpeed), micros(rightSpeed))
}
