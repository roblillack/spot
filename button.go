package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "button.h"
import "C"
import "unsafe"

// Button represents a button control that can trigger actions
type Button struct {
	buttonPtr unsafe.Pointer
}

// NewButton constructs a new button at position (x, y) and with size (width x height).
// Gotcha! It is currently hard-coded to quit the app when the button is being pressed, until
// callbacks have been implemented.
func NewButton(x int, y int, width int, height int) *Button {
	return &Button{buttonPtr: C.Button_New(C.int(x), C.int(y), C.int(width), C.int(height))}
}

// SetTitle sets the title of the button, which is the text displayed on the button.
func (btn *Button) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(btn.buttonPtr, cTitle)
}
