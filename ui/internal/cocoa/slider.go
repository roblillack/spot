package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "slider.h"
// #import "view.h"
import "C"

type Slider struct {
	ptr      C.SliderPtr
	callback func()
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
		ptr: sliderPtr,
	}
	sliders = append(sliders, slider)
	return slider
}

func (slider *Slider) SetMaximumValue(val float64) {
	C.Slider_SetMaximumValue(slider.ptr, C.double(val))
}

func (slider *Slider) SetMinimumValue(val float64) {
	C.Slider_SetMinimumValue(slider.ptr, C.double(val))
}

func (slider *Slider) SetValue(val float64) {
	C.Slider_SetValue(slider.ptr, C.double(val))
}

func (slider *Slider) Value() float64 {
	return float64(C.Slider_Value(slider.ptr))
}

func (slider *Slider) SetSliderType(sliderType SliderType) {
	C.Slider_SetSliderType(slider.ptr, C.int(sliderType))
}

func (slider *Slider) OnSliderValueChanged(fn func()) {
	slider.callback = fn
}

// Remove removes a Slider from the parent view again.
func (slider *Slider) Remove() {
	C.Slider_Remove(slider.ptr)
}

func (c *Slider) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *Slider) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *Slider) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *Slider) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
