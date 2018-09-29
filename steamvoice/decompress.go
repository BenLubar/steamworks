package steamvoice

import (
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// OptimalSampleRate returns the native sample rate of the Steam voice decoder.
//
// Using this sample rate for DecompressVoice will perform the least CPU
// processing. However, the final audio quality will depend on how well the
// audio device (and/or your application's audio output SDK) deals with lower
// sample rates. You may find that you get the best audio output quality when
// you ignore this function and use the native sample rate of your audio output
// device, which is usually 48000 or 44100.
func OptimalSampleRate() uint32 {
	return internal.SteamAPI_ISteamUser_GetVoiceOptimalSampleRate()
}

// DecompressVoice decodes the compressed voice data returned by GetVoice.
//
// The output data is raw single-channel 16-bit PCM audio. The decoder supports
// any sample rate from 11025 to 48000. See OptimalSampleRate for more
// information.
func DecompressVoice(compressed []byte, sampleRate uint32) ([]uint16, error) {
	buffer := make([]uint16, 10<<10)

	for {
		var bytesWritten uint32
		result := internal.SteamAPI_ISteamUser_DecompressVoice(unsafe.Pointer(&compressed[0]), uint32(len(compressed)), unsafe.Pointer(&buffer[0]), uint32(len(buffer))*2, &bytesWritten, sampleRate)

		if result == internal.EVoiceResult_BufferTooSmall {
			buffer = make([]uint16, len(buffer)*2)
			continue
		}

		return buffer[:bytesWritten/2], toError(result)
	}
}
