package steamutils

import (
	"sync"

	"github.com/BenLubar/steamworks"
	"github.com/BenLubar/steamworks/internal"
)

func init() {
	internal.OnDebugMessage = onMessage(&debugMessageHooks)
	internal.OnWarningMessage = onMessage(&warningMessageHooks)
}

var messageHookLock sync.Mutex
var debugMessageHooks []func(string)
var warningMessageHooks []func(string)

func onMessage(phooks *[]func(string)) func(string) {
	return func(msg string) {
		messageHookLock.Lock()
		hooks := *phooks
		messageHookLock.Unlock()

		for _, h := range hooks {
			if h != nil {
				h(msg)
			}
		}
	}
}

type hookRegistration struct {
	phooks *[]func(string)
	index  int
}

func (hr hookRegistration) Unregister() {
	messageHookLock.Lock()
	// Don't modify the slice directly as the underlying array might be in use
	// by onMessage.
	hooks := make([]func(string), len(*hr.phooks), cap(*hr.phooks))
	copy(hooks, *hr.phooks)

	hooks[hr.index] = nil

	// Reclaim space, but only from the end.
	for len(hooks) > 0 && hooks[len(hooks)-1] == nil {
		hooks = hooks[:len(hooks)-1]
	}

	*hr.phooks = hooks
	messageHookLock.Unlock()
}

func registerMessageHook(phooks *[]func(string), f func(string)) hookRegistration {
	initOnce.Do(doInit)

	messageHookLock.Lock()
	index := len(*phooks)
	*phooks = append(*phooks, f)
	messageHookLock.Unlock()

	return hookRegistration{
		phooks: phooks,
		index:  index,
	}
}

// RegisterDebugMessageHook registers a function to be called when Steam
// produces a debug message. This will only happen if Steam is started with
// -debug_steamapi.
func RegisterDebugMessageHook(f func(string)) steamworks.Registration {
	return registerMessageHook(&debugMessageHooks, f)
}

// RegisterWarningMessageHook registers a function to be called when Steam
// produces a warning message.
func RegisterWarningMessageHook(f func(string)) steamworks.Registration {
	return registerMessageHook(&warningMessageHooks, f)
}

var initOnce sync.Once

func doInit() {
	internal.SetWarningMessageHook()
}
