package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "window.h"
// #import "interactiveview.h"
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
