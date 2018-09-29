// +build windows linux darwin
// +build 386 amd64

// Package internal wraps the Steamworks API.
package internal

/*
#include "api.gen.h"
#include "callback.h"
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"strconv"
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

//export warningMessageHook
func warningMessageHook(severity C.int, debugText *C.char) {
	msg := C.GoString(debugText)
	switch severity {
	case 0:
		OnDebugMessage(msg)
	case 1:
		OnWarningMessage(msg)
	default:
		panic("steamworks: unexpected message level " + strconv.FormatInt(int64(severity), 10) + ": " + msg)
	}
}

// Message hook stubs (overwritten by steamutils)
var (
	OnDebugMessage   = func(string) {}
	OnWarningMessage = func(string) {}
)

// SetWarningMessageHook sets the C function pointer to make OnDebugMessage and
// OnWarningMessage work.
func SetWarningMessageHook() {
	C.SetWarningMessageHookGo()
}

// Helpful C functions for other packages to use internally:

// Malloc wraps C.malloc.
func Malloc(size uintptr) unsafe.Pointer { return C.malloc(C.size_t(size)) }

// Free wraps C.free.
func Free(ptr unsafe.Pointer) { C.free(ptr) }

// CChar is a C char.
type CChar = C.char

// CString wraps C.CString.
func CString(str string) *C.char { return C.CString(str) }

// GoString wraps C.GoString.
func GoString(str *C.char) string { return C.GoString(str) }

// GoStringN wraps C.GoStringN.
func GoStringN(str *C.char, maxSize uintptr) string { return C.GoStringN(str, C.int(maxSize)) }
