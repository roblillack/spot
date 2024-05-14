//go:build !fltk && (darwin || cocoa)

package ui

import (
	"runtime"

	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

const BackendName = "cocoa"

func Init() {
	spot.RunOnMainLoop = gocoa.RunOnMainLoop

	runtime.LockOSThread()
	gocoa.InitApplication()
}

func Run() {
	gocoa.RunApplication()
}
