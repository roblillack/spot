package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
import "C"

// TextField -Button represents a button control that can trigger actions.
type TextField struct {
	textField C.pTextField
	callback  func()
}

var textfields []*TextField

// NewTextField - This func is not thread safe.
func NewTextField(x int, y int, width int, height int) *TextField {
	textFieldID := len(textfields)
	textField := C.TextField_New(C.int(textFieldID), C.int(x), C.int(y), C.int(width), C.int(height))

	tf := &TextField{
		textField: textField,
	}
	textfields = append(textfields, tf)
	return tf
}
