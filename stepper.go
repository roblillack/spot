package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "stepper.h"
import "C"

type Stepper struct {
	ptr C.StepperPtr
	id  int
	cb  func(value float64)
}

var steppers []*Stepper

//export onStepperValueChanged
func onStepperValueChanged(id C.int) {
	stepper := getStepper(id)
	if stepper != nil && stepper.cb != nil {
		stepper.cb(stepper.Value())
	}
}

func getStepper(id C.int) *Stepper {
	stepperID := int(id)
	if stepperID < len(steppers) {
		return steppers[stepperID]
	}

	return nil
}

func NewStepper(x int, y int, width int, height int) *Stepper {
	stepperID := len(steppers)
	stepperPtr := C.Stepper_New(C.int(stepperID), C.int(x), C.int(y), C.int(width), C.int(height))

	stepper := &Stepper{
		ptr: stepperPtr,
		id:  stepperID,
	}
	steppers = append(steppers, stepper)
	return stepper
}

func (stepper *Stepper) SetMaxValue(val float64) {
	C.Stepper_SetMaxValue(stepper.ptr, C.double(val))
}

func (stepper *Stepper) SetMinValue(val float64) {
	C.Stepper_SetMinValue(stepper.ptr, C.double(val))
}

func (stepper *Stepper) SetIncrement(val float64) {
	C.Stepper_SetIncrement(stepper.ptr, C.double(val))
}

func (stepper *Stepper) SetValue(val float64) {
	C.Stepper_SetValue(stepper.ptr, C.double(val))
}

func (stepper *Stepper) SetValueWraps(val bool) {
	C.Stepper_SetValueWraps(stepper.ptr, C.bool(val))
}

func (stepper *Stepper) Value() float64 {
	return float64(C.Stepper_Value(stepper.ptr))
}

func (stepper *Stepper) OnStepperValueChanged(fn func(value float64)) {
	stepper.cb = fn
}

// Remove removes a Slider from the parent view again.
func (stepper *Stepper) Remove() {
	C.Stepper_Remove(stepper.ptr)
	steppers[stepper.id] = nil
}
