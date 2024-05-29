package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "combobox.h"
import "C"

type ComboBox struct {
	comboBoxPtr C.ComboBoxPtr
	comboBoxID  int
	callback    func()
}

var comboBoxes []*ComboBox

func NewComboBox(x int, y int, width int, height int) *ComboBox {
	comboBoxID := len(comboBoxes)
	comboBoxPtr := C.ComboBox_New(C.int(comboBoxID), C.int(x), C.int(y), C.int(width), C.int(height))

	comboBox := &ComboBox{
		comboBoxPtr: comboBoxPtr,
		comboBoxID:  comboBoxID,
	}
	comboBoxes = append(comboBoxes, comboBox)
	return comboBox
}

func (comboBox *ComboBox) AddItem(item string) {
	C.ComboBox_AddItem(comboBox.comboBoxPtr, C.CString(item))
}

func (comboBox *ComboBox) SetEditable(editable bool) {
	if editable {
		C.ComboBox_SetEditable(comboBox.comboBoxPtr, C.int(1))
	} else {
		C.ComboBox_SetEditable(comboBox.comboBoxPtr, C.int(0))
	}
}

func (comboBox *ComboBox) SelectedIndex() int {
	return int(C.ComboBox_SelectedIndex(comboBox.comboBoxPtr))
}

func (comboBox *ComboBox) SelectedText() string {
	return C.GoString(C.ComboBox_SelectedText(comboBox.comboBoxPtr))
}

func (comboBox *ComboBox) StringValue() string {
	return C.GoString(C.ComboBox_StringValue(comboBox.comboBoxPtr))
}

func (comboBox *ComboBox) SetSelectedIndex(selectedIndex int) {
	C.ComboBox_SetSelectedIndex(comboBox.comboBoxPtr, C.int(selectedIndex))
}

func (comboBox *ComboBox) SetSelectedText(selectedText string) {
	C.ComboBox_SetSelectedText(comboBox.comboBoxPtr, C.CString(selectedText))
}

func (comboBox *ComboBox) SetStringValue(stringValue string) {
	C.ComboBox_SetStringValue(comboBox.comboBoxPtr, C.CString(stringValue))
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
	C.ComboBox_Remove(comboBox.comboBoxPtr)
}
