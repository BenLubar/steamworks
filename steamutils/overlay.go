package steamutils

import (
	"sync"
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// OverlayNeedsPresent checks if the Overlay needs a present.
// Only required if using event driven render updates.
//
// Typically this call is unneeded if your game has a constantly running frame
// loop that calls the D3D Present API, or OGL SwapBuffers API every frame as
// is the case in most games. However, if you have a game that only refreshes
// the screen on an event driven basis then that can break the overlay, as it
// uses your Present/SwapBuffers calls to drive it's internal frame loop and it
// may also need to Present() to the screen any time a notification happens or
// when the overlay is brought up over the game by a user.
//
// You can use this API to ask the overlay if it currently need a present in
// that case, and then you can check for this periodically (roughly 33hz is
// desirable) and make sure you refresh the screen with Present or SwapBuffers
// to allow the overlay to do its work.
func OverlayNeedsPresent() bool {
	return internal.SteamAPI_ISteamUtils_BOverlayNeedsPresent()
}

// IsOverlayEnabled checks if the Steam Overlay is running and the user can
// access it.
//
// The overlay process could take a few seconds to start and hook the game
// process, so this function will initially return false while the overlay
// is loading.
func IsOverlayEnabled() bool {
	return internal.SteamAPI_ISteamUtils_IsOverlayEnabled()
}

// IsSteamInBigPictureMode checks if Steam and the Steam Overlay are running in
// Big Picture mode.
//
// Games must be launched through the Steam client to enable the Big Picture
// overlay. During development, a game can be added as a non-steam game to the
// developers library to test this feature.
//
// This will always return false if your app is not the 'game' application type.
func IsSteamInBigPictureMode() bool {
	return internal.SteamAPI_ISteamUtils_IsSteamInBigPictureMode()
}

var gamepadTextLock sync.Mutex

// GamepadTextInput activates the Big Picture text input dialog which only
// supports gamepad input.
//
// This function will return immediately if the big picture overlay is not
// available. Otherwise, it waits for the user to close the text input.
//
// If the user closes the text input by confirming it, the text they entered
// and true are returned.
//
// In any other case, the value of the existingText parameter and false are
// returned.
//
// Because this function blocks, the steam callback loop and your game's render
// loop must be called in another goroutine.
func GamepadTextInput(password, multiLine bool, description string, maxLength uint32, existingText string) (string, bool) {
	gamepadTextLock.Lock()
	defer gamepadTextLock.Unlock()

	defer internal.Cleanup()()

	inputMode := internal.EGamepadTextInputMode_Normal
	if password {
		inputMode = internal.EGamepadTextInputMode_Password
	}
	lineInputMode := internal.EGamepadTextInputLineMode_SingleLine
	if multiLine {
		lineInputMode = internal.EGamepadTextInputLineMode_MultipleLines
	}

	cdescription := internal.CString(description)
	defer internal.Free(unsafe.Pointer(cdescription))
	cexistingText := internal.CString(existingText)
	defer internal.Free(unsafe.Pointer(cexistingText))

	ch := make(chan *string, 1)

	registration := internal.RegisterCallback_GamepadTextInputDismissed(func(data *internal.GamepadTextInputDismissed, _ bool) {
		if !data.BSubmitted {
			ch <- nil
			return
		}

		ctextBuf := internal.Malloc(uintptr(data.UnSubmittedText) + 1)
		defer internal.Free(ctextBuf)
		ctext := (*internal.CChar)(ctextBuf)
		if !internal.SteamAPI_ISteamUtils_GetEnteredGamepadTextInput(ctext, uint32(data.UnSubmittedText)) {
			ch <- nil
			return
		}
		text := internal.GoStringN(ctext, uintptr(data.UnSubmittedText))
		ch <- &text
	}, 0)
	defer registration.Unregister()

	if !internal.SteamAPI_ISteamUtils_ShowGamepadTextInput(inputMode, lineInputMode, cdescription, maxLength, cexistingText) {
		return existingText, false
	}

	text := <-ch
	if text == nil {
		return existingText, false
	}
	return *text, true
}

// SetOverlayNotificationInset sets the inset of the overlay notification from
// the corner specified by SetOverlayNotificationPosition.
//
// A value of (0, 0) resets the position into the corner.
//
// This position is per-game and is reset each launch.
func SetOverlayNotificationInset(horizontal, vertical int) {
	defer internal.Cleanup()()

	internal.SteamAPI_ISteamUtils_SetOverlayNotificationInset(int32(horizontal), int32(vertical))
}

// SetOverlayNotificationPosition sets which corner the Steam overlay
// notification popup should display itself in.
//
// You can also set the distance from the specified corner by using
// SetOverlayNotificationInset.
//
// This position is per-game and is reset each launch.
func SetOverlayNotificationPosition(left, top bool) {
	defer internal.Cleanup()()

	switch {
	case left && top:
		internal.SteamAPI_ISteamUtils_SetOverlayNotificationPosition(internal.ENotificationPosition_EPositionTopLeft)
	case left:
		internal.SteamAPI_ISteamUtils_SetOverlayNotificationPosition(internal.ENotificationPosition_EPositionBottomLeft)
	case top:
		internal.SteamAPI_ISteamUtils_SetOverlayNotificationPosition(internal.ENotificationPosition_EPositionTopRight)
	default:
		internal.SteamAPI_ISteamUtils_SetOverlayNotificationPosition(internal.ENotificationPosition_EPositionBottomRight)
	}
}
