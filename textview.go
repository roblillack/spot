package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textview.h"
import "C"

// TextView - represents a textView control that can trigger actions.
type TextView struct {
	textView C.pTextView
	callback  func()
}

var textviews []*TextView

// NewTextView - This func is not thread safe.
func NewTextView(x int, y int, width int, height int) *TextView {
	textViewID := len(textviews)
	textView := C.TextView_New(C.int(textViewID), C.int(x), C.int(y), C.int(width), C.int(height))

	tv := &TextView{
		textView: textView,
	}
	textviews = append(textviews, tv)
	return tv
}
