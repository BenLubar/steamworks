package steamcontroller

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// ActionSetHandle is used to refer to specific in-game actions or action sets.
type ActionSetHandle = internal.ControllerActionSetHandle

// ActivateActionSet reconfigures the controller to use the specified action
// set (ie "Menu", "Walk", or "Drive").
//
// This is cheap, and can be safely called repeatedly. It's often easier to
// repeatedly call it in your state loops, instead of trying to place it in all
// of your state transitions.
//
// Example:
//
//    func updateStateLoop(current steamcontroller.Handle) {
//        switch currentState {
//        case MENU:
//            steamcontroller.ActivateActionSet(current, menuSetHandle)
//            doMenuStuff()
//        case WALKING:
//            steamcontroller.ActivateActionSet(current, walkingSetHandle)
//            doWalkingStuff()
//        case DRIVING:
//            steamcontroller.ActivateActionSet(current, drivingSetHandle)
//            doDrivingStuff()
//        case FIGHTING:
//            steamcontroller.ActivateActionSet(current, fightingSetHandle)
//            doFightingStuff()
//        }
//    }
func ActivateActionSet(controller Handle, actionSet ActionSetHandle) {
	defer cleanup()()

	internal.SteamAPI_ISteamController_ActivateActionSet(controller, actionSet)
}

// GetActionSetHandle looks up the handle for an Action Set. Best to do this
// once on startup, and store the handles for all future API calls.
//
// The name refers to an identifier in the game's VDF file.
func GetActionSetHandle(name string) ActionSetHandle {
	defer cleanup()()

	cname := internal.CString(name)
	defer internal.Free(unsafe.Pointer(cname))

	return internal.SteamAPI_ISteamController_GetActionSetHandle(cname)
}

// GetCurrentActionSet returns the current action set for the specified
// controller.
func GetCurrentActionSet(controller Handle) ActionSetHandle {
	defer cleanup()()

	return internal.SteamAPI_ISteamController_GetCurrentActionSet(controller)
}
