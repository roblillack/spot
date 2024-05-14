//go:build !fltk && (darwin || cocoa)

package ui

import (
	"runtime"

	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

// BackendName is the name of the backend. It can be used to check which backend
// is currently in use by the application during runtime.
const BackendName = "cocoa"

// Init initializes the UI library for the Cocoa backend. It locks the OS thread
// and sets up Spot to be able to intercept the main loop.
func Init() {
	spot.RunOnMainLoop = gocoa.RunOnMainLoop

	runtime.LockOSThread()
	gocoa.InitApplication()
}

// Run starts the main loop for the Cocoa backend.
func Run() {
	gocoa.RunApplication()
}
