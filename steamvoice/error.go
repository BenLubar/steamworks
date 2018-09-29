package steamvoice

import (
	"errors"
	"io"

	"github.com/BenLubar/steamworks/internal"
)

// Errors returned by this package.
//
// Additional errors include io.EOF and io.ErrShortBuffer.
var (
	ErrNotInitialized = errors.New("steamworks/steamvoice: interface has not been initialized")
	ErrDataCorrupted  = errors.New("steamworks/steamvoice: voice data has been corrupted")
	ErrRestricted     = errors.New("steamworks/steamvoice: user is chat restricted")
	ErrUnknown        = errors.New("steamworks/steamvoice: unknown error")
)

func toError(result internal.EVoiceResult) error {
	switch result {
	case internal.EVoiceResult_OK:
		return nil
	case internal.EVoiceResult_NotInitialized:
		return ErrNotInitialized
	case internal.EVoiceResult_NotRecording:
		return io.EOF
	case internal.EVoiceResult_NoData:
		return nil
	case internal.EVoiceResult_BufferTooSmall:
		return io.ErrShortBuffer
	case internal.EVoiceResult_DataCorrupted:
		return ErrDataCorrupted
	case internal.EVoiceResult_Restricted:
		return ErrRestricted
	default:
		return ErrUnknown
	}
}
