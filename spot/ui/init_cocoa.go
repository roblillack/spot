//go:build !fltk && (darwin || cocoa)

package ui

import (
	"journey/spot"
	"runtime"

	"github.com/mojbro/gocoa"
)

func Init() {
	spot.RunOnMainLoop = gocoa.RunOnMainLoop

	runtime.LockOSThread()
	gocoa.InitApplication()
}

func Run() {
	gocoa.RunApplication()
}
