package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"unsafe"
)

// Button represents a button control that can trigger actions.
type Button struct {
	buttonPtr C.ButtonPtr
	callback  func()
}

type ButtonType int32
type ButtonBezelStyle int32

const (
	ButtonTypeMomentaryPushIn       ButtonType = 7
	ButtonTypeMomentaryLight        ButtonType = 0
	ButtonTypeMomentaryChange       ButtonType = 5
	ButtonTypePushOnPushOff         ButtonType = 1
	ButtonTypeOnOff                 ButtonType = 6
	ButtonTypeToggle                ButtonType = 2
	ButtonTypeSwitch                ButtonType = 3
	ButtonTypeRadio                 ButtonType = 4
	ButtonTypeAccelerator           ButtonType = 8
	ButtonTypeMultiLevelAccelerator ButtonType = 9
)

const (
	ButtonBezelStyleRounded           ButtonBezelStyle = 1
	ButtonBezelStyleRegularSquare     ButtonBezelStyle = 2
	ButtonBezelStyleShadowlessSquare  ButtonBezelStyle = 6
	ButtonBezelStyleSmallSquare       ButtonBezelStyle = 10
	ButtonBezelStyleRoundRect         ButtonBezelStyle = 12
	ButtonBezelStyleInline            ButtonBezelStyle = 15
	ButtonBezelStyleRecessed          ButtonBezelStyle = 13
	ButtonBezelStyleDisclosure        ButtonBezelStyle = 5
	ButtonBezelStyleRoundedDisclosure ButtonBezelStyle = 14
	ButtonBezelStyleCircular          ButtonBezelStyle = 7
	ButtonBezelStyleHelpButton        ButtonBezelStyle = 9
	ButtonBezelStyleTexturedRounded   ButtonBezelStyle = 11
	ButtonBezelStyleTexturedSquare    ButtonBezelStyle = 8
)

var buttons []*Button

//export onButtonClicked
func onButtonClicked(id C.int) {
	buttonID := int(id)
	if buttonID < len(buttons) && buttons[buttonID].callback != nil {
		buttons[buttonID].callback()
	}
}

// NewButton constructs a new button at position (x, y) and with size (width x height).
// Gotcha! It is currently hard-coded to quit the app when the button is being pressed, until
// callbacks have been implemented.
// This func is not thread safe.
func NewButton(x int, y int, width int, height int) *Button {
	buttonID := len(buttons)
	buttonPtr := C.Button_New(C.int(buttonID), C.int(x), C.int(y), C.int(width), C.int(height))

	btn := &Button{
		buttonPtr: buttonPtr,
	}
	buttons = append(buttons, btn)
	return btn
}

// SetTitle sets the title of the button, which is the text displayed on the button.
func (btn *Button) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(btn.buttonPtr, cTitle)
}

func (btn *Button) SetButtonType(buttonType ButtonType) {
	C.Button_SetButtonType(btn.buttonPtr, C.int(buttonType))
}

func (btn *Button) SetBezelStyle(bezelStyle ButtonBezelStyle) {
	C.Button_SetBezelStyle(btn.buttonPtr, C.int(bezelStyle))
}

// OnClick - function, that will be triggered, if the button is clicked.
func (btn *Button) OnClick(fn func()) {
	btn.callback = fn
}
