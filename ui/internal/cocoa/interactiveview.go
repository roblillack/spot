package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "window.h"
// #import "interactiveview.h"
// #import "view.h"
// #include <stdlib.h>
import "C"
import (
	"image"
)

type InteractiveView struct {
	ptr      C.InteractiveViewPtr
	callback func(x, y int, secondary bool)
}

var interactiveViews map[C.InteractiveViewPtr]*InteractiveView = make(map[C.InteractiveViewPtr]*InteractiveView)

//export onInteractiveViewClicked
func onInteractiveViewClicked(ptr C.InteractiveViewPtr, x C.int, y C.int, secondary C.bool) {
	if b, ok := interactiveViews[ptr]; ok {
		b.callback(int(x), int(y), bool(secondary))
	}
}

func NewInteractiveView(x int, y int, width int, height int) *InteractiveView {
	ptr := C.InteractiveView_New(C.int(x), C.int(y), C.int(width), C.int(height))
	c := &InteractiveView{ptr: ptr}
	interactiveViews[ptr] = c
	return c
}

func (c *InteractiveView) OnClick(fn func(x, y int, secondary bool)) {
	c.callback = fn
}

func (c *InteractiveView) Remove() {
	C.InteractiveView_Remove(c.ptr)
	delete(interactiveViews, c.ptr)
}

func (c *InteractiveView) SetImage(img *image.RGBA) {
	bytes := C.CBytes(img.Pix)
	nsImage := C.Image_NewWithRGBA(C.int(img.Bounds().Dx()), C.int(img.Bounds().Dy()), (*C.uchar)(bytes))
	C.InteractiveView_SetImage(c.ptr, nsImage)
	C.free(bytes)
}

func (w *Window) AddInteractiveView(c *InteractiveView) {
	C.Window_AddInteractiveView(w.winPtr, c.ptr)
}

func (c *InteractiveView) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *InteractiveView) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *InteractiveView) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *InteractiveView) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
