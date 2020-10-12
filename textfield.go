package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
import "C"
import "unsafe"

// TextField -Button represents a button control that can trigger actions.
type TextField struct {
	textFieldPtr C.TextFieldPtr
	callback     func()
}

var textfields []*TextField

// NewTextField - This func is not thread safe.
func NewTextField(x int, y int, width int, height int) *TextField {
	textFieldID := len(textfields)
	textFieldPtr := C.TextField_New(C.int(textFieldID), C.int(x), C.int(y), C.int(width), C.int(height))

	tf := &TextField{
		textFieldPtr: textFieldPtr,
	}
	textfields = append(textfields, tf)
	return tf
}

// StringValue - returns the string value of the text field
func (textField *TextField) StringValue() string {
	return C.GoString(C.TextField_StringValue(textField.textFieldPtr))
}

// SetStringValue sets the string value of the text field
func (textField *TextField) SetStringValue(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.TextField_SetStringValue(textField.textFieldPtr, cText)
}
