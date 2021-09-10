package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
import "C"

// Label is only for convenience. Under the hood it's still a textfield with some preconfiguration in place
type Label struct {
	textFieldPtr C.TextFieldPtr
	callback     func()
}

var labels []*Label

// NewLabel - This func is not thread safe.
func NewLabel(x int, y int, width int, height int) *TextField {
	textField := NewTextField(x, y, width, height)
	textField.SetBezeled(false)
	textField.SetDrawsBackground(false)
	textField.SetEditable(false)
	textField.SetSelectable(false)
	return textField
}
