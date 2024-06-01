package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "datepicker.h"
// #import "view.h"
import "C"

// DatePicker represents a date picker control that can trigger actions.
type DatePicker struct {
	ptr          C.DatePickerPtr
	datePickerID int
	callback     func()
}

type DatePickerStyle int32
type DatePickerMode int32

const (
	DatePickerStyleClockAndCalendar    DatePickerStyle = 1
	DatePickerStyleTextField           DatePickerStyle = 2
	DatePickerStyleTextFieldAndStepper DatePickerStyle = 0
	DatePickerModeRange                DatePickerMode  = 1
	DatePickerModeSingle               DatePickerMode  = 0
)

var datePickers []*DatePicker
var datePickerFormats []string

func NewDatePicker(x int, y int, width int, height int) *DatePicker {
	datePickerID := len(datePickers)
	datePickerPtr := C.DatePicker_New(C.int(datePickerID), C.int(x), C.int(y), C.int(width), C.int(height))

	datePicker := &DatePicker{
		ptr:          datePickerPtr,
		datePickerID: datePickerID,
	}
	datePickers = append(datePickers, datePicker)
	datePickerFormats = append(datePickerFormats, "yyyy-MM-dd")
	return datePicker
}

func (datePicker *DatePicker) SetStyle(style DatePickerStyle) {
	C.DatePicker_SetStyle(datePicker.ptr, C.int(style))
}

func (datePicker *DatePicker) SetMode(mode DatePickerMode) {
	C.DatePicker_SetMode(datePicker.ptr, C.int(mode))
}

func (datePicker *DatePicker) SetDate(date string) {
	C.DatePicker_SetDate(datePicker.ptr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetMinimumDate(date string) {
	C.DatePicker_SetMinimumDate(datePicker.ptr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetMaximumDate(date string) {
	C.DatePicker_SetMaximumDate(datePicker.ptr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetDateFormat(dateFormat string) {
	datePickerFormats[datePicker.datePickerID] = dateFormat
}

func (datePicker *DatePicker) Date() string {
	return C.GoString(C.DatePicker_Date(datePicker.ptr, C.CString(datePickerFormats[datePicker.datePickerID])))
}

// Remove removes a DatePicker from the parent view again.
func (datePicker *DatePicker) Remove() {
	C.DatePicker_Remove(datePicker.ptr)
}

func (c *DatePicker) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *DatePicker) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *DatePicker) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *DatePicker) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
