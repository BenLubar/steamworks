// Package steamvoice wraps Steam's voice chat API.
//
// This package is only available on clients.
//
// See the Steam Voice documentation for more details.
// <https://partner.steamgames.com/doc/features/voice>
package steamvoice

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// StartRecording starts voice recording.
//
// Once started, use Reader to get the data, and then call StopRecording when
// the user has released their push-to-talk hotkey or the game session has
// completed.
func StartRecording() {
	internal.SteamAPI_ISteamUser_StartVoiceRecording()
}

// StopRecording stops voice recording.
//
// Because people often release push-to-talk keys early, the system will keep
// recording for a little bit after this function is called. As such, Reader
// should continue to be read from until it returns io.EOF. Only then will
// voice recording be stopped.
func StopRecording() {
	internal.SteamAPI_ISteamUser_StopVoiceRecording()
}

// SetInGameSpeaking lets Steam know that the user is currently using voice
// chat in game.
//
// This will suppress the microphone for all voice communication in the Steam
// UI.
func SetInGameSpeaking(speaking bool) {
	internal.SteamAPI_ISteamFriends_SetInGameVoiceSpeaking(0, speaking)
}

// Reader is a stream of captured audio data from the microphone buffer.
//
// The compressed data can be transmitted by your application and decoded back
// into raw audio data using DecompressVoice on the other side. The compressed
// data provided is in an arbitrary format and is not meant to be played
// directly.
//
// This should be read from once per frame, and at worst no more than four
// times a second to keep the microphone input delay as low as possible.
// Reading any less frequently may result in gaps in the returned stream.
//
// It is recommended that you pass in an 8 kilobytes or larger destination
// buffer for compressed audio. Static buffers are recommended for performance
// reasons. However, if you would like to allocate precisely the right amount
// of space for a buffer before each call you may use Reader.Available() to find
// out how much data is available to be read.
var Reader VoiceReader

// VoiceReader is the type of Reader. It contains no state and should not be
// used directly. Instead, use the steamvoice.Reader variable.
type VoiceReader struct{}

// Read implements io.Reader. Reading is non-blocking, and if no data is
// available, (0, nil) will be returned.
func (VoiceReader) Read(p []byte) (int, error) {
	defer internal.Cleanup()()

	var bytesWritten uint32
	result := internal.SteamAPI_ISteamUser_GetVoice(true, unsafe.Pointer(&p[0]), uint32(len(p)), &bytesWritten, false, nil, 0, nil, 0)
	return int(bytesWritten), toError(result)
}

// Available returns the number of bytes of compressed voice data currently
// available from Read.
func (VoiceReader) Available() (int, error) {
	defer internal.Cleanup()()

	var bytesAvailable uint32
	result := internal.SteamAPI_ISteamUser_GetAvailableVoice(&bytesAvailable, nil, 0)
	return int(bytesAvailable), toError(result)
}
