// Package steamworks wraps the Steamworks API.
//
// See the official API documentation on the Steam Partner website:
// <https://partner.steamgames.com/doc/api>
//
// This package attempts to make API functions available in a more
// idiomatic Go style.
package steamworks

import (
	"encoding/binary"
	"errors"
	"net"
	"runtime"
	"time"
	"unsafe"

	"github.com/BenLubar/steamworks/internal"
)

// Registration is an opaque type that represents a registered callback.
type Registration interface {
	// Unregister unregisters the callback. It is unsafe to call Unregister
	// multiple times on the same Registration.
	Unregister()
}

var callbackShutdown chan<- chan<- struct{}

// RestartAppIfNecessary checks if your executable was launched through Steam
// and relaunches it through Steam if it wasn't started by Steam originally.
//
// See Initialization and Shutdown for additional information.
// <https://partner.steamgames.com/doc/sdk/api#initialization_and_shutdown>
//
// If this returns true then it starts the Steam client if required and
// launches your game again through it, and you should quit your process as
// soon as possible. This effectively runs `steam://run/<AppId>` so it may not
// relaunch the exact executable that called it, as it will always relaunch
// from the version installed in your Steam library folder.
//
// If it returns false, then your game was launched by the Steam client and no
// action needs to be taken. One exception is if a steam_appid.txt file is
// present then this will return false regardless. This allows you to develop
// and test without launching your game through the Steam client. Make sure to
// remove the steam_appid.txt file when uploading the game to your Steam depot!
//
// Example:
//
//    const appId = ...
//
//    func main() {
//        if steamworks.RestartAppIfNecessary(appId) {
//            return
//        }
//
//        if err := steamworks.InitClient(true); err != nil {
//            handleFatalError(err)
//        }
//        defer steamworks.Shutdown()
//
//        mainGameLoop()
//    }
func RestartAppIfNecessary(ownAppID AppID) bool {
	defer internal.Cleanup()()

	return internal.SteamAPI_RestartAppIfNecessary(uint32(ownAppID))
}

// Errors that can be returned by InitClient or InitServer.
var (
	ErrSteamNotRunning = errors.New("steamworks: the Steam client is not running")
	ErrIPv4Only        = errors.New("steamworks: only IPv4 addresses are supported")
	errSteamInitFailed = errors.New("steamworks: failed to initialize") // TODO: be more specific
)

// InitClient initializes the Steamworks API for game clients.
//
// See Initialization and Shutdown for additional information.
// <https://partner.steamgames.com/doc/sdk/api#initialization_and_shutdown>
//
// If startCallbackGoroutine is true, RunCallbacks will automatically be called
// in a loop. Set startCallbackGoroutine to false if you plan to call
// RunCallbacks manually, e.g. if your callback code is not thread-safe.
//
// Returns nil if all required interfaces have been acquired and are accessible.
//
// A non-nil error indicates one of the following conditions:
//
//    - The Steam client isn't running. A running Steam client is required to
//      provide implementations of the various Steamworks interfaces.
//    - The Steam client couldn't determine the App ID of game. If you're
//      running your application from the executable or debugger directly then
//      you must have a steam_appid.txt in your game directory next to the
//      executable, with your app ID in it and nothing else. Steam will look
//      for this file in the current working directory. If you are running your
//      executable from a different directory you may need to relocate the
//      steam_appid.txt file.
//    - Your application is not running under the same OS user context as the
//      Steam client, such as a different user or administration access level.
//    - Ensure that you own a license for the App ID on the currently active
//      Steam account. Your game must show up in your Steam library.
//    - Your App ID is not completely set up, i.e. in Release State:
//      Unavailable, or it's missing default packages.
func InitClient(startCallbackGoroutine bool) error {
	if !internal.SteamAPI_Init() {
		if !internal.SteamAPI_IsSteamRunning() {
			return ErrSteamNotRunning
		}
		return errSteamInitFailed
	}

	stopCallbackGoroutine()

	if startCallbackGoroutine {
		ch := make(chan chan<- struct{}, 1)
		callbackShutdown = ch
		go runCallbacksForever(ch)
	}

	return nil
}

