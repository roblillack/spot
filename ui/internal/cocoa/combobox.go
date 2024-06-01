package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "combobox.h"
// #import "view.h"
import "C"

type ComboBox struct {
	ptr        C.ComboBoxPtr
	comboBoxID int
	callback   func()
}

var comboBoxes []*ComboBox

func NewComboBox(x int, y int, width int, height int) *ComboBox {
	comboBoxID := len(comboBoxes)
	comboBoxPtr := C.ComboBox_New(C.int(comboBoxID), C.int(x), C.int(y), C.int(width), C.int(height))

	comboBox := &ComboBox{
		ptr:        comboBoxPtr,
		comboBoxID: comboBoxID,
	}
	comboBoxes = append(comboBoxes, comboBox)
	return comboBox
}

func (comboBox *ComboBox) AddItem(item string) {
	C.ComboBox_AddItem(comboBox.ptr, C.CString(item))
}

func (comboBox *ComboBox) SetEditable(editable bool) {
	if editable {
		C.ComboBox_SetEditable(comboBox.ptr, C.int(1))
	} else {
		C.ComboBox_SetEditable(comboBox.ptr, C.int(0))
	}
}

func (comboBox *ComboBox) SelectedIndex() int {
	return int(C.ComboBox_SelectedIndex(comboBox.ptr))
}

func (comboBox *ComboBox) SelectedText() string {
	return C.GoString(C.ComboBox_SelectedText(comboBox.ptr))
}

func (comboBox *ComboBox) StringValue() string {
	return C.GoString(C.ComboBox_StringValue(comboBox.ptr))
}

func (comboBox *ComboBox) SetSelectedIndex(selectedIndex int) {
	C.ComboBox_SetSelectedIndex(comboBox.ptr, C.int(selectedIndex))
}

func (comboBox *ComboBox) SetSelectedText(selectedText string) {
	C.ComboBox_SetSelectedText(comboBox.ptr, C.CString(selectedText))
}

func (comboBox *ComboBox) SetStringValue(stringValue string) {
	C.ComboBox_SetStringValue(comboBox.ptr, C.CString(stringValue))
}

//export onSelectionDidChange
func onSelectionDidChange(id C.int) {
	comboBoxID := int(id)
	if comboBoxID < len(comboBoxes) && comboBoxes[comboBoxID].callback != nil {
		comboBoxes[comboBoxID].callback()
	}
}

func (comboBox *ComboBox) OnSelectionDidChange(fn func()) {
	comboBox.callback = fn
}

// Remove removes a ComboBox from the parent view again.
func (comboBox *ComboBox) Remove() {
	C.ComboBox_Remove(comboBox.ptr)
}

func (c *ComboBox) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *ComboBox) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *ComboBox) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *ComboBox) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
