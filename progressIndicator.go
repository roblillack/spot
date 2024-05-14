package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "progressIndicator.h"
import "C"

// ProgressIndicator represents a indicator control that can trigger actions.
type ProgressIndicator struct {
	progressIndicatorPtr C.ProgressIndicatorPtr
	minValue             float64
	maxValue             float64
	value                float64
	autohide             bool
}

// NewProgressIndicator constructs a new ProgressIndicator
func NewProgressIndicator(x int, y int, w int, h int) *ProgressIndicator {
	indicatorPtr := C.ProgressIndicator_New(C.int(x), C.int(y), C.int(w), C.int(h))
	indicator := &ProgressIndicator{
		progressIndicatorPtr: indicatorPtr,
		minValue:             0.00,
		maxValue:             100.00,
		value:                0.00,
		autohide:             false,
	}
	indicator.SetIsIndeterminate(false)
	// indicator.Hide()
	return indicator
}

// StartAnimation - Starts the animation of an indeterminate progress indicator.
func (indicator *ProgressIndicator) StartAnimation() {
	C.ProgressIndicator_StartAnimation(indicator.progressIndicatorPtr)
}

// StopAnimation - Stops the animation of an indeterminate progress indicator.
func (indicator *ProgressIndicator) StopAnimation() {
	C.ProgressIndicator_StopAnimation(indicator.progressIndicatorPtr)
}

// SetAutohide - if set to true, the indicator will disapear after 100%, default false
func (indicator *ProgressIndicator) SetAutohide(value bool) {
	indicator.autohide = value
}

// SetLimits - sets min and max values
func (indicator *ProgressIndicator) SetLimits(minValue float64, maxValue float64) {
	C.ProgressIndicator_SetLimits(indicator.progressIndicatorPtr, C.double(minValue), C.double(maxValue))
}

// GetValue - returns the current value of the indicator
func (indicator *ProgressIndicator) GetValue() float64 {
	return indicator.value
}

// SetValue - sets the value to `value`
func (indicator *ProgressIndicator) SetValue(value float64) {
	indicator.updateValue(value)
	C.ProgressIndicator_SetValue(indicator.progressIndicatorPtr, C.double(value))
	if indicator.autohide && value > indicator.maxValue {
		C.ProgressIndicator_Hide(indicator.progressIndicatorPtr)
	}
}

// IncrementBy - increments by `inc`
func (indicator *ProgressIndicator) IncrementBy(inc float64) {
	value := indicator.value + inc
	indicator.updateValue(value)
	C.ProgressIndicator_IncrementBy(indicator.progressIndicatorPtr, C.double(inc))
	if indicator.autohide && value > indicator.maxValue {
		C.ProgressIndicator_Hide(indicator.progressIndicatorPtr)
	}
}

func (indicator *ProgressIndicator) updateValue(value float64) {
	if value < 0 {
		indicator.value = 0.00
		return
	}
	indicator.value = value
}

// GetIsIndeterminate - return if the progressbar is indeterminate
func (indicator *ProgressIndicator) GetIsIndeterminate() bool {
	value := C.ProgressIndicator_IsIndeterminate(indicator.progressIndicatorPtr)
	return value == 1
}

// SetIsIndeterminate - sets if the progressbar is indeterminate
func (indicator *ProgressIndicator) SetIsIndeterminate(value bool) {
	if value {
		C.ProgressIndicator_SetIsIndeterminate(indicator.progressIndicatorPtr, 1)
	} else {
		C.ProgressIndicator_SetIsIndeterminate(indicator.progressIndicatorPtr, 0)
	}
}

// SetDisplayedWhenStopped - A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
func (indicator *ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	if value {
		C.ProgressIndicator_SetDisplayedWhenStopped(indicator.progressIndicatorPtr, 1)
	} else {
		C.ProgressIndicator_SetDisplayedWhenStopped(indicator.progressIndicatorPtr, 0)
	}
}

// Show - shows the Progressbar
func (indicator *ProgressIndicator) Show() {
	C.ProgressIndicator_Show(indicator.progressIndicatorPtr)
}

// Hide - hides the Progressbar
func (indicator *ProgressIndicator) Hide() {
	C.ProgressIndicator_Hide(indicator.progressIndicatorPtr)
}

// Remove - removes the indicator from the superview
func (indicator *ProgressIndicator) Remove() {
	C.ProgressIndicator_Remove(indicator.progressIndicatorPtr)
}
