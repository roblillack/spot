package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "button.h"
import "C"
import "unsafe"

type Button struct {
	buttonPtr unsafe.Pointer
}

// NewButton constructs a new button
func NewButton() *Button {
	return &Button{buttonPtr: C.Button_New()}
}

func (btn *Button) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(btn.buttonPtr, cTitle)
}
