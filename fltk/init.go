package fltk

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
)

func RunOnMainLoop(fn func()) {
	// goFltk.Lock()
	// fn()
	// goFltk.Unlock()
	// goFltk.AwakeNullMessage()

	goFltk.Awake(fn)
}

func init() {
	goFltk.Lock()
	spot.RunOnMainLoop = RunOnMainLoop
}
