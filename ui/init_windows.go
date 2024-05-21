//go:build !fltk && windows

package ui

import (
	"runtime"

	"github.com/roblillack/spot"
)

// BackendName is the name of the backend. It can be used to check which backend
// is currently in use by the application during runtime.
const BackendName = "windows"

var activeWindow *Window

func runOnMainLoop(fn func()) {
	if activeWindow == nil {
		fn()
	}

	activeWindow.ref.RunUiThread(fn)
}

// Init initializes the UI library for the FLTK backend. It locks the OS thread
// and sets up Spot to be able to intercept the main loop.
func Init() {
	runtime.LockOSThread()
	spot.RunOnMainLoop = runOnMainLoop
}

// Run starts the main loop for the FLTK backend.
func Run() {
	if activeWindow == nil {
		return
	}

	ref := activeWindow.ref
	if ref == nil {
		return
	}

	ref.RunAsMain()
}
