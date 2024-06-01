package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import (
	"unsafe"
)

// WindowEvent - different window delegate Events
type WindowEvent int

const (
	didResize        WindowEvent = 0
	didMove          WindowEvent = 1
	didMiniaturize   WindowEvent = 2
	didDeminiaturize WindowEvent = 3
)

// EventHandler - handler functions that accepts the updated window as parameter
type EventHandler func(wnd *Window)

// Window is just that.
type Window struct {
	title     string
	x         int
	y         int
	w         int
	h         int
	callbacks map[WindowEvent]EventHandler
	winPtr    unsafe.Pointer
}

// Screen the screen of the window.
type Screen struct {
	X int
	Y int
}

var windows []*Window

// NewWindow constructs and returns a new window.
func NewWindow(title string, x int, y int, w int, h int) *Window {
	windowID := len(windows)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	wnd := &Window{
		title:     title,
		x:         x,
		y:         y,
		w:         w,
		h:         h,
		callbacks: make(map[WindowEvent]EventHandler),
		winPtr:    C.Window_New(C.int(windowID), C.int(x), C.int(y), C.int(w), C.int(h), cTitle)}
	windows = append(windows, wnd)
	return wnd
}

// NewCenteredWindow constructs and returns a new window.
func NewCenteredWindow(title string, w int, h int) *Window {
	windowID := len(windows)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	wnd := &Window{
		title:     title,
		w:         w,
		h:         h,
		callbacks: make(map[WindowEvent]EventHandler),
		winPtr:    C.Centered_Window_New(C.int(windowID), C.int(w), C.int(h), cTitle)}
	wnd.x = int(C.Screen_Center_X(wnd.winPtr))
	wnd.y = int(C.Screen_Center_Y(wnd.winPtr))
	windows = append(windows, wnd)
	return wnd
}

// GetScreen - returns the screen dimensions
func (wnd *Window) GetScreen() *Screen {
	return &Screen{
		X: int(C.Screen_X(wnd.winPtr)),
		Y: int(C.Screen_Y(wnd.winPtr))}
}

// MakeKeyAndOrderFront moves the window to the front of the screen list, within its
// level and it shows the window.
func (wnd *Window) MakeKeyAndOrderFront() {
	C.Window_MakeKeyAndOrderFront(wnd.winPtr)
}

// AddButton adds a Button to the window.
func (wnd *Window) AddButton(btn *Button) {
	C.Window_AddButton(wnd.winPtr, btn.ptr)
}

// AddDatePicker adds a DatePicker to the window.
func (wnd *Window) AddDatePicker(datePicker *DatePicker) {
	C.Window_AddDatePicker(wnd.winPtr, datePicker.ptr)
}

// AddTextView - adds a Text View to the window.
func (wnd *Window) AddTextView(tv *TextView) {
	C.Window_AddTextView(wnd.winPtr, tv.ptr)
}

// AddTextField - adds a Text Field to the window.
func (wnd *Window) AddTextField(tf *TextField) {
	C.Window_AddTextField(wnd.winPtr, tf.ptr)
}

// AddTextField - adds a Button to the window.
func (wnd *Window) AddLabel(tv *TextField) {
	C.Window_AddTextField(wnd.winPtr, tv.ptr)
}

// AddProgressIndicator adds a ProgressIndicator to the window.
func (wnd *Window) AddProgressIndicator(indicator *ProgressIndicator) {
	C.Window_AddProgressIndicator(wnd.winPtr, indicator.ptr)
}

// AddImageView adds an ImageView to the window.
func (wnd *Window) AddImageView(imageView *ImageView) {
	C.Window_AddImageView(wnd.winPtr, imageView.ptr)
}

func (wnd *Window) AddSlider(slider *Slider) {
	C.Window_AddSlider(wnd.winPtr, slider.ptr)
}

func (wnd *Window) AddStepper(stepper *Stepper) {
	C.Window_AddStepper(wnd.winPtr, stepper.ptr)
}

func (wnd *Window) AddComboBox(comboBox *ComboBox) {
	C.Window_AddComboBox(wnd.winPtr, comboBox.ptr)
}

func (wnd *Window) AddTableView(tableView *TableView) {
	C.Window_AddTableView(wnd.winPtr, tableView.ptr)
}

// Update - forces the whole window to repaint
func (wnd *Window) Update() {
	C.Window_Update(wnd.winPtr)
}

func (wnd *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Window_SetTitle(wnd.winPtr, cTitle)
}

func (wnd *Window) SetMiniaturizeButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetMiniaturizeButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetMiniaturizeButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetZoomButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetZoomButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetZoomButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetCloseButtonEnabled(enabled bool) {
	if enabled {
		C.Window_SetCloseButtonEnabled(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetCloseButtonEnabled(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) SetAllowsResizing(allowsResizing bool) {
	if allowsResizing {
		C.Window_SetAllowsResizing(wnd.winPtr, C.int(1))
	} else {
		C.Window_SetAllowsResizing(wnd.winPtr, C.int(0))
	}
}

func (wnd *Window) AddDefaultQuitMenu() {
	C.Window_AddDefaultQuitMenu(wnd.winPtr)
}

func (wnd *Window) OnDidResize(fn EventHandler) {
	wnd.callbacks[didResize] = fn
}

func (wnd *Window) OnDidMiniaturize(fn EventHandler) {
	wnd.callbacks[didMiniaturize] = fn
}

func (wnd *Window) OnDidDeminiaturize(fn EventHandler) {
	wnd.callbacks[didDeminiaturize] = fn
}

func (wnd *Window) OnDidMove(fn EventHandler) {
	wnd.callbacks[didMove] = fn
}

func (wnd *Window) Size() (int, int) {
	return wnd.w, wnd.h
}

func (wnd *Window) Position() (int, int) {
	return wnd.x, wnd.y
}

//export onWindowEvent
func onWindowEvent(id C.int, eventID C.int, x C.int, y C.int, w C.int, h C.int) {
	windowID := int(id)
	event := WindowEvent(eventID)
	if windowID < len(windows) && windows[windowID].callbacks[event] != nil {
		wnd := windows[windowID]
		wnd.x = int(x)
		wnd.y = int(y)
		wnd.w = int(w)
		wnd.h = int(h)
		windows[windowID].callbacks[event](wnd)
	}
}
