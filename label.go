package gocoa

// NewLabel - This func is not thread safe.
// Label is only for convenience. Under the hood it's still a textfield with some preconfiguration in place
func NewLabel(x int, y int, width int, height int) *TextField {
	textField := NewTextField(x, y, width, height)
	textField.SetBezeled(false)
	textField.SetDrawsBackground(false)
	textField.SetEditable(false)
	textField.SetSelectable(false)
	return textField
}
