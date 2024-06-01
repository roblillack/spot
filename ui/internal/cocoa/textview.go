package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textview.h"
// #import "view.h"
import "C"
import "unsafe"

// TextView - represents a textView control that can trigger actions.
type TextView struct {
	ptr      C.TextViewPtr
	callback func()
}

var textviews []*TextView

// NewTextView - This func is not thread safe.
func NewTextView(x int, y int, width int, height int) *TextView {
	textViewID := len(textviews)
	textViewPtr := C.TextView_New(C.int(textViewID), C.int(x), C.int(y), C.int(width), C.int(height))

	tv := &TextView{
		ptr: textViewPtr,
	}
	textviews = append(textviews, tv)
	return tv
}

// SetText sets the text of the text view
func (textview *TextView) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.TextView_SetText(textview.ptr, cText)
}

// Remove - removes a Text View from the parent view
func (textview *TextView) Remove() {
	C.TextView_Remove(textview.ptr)
}

// SetText sets the text of the text view
func (textview *TextView) SetFontSize(size int) {
	C.TextView_SetFontSize(textview.ptr, C.int(size))
}

func (c *TextView) SetEditable(editable bool) {
	if editable {
		C.TextView_SetEditable(c.ptr, 1)
	} else {
		C.TextView_SetEditable(c.ptr, 0)
	}
}

func (c *TextView) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *TextView) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *TextView) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *TextView) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
