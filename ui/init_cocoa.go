//go:build !fltk && (darwin || cocoa)

package ui

import (
	"runtime"

	"github.com/mojbro/gocoa"
	"github.com/roblillack/spot"
)

func Init() {
	spot.RunOnMainLoop = gocoa.RunOnMainLoop

	runtime.LockOSThread()
	gocoa.InitApplication()
}

func Run() {
	gocoa.RunApplication()
}
