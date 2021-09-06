package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
import "C"
import (
	"fmt"
	"unsafe"
)

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

func (textField *TextField) Enabled() bool {
	return C.TextField_Enabled(textField.textFieldPtr) == 1
}

// SetEnabled sets the enabled value of the text field. CANNOT BE CHANGED AT RUNTIME
func (textField *TextField) SetEnabled(enabled bool) {
	if enabled {
		C.TextField_SetEnabled(textField.textFieldPtr, 1)
	} else {
		C.TextField_SetEnabled(textField.textFieldPtr, 0)
	}
}

func (textField *TextField) Editable() bool {
	return C.TextField_Editable(textField.textFieldPtr) == 1
}

func (textField *TextField) SetEditable(editable bool) {
	if editable {
		C.TextField_SetEditable(textField.textFieldPtr, 1)
	} else {
		C.TextField_SetEditable(textField.textFieldPtr, 0)
	}
}

func (textField *TextField) SetFontFamily(fontFamily string) {
	cText := C.CString(fontFamily)
	defer C.free(unsafe.Pointer(cText))
	C.TextField_SetFontFamily(textField.textFieldPtr, cText)
}

func (textField *TextField) SetFontSize(fontSize int) {
	C.TextField_SetFontSize(textField.textFieldPtr, C.int(fontSize))
}

func (textField *TextField) SetColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBackgroundColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetBackgroundColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetBorderColor(textField.textFieldPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderWidth(borderWidth int) {
	C.TextField_SetBorderWidth(textField.textFieldPtr, C.int(borderWidth))
}
