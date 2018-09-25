package internal

/*
#include "api.gen.h"
#include "callback.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

var (
	callbackLock sync.Mutex
	callbacks    = make(map[C.CallbackID_t]func(unsafe.Pointer, uintptr, bool, SteamAPICall))
)

// Cleanup should be called as follows:
//
//    defer internal.Cleanup()()
//
// It locks the current OS thread and releases Steam API thread-local memory in the returned func.
func Cleanup() func() {
	runtime.LockOSThread()

	return func() {
		SteamAPI_ReleaseCurrentThreadMemory()
		runtime.UnlockOSThread()
	}
}

//export onCallback
func onCallback(cbid C.CallbackID_t, data unsafe.Pointer, dataLength uintptr, ioFailure bool, apiCallID SteamAPICall) {
	callbackLock.Lock()
	cb := callbacks[cbid]
	callbackLock.Unlock()

	if cb != nil {
		cb(data, dataLength, ioFailure, apiCallID)
	}
}

func registerCallback(cb func(unsafe.Pointer, uintptr, bool, SteamAPICall), size uintptr, callbackType int32, apiCallID SteamAPICall, gameServer bool) registeredCallback {
	cbid := C.Register_Callback(C.size_t(size), C.int(callbackType), apiCallID, C.bool(gameServer))

	callbackLock.Lock()
	callbacks[cbid] = cb
	callbackLock.Unlock()

	return registeredCallback(cbid)
}

type registeredCallback C.CallbackID_t

func (r registeredCallback) Unregister() {
	cbid := C.CallbackID_t(r)

	callbackLock.Lock()
	delete(callbacks, cbid)
	callbackLock.Unlock()

	C.Unregister_Callback(cbid)
}
