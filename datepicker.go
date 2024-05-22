package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "datepicker.h"
import "C"

// DatePicker represents a date picker control that can trigger actions.
type DatePicker struct {
	datePickerPtr C.DatePickerPtr
	datePickerID  int
	callback      func()
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
		datePickerPtr: datePickerPtr,
		datePickerID:  datePickerID,
	}
	datePickers = append(datePickers, datePicker)
	datePickerFormats = append(datePickerFormats, "yyyy-MM-dd")
	return datePicker
}

func (datePicker *DatePicker) SetStyle(style DatePickerStyle) {
	C.DatePicker_SetStyle(datePicker.datePickerPtr, C.int(style))
}

func (datePicker *DatePicker) SetMode(mode DatePickerMode) {
	C.DatePicker_SetMode(datePicker.datePickerPtr, C.int(mode))
}

func (datePicker *DatePicker) SetDate(date string) {
	C.DatePicker_SetDate(datePicker.datePickerPtr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetMinimumDate(date string) {
	C.DatePicker_SetMinimumDate(datePicker.datePickerPtr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetMaximumDate(date string) {
	C.DatePicker_SetMaximumDate(datePicker.datePickerPtr, C.CString(date), C.CString(datePickerFormats[datePicker.datePickerID]))
}

func (datePicker *DatePicker) SetDateFormat(dateFormat string) {
	datePickerFormats[datePicker.datePickerID] = dateFormat
}

func (datePicker *DatePicker) Date() string {
	return C.GoString(C.DatePicker_Date(datePicker.datePickerPtr, C.CString(datePickerFormats[datePicker.datePickerID])))
}

// Remove removes a DatePicker from the parent view again.
func (datePicker *DatePicker) Remove() {
	C.DatePicker_Remove(datePicker.datePickerPtr)
}
