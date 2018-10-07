package internal

import (
	"sync"
	"sync/atomic"
)

var globalLock sync.Mutex
var onces []*Once
var onShutdown []func()

// Once is a modified version of sync.Once that allows resetting under certain
// circumstances.
type Once struct {
	done int32
}

// Do runs f if and only if Do has not been called on this Once since process
// startup or the last call to ResetOnce.
//
// If multiple calls to Do the same Once are made, the first will run the
// function and the others will wait until it has completed.
func (o *Once) Do(f func()) {
	if atomic.LoadInt32(&o.done) == 1 {
		return
	}

	globalLock.Lock()
	defer globalLock.Unlock()

	if atomic.LoadInt32(&o.done) == 1 {
		return
	}

	onces = append(onces, o)
	defer atomic.StoreInt32(&o.done, 1)

	f()
}

// OnShutdown registers f to be run on the next call to ResetOnce.
//
// This function is only safe to call from the body of a Once.Do function.
func OnShutdown(f func()) {
	onShutdown = append(onShutdown, f)
}

// ResetOnce resets all internal.Once instances to their initial state.
func ResetOnce() {
	globalLock.Lock()
	defer globalLock.Unlock()

	for _, o := range onces {
		atomic.StoreInt32(&o.done, 0)
	}

	onces = nil

	for _, f := range onShutdown {
		f()
	}

	onShutdown = nil
}
