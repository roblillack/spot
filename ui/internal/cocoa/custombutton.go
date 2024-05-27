package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "window.h"
// #import "custombutton.h"
// #include <stdlib.h>
import "C"
import "image"

// Button represents a button control that can trigger actions.
type CustomButton struct {
	ptr      C.ButtonPtr
	callback func(x, y int, secondary bool)
}

var customButtons []*CustomButton

//export onCustomButtonClicked
func onCustomButtonClicked(id C.int, x C.int, y C.int, secondary C.bool) {
	buttonID := int(id)
	if buttonID < len(buttons) && customButtons[buttonID].callback != nil {
		customButtons[buttonID].callback(int(x), int(y), bool(secondary))
	}
}

func NewCustomButton(x int, y int, width int, height int) *CustomButton {
	buttonID := len(customButtons)
	ptr := C.CustomButton_New(C.int(buttonID), C.int(x), C.int(y), C.int(width), C.int(height))
	btn := &CustomButton{ptr: ptr}
	customButtons = append(customButtons, btn)
	return btn
}

func (c *CustomButton) OnClick(fn func(x, y int, secondary bool)) {
	c.callback = fn
}

func (c *CustomButton) Remove() {
	C.Button_Remove(c.ptr)
}

func (c *CustomButton) SetImage(img *image.RGBA) {
	bytes := C.CBytes(img.Pix)
	nsImage := C.Image_NewWithRGBA(C.int(img.Bounds().Dx()), C.int(img.Bounds().Dy()), (*C.uchar)(bytes))
	C.Button_SetImage(c.ptr, nsImage)
	C.free(bytes)
}

func (w *Window) AddCustomButton(c *CustomButton) {
	C.Window_AddButton(w.winPtr, c.ptr)
}
