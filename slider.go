package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "slider.h"
import "C"

type Slider struct {
	sliderPtr C.SliderPtr
	callback  func()
}

var sliders []*Slider

type SliderType int32

const (
	SliderTypeCircular SliderType = 1
	SliderTypeLinear   SliderType = 0
)

//export onSliderValueChanged
func onSliderValueChanged(id C.int) {
	sliderID := int(id)
	if sliderID < len(sliders) && sliders[sliderID].callback != nil {
		sliders[sliderID].callback()
	}
}

func NewSlider(x int, y int, width int, height int) *Slider {
	sliderID := len(sliders)
	sliderPtr := C.Slider_New(C.int(sliderID), C.int(x), C.int(y), C.int(width), C.int(height))

	slider := &Slider{
		sliderPtr: sliderPtr,
	}
	sliders = append(sliders, slider)
	return slider
}

func (slider *Slider) SetMaximumValue(val float64) {
	C.Slider_SetMaximumValue(slider.sliderPtr, C.double(val))
}

func (slider *Slider) SetMinimumValue(val float64) {
	C.Slider_SetMinimumValue(slider.sliderPtr, C.double(val))
}

func (slider *Slider) SetValue(val float64) {
	C.Slider_SetValue(slider.sliderPtr, C.double(val))
}

func (slider *Slider) Value() float64 {
	return float64(C.Slider_Value(slider.sliderPtr))
}

func (slider *Slider) SetSliderType(sliderType SliderType) {
	C.Slider_SetSliderType(slider.sliderPtr, C.int(sliderType))
}

func (slider *Slider) OnSliderValueChanged(fn func()) {
	slider.callback = fn
}
