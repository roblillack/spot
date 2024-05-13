//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"journey/spot"
	"runtime"

	goFltk "github.com/pwiecz/go-fltk"
)

func RunOnMainLoop(fn func()) {
	// goFltk.Lock()
	// fn()
	// goFltk.Unlock()
	// goFltk.AwakeNullMessage()

	goFltk.Awake(fn)
}

func Init() {
	runtime.LockOSThread()
	goFltk.Lock()
	spot.RunOnMainLoop = RunOnMainLoop
}

func Run() {
	goFltk.Run()
}
