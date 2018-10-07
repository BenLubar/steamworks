package steamcontroller

import "github.com/BenLubar/steamworks/internal"

// STEAM_CONTROLLER_MAX_ORIGINS
const maxOrigins = 8

// ActionOrigin represents an input the player binds to an action in the Steam
// Controller Configurator. The chief purpose of these values is to direct
// which on-screen button glyphs should appear for a given action, such as
// "Press [A] to Jump".
type ActionOrigin = internal.EControllerActionOrigin

// ActionOrigin enum values.
const (
	AONone                             ActionOrigin = internal.EControllerActionOrigin_None
	AOA                                ActionOrigin = internal.EControllerActionOrigin_A
	AOB                                ActionOrigin = internal.EControllerActionOrigin_B
	AOX                                ActionOrigin = internal.EControllerActionOrigin_X
	AOY                                ActionOrigin = internal.EControllerActionOrigin_Y
	AOLeftBumper                       ActionOrigin = internal.EControllerActionOrigin_LeftBumper
	AORightBumper                      ActionOrigin = internal.EControllerActionOrigin_RightBumper
	AOLeftGrip                         ActionOrigin = internal.EControllerActionOrigin_LeftGrip
	AORightGrip                        ActionOrigin = internal.EControllerActionOrigin_RightGrip
	AOStart                            ActionOrigin = internal.EControllerActionOrigin_Start
	AOBack                             ActionOrigin = internal.EControllerActionOrigin_Back
	AOLeftPad_Touch                    ActionOrigin = internal.EControllerActionOrigin_LeftPad_Touch
	AOLeftPad_Swipe                    ActionOrigin = internal.EControllerActionOrigin_LeftPad_Swipe
	AOLeftPad_Click                    ActionOrigin = internal.EControllerActionOrigin_LeftPad_Click
	AOLeftPad_DPadNorth                ActionOrigin = internal.EControllerActionOrigin_LeftPad_DPadNorth
	AOLeftPad_DPadSouth                ActionOrigin = internal.EControllerActionOrigin_LeftPad_DPadSouth
	AOLeftPad_DPadWest                 ActionOrigin = internal.EControllerActionOrigin_LeftPad_DPadWest
	AOLeftPad_DPadEast                 ActionOrigin = internal.EControllerActionOrigin_LeftPad_DPadEast
	AORightPad_Touch                   ActionOrigin = internal.EControllerActionOrigin_RightPad_Touch
	AORightPad_Swipe                   ActionOrigin = internal.EControllerActionOrigin_RightPad_Swipe
	AORightPad_Click                   ActionOrigin = internal.EControllerActionOrigin_RightPad_Click
	AORightPad_DPadNorth               ActionOrigin = internal.EControllerActionOrigin_RightPad_DPadNorth
	AORightPad_DPadSouth               ActionOrigin = internal.EControllerActionOrigin_RightPad_DPadSouth
	AORightPad_DPadWest                ActionOrigin = internal.EControllerActionOrigin_RightPad_DPadWest
	AORightPad_DPadEast                ActionOrigin = internal.EControllerActionOrigin_RightPad_DPadEast
	AOLeftTrigger_Pull                 ActionOrigin = internal.EControllerActionOrigin_LeftTrigger_Pull
	AOLeftTrigger_Click                ActionOrigin = internal.EControllerActionOrigin_LeftTrigger_Click
	AORightTrigger_Pull                ActionOrigin = internal.EControllerActionOrigin_RightTrigger_Pull
	AORightTrigger_Click               ActionOrigin = internal.EControllerActionOrigin_RightTrigger_Click
	AOLeftStick_Move                   ActionOrigin = internal.EControllerActionOrigin_LeftStick_Move
	AOLeftStick_Click                  ActionOrigin = internal.EControllerActionOrigin_LeftStick_Click
	AOLeftStick_DPadNorth              ActionOrigin = internal.EControllerActionOrigin_LeftStick_DPadNorth
	AOLeftStick_DPadSouth              ActionOrigin = internal.EControllerActionOrigin_LeftStick_DPadSouth
	AOLeftStick_DPadWest               ActionOrigin = internal.EControllerActionOrigin_LeftStick_DPadWest
	AOLeftStick_DPadEast               ActionOrigin = internal.EControllerActionOrigin_LeftStick_DPadEast
	AOGyro_Move                        ActionOrigin = internal.EControllerActionOrigin_Gyro_Move
	AOGyro_Pitch                       ActionOrigin = internal.EControllerActionOrigin_Gyro_Pitch
	AOGyro_Yaw                         ActionOrigin = internal.EControllerActionOrigin_Gyro_Yaw
	AOGyro_Roll                        ActionOrigin = internal.EControllerActionOrigin_Gyro_Roll
	AOPS4_X                            ActionOrigin = internal.EControllerActionOrigin_PS4_X
	AOPS4_Circle                       ActionOrigin = internal.EControllerActionOrigin_PS4_Circle
	AOPS4_Triangle                     ActionOrigin = internal.EControllerActionOrigin_PS4_Triangle
	AOPS4_Square                       ActionOrigin = internal.EControllerActionOrigin_PS4_Square
	AOPS4_LeftBumper                   ActionOrigin = internal.EControllerActionOrigin_PS4_LeftBumper
	AOPS4_RightBumper                  ActionOrigin = internal.EControllerActionOrigin_PS4_RightBumper
	AOPS4_Options                      ActionOrigin = internal.EControllerActionOrigin_PS4_Options
	AOPS4_Share                        ActionOrigin = internal.EControllerActionOrigin_PS4_Share
	AOPS4_LeftPad_Touch                ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_Touch
	AOPS4_LeftPad_Swipe                ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_Swipe
	AOPS4_LeftPad_Click                ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_Click
	AOPS4_LeftPad_DPadNorth            ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_DPadNorth
	AOPS4_LeftPad_DPadSouth            ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_DPadSouth
	AOPS4_LeftPad_DPadWest             ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_DPadWest
	AOPS4_LeftPad_DPadEast             ActionOrigin = internal.EControllerActionOrigin_PS4_LeftPad_DPadEast
	AOPS4_RightPad_Touch               ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_Touch
	AOPS4_RightPad_Swipe               ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_Swipe
	AOPS4_RightPad_Click               ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_Click
	AOPS4_RightPad_DPadNorth           ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_DPadNorth
	AOPS4_RightPad_DPadSouth           ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_DPadSouth
	AOPS4_RightPad_DPadWest            ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_DPadWest
	AOPS4_RightPad_DPadEast            ActionOrigin = internal.EControllerActionOrigin_PS4_RightPad_DPadEast
	AOPS4_CenterPad_Touch              ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_Touch
	AOPS4_CenterPad_Swipe              ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_Swipe
	AOPS4_CenterPad_Click              ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_Click
	AOPS4_CenterPad_DPadNorth          ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_DPadNorth
	AOPS4_CenterPad_DPadSouth          ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_DPadSouth
	AOPS4_CenterPad_DPadWest           ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_DPadWest
	AOPS4_CenterPad_DPadEast           ActionOrigin = internal.EControllerActionOrigin_PS4_CenterPad_DPadEast
	AOPS4_LeftTrigger_Pull             ActionOrigin = internal.EControllerActionOrigin_PS4_LeftTrigger_Pull
	AOPS4_LeftTrigger_Click            ActionOrigin = internal.EControllerActionOrigin_PS4_LeftTrigger_Click
	AOPS4_RightTrigger_Pull            ActionOrigin = internal.EControllerActionOrigin_PS4_RightTrigger_Pull
	AOPS4_RightTrigger_Click           ActionOrigin = internal.EControllerActionOrigin_PS4_RightTrigger_Click
	AOPS4_LeftStick_Move               ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_Move
	AOPS4_LeftStick_Click              ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_Click
	AOPS4_LeftStick_DPadNorth          ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_DPadNorth
	AOPS4_LeftStick_DPadSouth          ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_DPadSouth
	AOPS4_LeftStick_DPadWest           ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_DPadWest
	AOPS4_LeftStick_DPadEast           ActionOrigin = internal.EControllerActionOrigin_PS4_LeftStick_DPadEast
	AOPS4_RightStick_Move              ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_Move
	AOPS4_RightStick_Click             ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_Click
	AOPS4_RightStick_DPadNorth         ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_DPadNorth
	AOPS4_RightStick_DPadSouth         ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_DPadSouth
	AOPS4_RightStick_DPadWest          ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_DPadWest
	AOPS4_RightStick_DPadEast          ActionOrigin = internal.EControllerActionOrigin_PS4_RightStick_DPadEast
	AOPS4_DPad_North                   ActionOrigin = internal.EControllerActionOrigin_PS4_DPad_North
	AOPS4_DPad_South                   ActionOrigin = internal.EControllerActionOrigin_PS4_DPad_South
	AOPS4_DPad_West                    ActionOrigin = internal.EControllerActionOrigin_PS4_DPad_West
	AOPS4_DPad_East                    ActionOrigin = internal.EControllerActionOrigin_PS4_DPad_East
	AOPS4_Gyro_Move                    ActionOrigin = internal.EControllerActionOrigin_PS4_Gyro_Move
	AOPS4_Gyro_Pitch                   ActionOrigin = internal.EControllerActionOrigin_PS4_Gyro_Pitch
	AOPS4_Gyro_Yaw                     ActionOrigin = internal.EControllerActionOrigin_PS4_Gyro_Yaw
	AOPS4_Gyro_Roll                    ActionOrigin = internal.EControllerActionOrigin_PS4_Gyro_Roll
	AOXBoxOne_A                        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_A
	AOXBoxOne_B                        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_B
	AOXBoxOne_X                        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_X
	AOXBoxOne_Y                        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_Y
	AOXBoxOne_LeftBumper               ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftBumper
	AOXBoxOne_RightBumper              ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightBumper
	AOXBoxOne_Menu                     ActionOrigin = internal.EControllerActionOrigin_XBoxOne_Menu
	AOXBoxOne_View                     ActionOrigin = internal.EControllerActionOrigin_XBoxOne_View
	AOXBoxOne_LeftTrigger_Pull         ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftTrigger_Pull
	AOXBoxOne_LeftTrigger_Click        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftTrigger_Click
	AOXBoxOne_RightTrigger_Pull        ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightTrigger_Pull
	AOXBoxOne_RightTrigger_Click       ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightTrigger_Click
	AOXBoxOne_LeftStick_Move           ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_Move
	AOXBoxOne_LeftStick_Click          ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_Click
	AOXBoxOne_LeftStick_DPadNorth      ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_DPadNorth
	AOXBoxOne_LeftStick_DPadSouth      ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_DPadSouth
	AOXBoxOne_LeftStick_DPadWest       ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_DPadWest
	AOXBoxOne_LeftStick_DPadEast       ActionOrigin = internal.EControllerActionOrigin_XBoxOne_LeftStick_DPadEast
	AOXBoxOne_RightStick_Move          ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_Move
	AOXBoxOne_RightStick_Click         ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_Click
	AOXBoxOne_RightStick_DPadNorth     ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_DPadNorth
	AOXBoxOne_RightStick_DPadSouth     ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_DPadSouth
	AOXBoxOne_RightStick_DPadWest      ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_DPadWest
	AOXBoxOne_RightStick_DPadEast      ActionOrigin = internal.EControllerActionOrigin_XBoxOne_RightStick_DPadEast
	AOXBoxOne_DPad_North               ActionOrigin = internal.EControllerActionOrigin_XBoxOne_DPad_North
	AOXBoxOne_DPad_South               ActionOrigin = internal.EControllerActionOrigin_XBoxOne_DPad_South
	AOXBoxOne_DPad_West                ActionOrigin = internal.EControllerActionOrigin_XBoxOne_DPad_West
	AOXBoxOne_DPad_East                ActionOrigin = internal.EControllerActionOrigin_XBoxOne_DPad_East
	AOXBox360_A                        ActionOrigin = internal.EControllerActionOrigin_XBox360_A
	AOXBox360_B                        ActionOrigin = internal.EControllerActionOrigin_XBox360_B
	AOXBox360_X                        ActionOrigin = internal.EControllerActionOrigin_XBox360_X
	AOXBox360_Y                        ActionOrigin = internal.EControllerActionOrigin_XBox360_Y
	AOXBox360_LeftBumper               ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftBumper
	AOXBox360_RightBumper              ActionOrigin = internal.EControllerActionOrigin_XBox360_RightBumper
	AOXBox360_Start                    ActionOrigin = internal.EControllerActionOrigin_XBox360_Start
	AOXBox360_Back                     ActionOrigin = internal.EControllerActionOrigin_XBox360_Back
	AOXBox360_LeftTrigger_Pull         ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftTrigger_Pull
	AOXBox360_LeftTrigger_Click        ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftTrigger_Click
	AOXBox360_RightTrigger_Pull        ActionOrigin = internal.EControllerActionOrigin_XBox360_RightTrigger_Pull
	AOXBox360_RightTrigger_Click       ActionOrigin = internal.EControllerActionOrigin_XBox360_RightTrigger_Click
	AOXBox360_LeftStick_Move           ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_Move
	AOXBox360_LeftStick_Click          ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_Click
	AOXBox360_LeftStick_DPadNorth      ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_DPadNorth
	AOXBox360_LeftStick_DPadSouth      ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_DPadSouth
	AOXBox360_LeftStick_DPadWest       ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_DPadWest
	AOXBox360_LeftStick_DPadEast       ActionOrigin = internal.EControllerActionOrigin_XBox360_LeftStick_DPadEast
	AOXBox360_RightStick_Move          ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_Move
	AOXBox360_RightStick_Click         ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_Click
	AOXBox360_RightStick_DPadNorth     ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_DPadNorth
	AOXBox360_RightStick_DPadSouth     ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_DPadSouth
	AOXBox360_RightStick_DPadWest      ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_DPadWest
	AOXBox360_RightStick_DPadEast      ActionOrigin = internal.EControllerActionOrigin_XBox360_RightStick_DPadEast
	AOXBox360_DPad_North               ActionOrigin = internal.EControllerActionOrigin_XBox360_DPad_North
	AOXBox360_DPad_South               ActionOrigin = internal.EControllerActionOrigin_XBox360_DPad_South
	AOXBox360_DPad_West                ActionOrigin = internal.EControllerActionOrigin_XBox360_DPad_West
	AOXBox360_DPad_East                ActionOrigin = internal.EControllerActionOrigin_XBox360_DPad_East
	AOSteamV2_A                        ActionOrigin = internal.EControllerActionOrigin_SteamV2_A
	AOSteamV2_B                        ActionOrigin = internal.EControllerActionOrigin_SteamV2_B
	AOSteamV2_X                        ActionOrigin = internal.EControllerActionOrigin_SteamV2_X
	AOSteamV2_Y                        ActionOrigin = internal.EControllerActionOrigin_SteamV2_Y
	AOSteamV2_LeftBumper               ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftBumper
	AOSteamV2_RightBumper              ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightBumper
	AOSteamV2_LeftGrip                 ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftGrip
	AOSteamV2_RightGrip                ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightGrip
	AOSteamV2_LeftGrip_Upper           ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftGrip_Upper
	AOSteamV2_RightGrip_Upper          ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightGrip_Upper
	AOSteamV2_LeftBumper_Pressure      ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftBumper_Pressure
	AOSteamV2_RightBumper_Pressure     ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightBumper_Pressure
	AOSteamV2_LeftGrip_Pressure        ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftGrip_Pressure
	AOSteamV2_RightGrip_Pressure       ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightGrip_Pressure
	AOSteamV2_LeftGrip_Upper_Pressure  ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftGrip_Upper_Pressure
	AOSteamV2_RightGrip_Upper_Pressure ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightGrip_Upper_Pressure
	AOSteamV2_Start                    ActionOrigin = internal.EControllerActionOrigin_SteamV2_Start
	AOSteamV2_Back                     ActionOrigin = internal.EControllerActionOrigin_SteamV2_Back
	AOSteamV2_LeftPad_Touch            ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_Touch
	AOSteamV2_LeftPad_Swipe            ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_Swipe
	AOSteamV2_LeftPad_Click            ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_Click
	AOSteamV2_LeftPad_Pressure         ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_Pressure
	AOSteamV2_LeftPad_DPadNorth        ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_DPadNorth
	AOSteamV2_LeftPad_DPadSouth        ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_DPadSouth
	AOSteamV2_LeftPad_DPadWest         ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_DPadWest
	AOSteamV2_LeftPad_DPadEast         ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftPad_DPadEast
	AOSteamV2_RightPad_Touch           ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_Touch
	AOSteamV2_RightPad_Swipe           ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_Swipe
	AOSteamV2_RightPad_Click           ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_Click
	AOSteamV2_RightPad_Pressure        ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_Pressure
	AOSteamV2_RightPad_DPadNorth       ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_DPadNorth
	AOSteamV2_RightPad_DPadSouth       ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_DPadSouth
	AOSteamV2_RightPad_DPadWest        ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_DPadWest
	AOSteamV2_RightPad_DPadEast        ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightPad_DPadEast
	AOSteamV2_LeftTrigger_Pull         ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftTrigger_Pull
	AOSteamV2_LeftTrigger_Click        ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftTrigger_Click
	AOSteamV2_RightTrigger_Pull        ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightTrigger_Pull
	AOSteamV2_RightTrigger_Click       ActionOrigin = internal.EControllerActionOrigin_SteamV2_RightTrigger_Click
	AOSteamV2_LeftStick_Move           ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_Move
	AOSteamV2_LeftStick_Click          ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_Click
	AOSteamV2_LeftStick_DPadNorth      ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_DPadNorth
	AOSteamV2_LeftStick_DPadSouth      ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_DPadSouth
	AOSteamV2_LeftStick_DPadWest       ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_DPadWest
	AOSteamV2_LeftStick_DPadEast       ActionOrigin = internal.EControllerActionOrigin_SteamV2_LeftStick_DPadEast
	AOSteamV2_Gyro_Move                ActionOrigin = internal.EControllerActionOrigin_SteamV2_Gyro_Move
	AOSteamV2_Gyro_Pitch               ActionOrigin = internal.EControllerActionOrigin_SteamV2_Gyro_Pitch
	AOSteamV2_Gyro_Yaw                 ActionOrigin = internal.EControllerActionOrigin_SteamV2_Gyro_Yaw
	AOSteamV2_Gyro_Roll                ActionOrigin = internal.EControllerActionOrigin_SteamV2_Gyro_Roll
	AOCount                            ActionOrigin = internal.EControllerActionOrigin_Count
)

// GetGlyphForActionOrigin returns a local path to art for on-screen glyph for
// a particular origin. The returned path refers to a PNG file.
func GetGlyphForActionOrigin(origin ActionOrigin) string {
	defer cleanup()()

	return internal.GoString(internal.SteamAPI_ISteamController_GetGlyphForActionOrigin(origin))
}

// GetStringForActionOrigin returns a localized string (from Steam's language
// setting) for the specified origin.
func GetStringForActionOrigin(origin ActionOrigin) string {
	defer cleanup()()

	return internal.GoString(internal.SteamAPI_ISteamController_GetStringForActionOrigin(origin))
}
