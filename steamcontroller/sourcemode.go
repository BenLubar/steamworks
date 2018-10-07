package steamcontroller

import "github.com/BenLubar/steamworks/internal"

// SourceMode is the virtual input mode imposed by the configurator upon a
// controller source. For instance, the configurator can make an analog
// joystick behave like a Dpad with four digital inputs; the Source would be
// SourceJoystick and the SourceMode would be SourceModeDpad. The mode also
// changes the input data received by any associated actions.
type SourceMode = internal.EControllerSourceMode

// SourceMode enum values.
const (
	SourceModeNone           SourceMode = internal.EControllerSourceMode_None
	SourceModeDpad           SourceMode = internal.EControllerSourceMode_Dpad
	SourceModeButtons        SourceMode = internal.EControllerSourceMode_Buttons
	SourceModeFourButtons    SourceMode = internal.EControllerSourceMode_FourButtons
	SourceModeAbsoluteMouse  SourceMode = internal.EControllerSourceMode_AbsoluteMouse
	SourceModeRelativeMouse  SourceMode = internal.EControllerSourceMode_RelativeMouse
	SourceModeJoystickMove   SourceMode = internal.EControllerSourceMode_JoystickMove
	SourceModeJoystickMouse  SourceMode = internal.EControllerSourceMode_JoystickMouse
	SourceModeJoystickCamera SourceMode = internal.EControllerSourceMode_JoystickCamera
	SourceModeScrollWheel    SourceMode = internal.EControllerSourceMode_ScrollWheel
	SourceModeTrigger        SourceMode = internal.EControllerSourceMode_Trigger
	SourceModeTouchMenu      SourceMode = internal.EControllerSourceMode_TouchMenu
	SourceModeMouseJoystick  SourceMode = internal.EControllerSourceMode_MouseJoystick
	SourceModeMouseRegion    SourceMode = internal.EControllerSourceMode_MouseRegion
	SourceModeRadialMenu     SourceMode = internal.EControllerSourceMode_RadialMenu
	SourceModeSingleButton   SourceMode = internal.EControllerSourceMode_SingleButton
	SourceModeSwitches       SourceMode = internal.EControllerSourceMode_Switches
)
