package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import "unsafe"

// Window represents a ... window
type Window struct {
	winPtr unsafe.Pointer
}

// NewWindow constructs a new window
func NewWindow(title string, x int, y int, w int, h int) *Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return &Window{winPtr: C.Window_New(C.int(x), C.int(y), C.int(w), C.int(h), cTitle)}
}

func (wnd *Window) MakeKeyAndOrderFront() {
	C.Window_MakeKeyAndOrderFront(wnd.winPtr)
}

func (wnd *Window) AddButton(btn *Button) {
	C.Window_AddButton(wnd.winPtr, btn.buttonPtr)
}