// ServerMode is the authentication mode for a Steam game server.
type ServerMode = internal.EServerMode

const (
	// NoAuthentication doesn't authenticate user logins and doesn't list the
	// server on the master server list.
	NoAuthentication ServerMode = internal.EServerMode_NoAuthentication
	// Authentication authenticates users and lists the server on the master
	// server list but doesn't run VAC on clients that connect.
	Authentication ServerMode = internal.EServerMode_Authentication
	// AuthenticationAndSecure authenticates users, lists the server on the
	// master server list, and VAC protects clients.
	AuthenticationAndSecure ServerMode = internal.EServerMode_AuthenticationAndSecure
)

// UseGameSocketShare is a placeholder value for queryPort that means the game
// is responsible for sending and receiving UDP packets for the master server
// updater.
const UseGameSocketShare = ^uint16(0)

// InitServer initializes the ISteamGameServer interface object, and sets
// server properties which may not be changed.
//
// After calling this function, you should set any additional server parameters,
// and then call steamgameserver.LogOnAnonymous or steamgameserver.LogOn.
//
// If you pass in UseGameSocketShare into queryPort, then the game server will
// use GameSocketShare mode, which means that the game is responsible for
// sending and receiving UDP packets for the master server updater.
func InitServer(ip net.IP, steamPort, gamePort, queryPort uint16, serverMode ServerMode, version string, startCallbackGoroutine bool) error {
	cversion := internal.CString(version)
	defer internal.Free(unsafe.Pointer(cversion))

	var ipInt uint32
	if ip4 := ip.To4(); ip4 != nil {
		ipInt = binary.BigEndian.Uint32(ip4)
	} else if len(ip) != 0 {
		return ErrIPv4Only
	}

	if !internal.SteamGameServer_Init(ipInt, steamPort, gamePort, queryPort, serverMode, cversion) {
		return errSteamInitFailed
	}

	stopCallbackGoroutine()

	if startCallbackGoroutine {
		ch := make(chan chan<- struct{}, 1)
		callbackShutdown = ch
		go runCallbacksForever(ch)
	}

	return nil
}

// Shutdown shuts down the Steamworks API, releases pointers and frees memory.
//
// You should call this during process shutdown if possible.
//
// This will not unhook the Steam Overlay from your game as there's no
// guarantee that your rendering API is done using it.
func Shutdown() {
	stopCallbackGoroutine()

	internal.ResetOnce()

	internal.SteamAPI_Shutdown()
}

func stopCallbackGoroutine() {
	if callbackShutdown == nil {
		return
	}

	ch := make(chan struct{}, 1)
	callbackShutdown <- ch
	<-ch

	callbackShutdown = nil
}

func runCallbacksForever(quit <-chan chan<- struct{}) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		internal.SteamAPI_RunCallbacks()

		select {
		case ch := <-quit:
			ch <- struct{}{}
			return
		case <-ticker.C:
		}
	}
}

// RunCallbacks dispatches callbacks and call results to all of the registered
// listeners.
//
// It's best to call this at >10Hz, the more time between calls, the more
// potential latency between receiving events or results from the Steamworks
// API. Most games call this once per render-frame. All registered listener
// functions will be invoked during this call, in the caller's thread context.
//
// RunCallbacks is safe to call from multiple goroutines simultaneously, but if
// you choose to do this, callback code could be executed on any goroutine. One
// alternative is to call RunCallbacks from the main thread only.
//
// Calling this function is required if and only if InitClient or InitServer
// was called with startCallbackGoroutine set to false.
func RunCallbacks() {
	internal.SteamAPI_RunCallbacks()
}
