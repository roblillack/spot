package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"unsafe"
)

// Button represents a button control that can trigger actions.
type Button struct {
	buttonPtr C.ButtonPtr
	callback  func()
}

var buttons []*Button

//export onButtonClicked
func onButtonClicked(id C.int) {
	buttonId := int(id)
	if buttonId < len(buttons) && buttons[buttonId].callback != nil {
		buttons[buttonId].callback()
	}
}

// NewButton constructs a new button at position (x, y) and with size (width x height).
// Gotcha! It is currently hard-coded to quit the app when the button is being pressed, until
// callbacks have been implemented.
// This func is not thread safe.
func NewButton(x int, y int, width int, height int) *Button {
	buttonId := len(buttons)
	buttonPtr := C.Button_New(C.int(buttonId), C.int(x), C.int(y), C.int(width), C.int(height))

	btn := &Button{
		buttonPtr: buttonPtr,
	}
	buttons = append(buttons, btn)
	return btn
}

// SetTitle sets the title of the button, which is the text displayed on the button.
func (btn *Button) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(btn.buttonPtr, cTitle)
}

func (btn *Button) OnClick(fn func()) {
	btn.callback = fn
}
