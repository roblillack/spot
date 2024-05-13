//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"runtime"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
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
