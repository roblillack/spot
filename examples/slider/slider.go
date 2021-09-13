package main

import (
	"fmt"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("Slider example", 400, 300)

	label := gocoa.NewLabel(165, 130, 300, 50)
	label.SetStringValue("test")
	wnd.AddLabel(label)

	slider := gocoa.NewSlider(25, 125, 300, 150)
	slider.SetMaximumValue(10.0)
	slider.SetMinimumValue(1.0)
	slider.SetValue(2.5)
	slider.SetSliderType(gocoa.SliderTypeLinear)
	slider.OnSliderValueChanged(func() {
		sliderVal := fmt.Sprintf("%f", slider.Value())
		label.SetStringValue(sliderVal)
	})
	wnd.AddSlider(slider)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
