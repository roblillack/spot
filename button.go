package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// Button represents a button control that can trigger actions.
type Button struct {
	buttonPtr C.ButtonPtr
	callback  func()
}

type ButtonType int32
type ButtonBezelStyle int32
type ButtonState int32

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

const (
	ButtonStateValueOff   ButtonState = 0
	ButtonStateValueOn    ButtonState = 1
	ButtonStateValueMixed ButtonState = 2
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

func (btn *Button) Title() string {
	return C.GoString(C.Button_Title(btn.buttonPtr))
}

func (btn *Button) SetButtonType(buttonType ButtonType) {
	C.Button_SetButtonType(btn.buttonPtr, C.int(buttonType))
}

func (btn *Button) SetBezelStyle(bezelStyle ButtonBezelStyle) {
	C.Button_SetBezelStyle(btn.buttonPtr, C.int(bezelStyle))
}

func (btn *Button) SetFontFamily(fontFamily string) {
	cText := C.CString(fontFamily)
	defer C.free(unsafe.Pointer(cText))
	C.Button_SetFontFamily(btn.buttonPtr, cText)
}

func (btn *Button) SetFontSize(fontSize int) {
	C.Button_SetFontSize(btn.buttonPtr, C.int(fontSize))
}

func (btn *Button) SetColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.Button_SetColor(btn.buttonPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (btn *Button) SetBackgroundColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.Button_SetBackgroundColor(btn.buttonPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (btn *Button) SetBorderColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.Button_SetBorderColor(btn.buttonPtr, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (btn *Button) SetBorderWidth(borderWidth int) {
	C.Button_SetBorderWidth(btn.buttonPtr, C.int(borderWidth))
}

func (btn *Button) SetState(state ButtonState) {
	C.Button_SetState(btn.buttonPtr, C.int(state))
}

func (btn *Button) State() ButtonState {
	return ButtonState(int(C.Button_State(btn.buttonPtr)))
}

// OnClick - function, that will be triggered, if the button is clicked.
func (btn *Button) OnClick(fn func()) {
	btn.callback = fn
}
