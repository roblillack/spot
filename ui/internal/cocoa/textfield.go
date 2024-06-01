package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textfield.h"
// #import "view.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// TextField -Button represents a button control that can trigger actions.
type TextField struct {
	ptr      C.TextFieldPtr
	callback func(value string)
}

var textfields []*TextField

// NewTextField - This func is not thread safe.
func NewTextField(x int, y int, width int, height int) *TextField {
	textFieldID := len(textfields)
	textFieldPtr := C.TextField_New(C.int(textFieldID), C.int(x), C.int(y), C.int(width), C.int(height))

	tf := &TextField{
		ptr: textFieldPtr,
	}
	textfields = append(textfields, tf)
	return tf
}

// StringValue - returns the string value of the text field
func (textField *TextField) StringValue() string {
	return C.GoString(C.TextField_StringValue(textField.ptr))
}

// SetStringValue sets the string value of the text field
func (textField *TextField) SetStringValue(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.TextField_SetStringValue(textField.ptr, cText)
}

// Remove - removes a Text Field from the parent view
func (textField *TextField) Remove() {
	C.TextField_Remove(textField.ptr)
}

func (textField *TextField) Enabled() bool {
	return C.TextField_Enabled(textField.ptr) == 1
}

// SetEnabled sets the enabled value of the text field. CANNOT BE CHANGED AT RUNTIME
func (textField *TextField) SetEnabled(enabled bool) {
	if enabled {
		C.TextField_SetEnabled(textField.ptr, 1)
	} else {
		C.TextField_SetEnabled(textField.ptr, 0)
	}
}

func (textField *TextField) Editable() bool {
	return C.TextField_Editable(textField.ptr) == 1
}

func (textField *TextField) SetEditable(editable bool) {
	if editable {
		C.TextField_SetEditable(textField.ptr, 1)
	} else {
		C.TextField_SetEditable(textField.ptr, 0)
	}
}

func (textField *TextField) SetFontFamily(fontFamily string) {
	cText := C.CString(fontFamily)
	defer C.free(unsafe.Pointer(cText))
	C.TextField_SetFontFamily(textField.ptr, cText)
}

func (textField *TextField) SetFontSize(fontSize int) {
	C.TextField_SetFontSize(textField.ptr, C.int(fontSize))
}

func (textField *TextField) SetColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetColor(textField.ptr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBackgroundColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetBackgroundColor(textField.ptr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.TextField_SetBorderColor(textField.ptr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (textField *TextField) SetBorderWidth(borderWidth int) {
	C.TextField_SetBorderWidth(textField.ptr, C.int(borderWidth))
}

func (textField *TextField) SetBezeled(bezeled bool) {
	if bezeled {
		C.TextField_SetBezeled(textField.ptr, C.int(1))
	} else {
		C.TextField_SetBezeled(textField.ptr, C.int(0))
	}
}

func (textField *TextField) SetDrawsBackground(drawsBackground bool) {
	if drawsBackground {
		C.TextField_SetDrawsBackground(textField.ptr, C.int(1))
	} else {
		C.TextField_SetDrawsBackground(textField.ptr, C.int(0))
	}
}

func (textField *TextField) SetSelectable(selectable bool) {
	if selectable {
		C.TextField_SetSelectable(textField.ptr, C.int(1))
	} else {
		C.TextField_SetSelectable(textField.ptr, C.int(0))
	}
}

func (textField *TextField) SetAlignmentCenter() {
	C.TextField_SetAlignmentCenter(textField.ptr)
}

func (textField *TextField) SetAlignmentLeft() {
	C.TextField_SetAlignmentLeft(textField.ptr)
}

func (textField *TextField) SetAlignmentRight() {
	C.TextField_SetAlignmentRight(textField.ptr)
}

func (textField *TextField) OnChange(fn func(value string)) {
	textField.callback = fn
}

func (c *TextField) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *TextField) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *TextField) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *TextField) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}

//export onTextFieldDidChange
func onTextFieldDidChange(id C.int) {
	textFieldID := int(id)
	if textFieldID < len(textfields) && textfields[textFieldID].callback != nil {
		tf := textfields[textFieldID]
		tf.callback(tf.StringValue())
	}
}
