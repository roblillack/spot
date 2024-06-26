package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "textview.h"
import "C"
import "unsafe"

// TextView - represents a textView control that can trigger actions.
type TextView struct {
	ptr      C.TextViewPtr
	callback func(string)
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

// GetText gets the text of the text view
func (textview *TextView) GetText() string {
	return C.GoString(C.TextView_Text(textview.ptr))
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

func (c *TextView) OnChange(fn func(value string)) {
	c.callback = fn
}

//export onTextViewDidChange
func onTextViewDidChange(id C.int) {
	textViewId := int(id)
	if textViewId < len(textviews) && textviews[textViewId].callback != nil {
		tf := textviews[textViewId]
		tf.callback(tf.GetText())
	}
}
