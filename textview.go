package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textview.h"
import "C"
import "unsafe"

// TextView - represents a textView control that can trigger actions.
type TextView struct {
	textViewPtr C.TextViewPtr
	callback    func()
}

var textviews []*TextView

// NewTextView - This func is not thread safe.
func NewTextView(x int, y int, width int, height int) *TextView {
	textViewID := len(textviews)
	textViewPtr := C.TextView_New(C.int(textViewID), C.int(x), C.int(y), C.int(width), C.int(height))

	tv := &TextView{
		textViewPtr: textViewPtr,
	}
	textviews = append(textviews, tv)
	return tv
}

// SetText sets the text of the text view
func (textview *TextView) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.TextView_SetText(textview.textViewPtr, cText)
}

// SetText sets the text of the text view
func (textview *TextView) SetFontSize(size int) {
	C.TextView_SetFontSize(textview.textViewPtr, C.int(size))
}
