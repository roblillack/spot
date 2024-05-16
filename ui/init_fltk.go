//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"runtime"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

// BackendName is the name of the backend. It can be used to check which backend
// is currently in use by the application during runtime.
const BackendName = "fltk"

func runOnMainLoop(fn func()) {
	goFltk.Awake(fn)
}

// Init initializes the UI library for the FLTK backend. It locks the OS thread
// and sets up Spot to be able to intercept the main loop.
func Init() {
	runtime.LockOSThread()
	goFltk.Lock()
	spot.RunOnMainLoop = runOnMainLoop
}

// Run starts the main loop for the FLTK backend.
func Run() {
	goFltk.Run()
}
