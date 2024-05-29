//go:build !fltk && (darwin || cocoa)

package ui

import (
	"fmt"
	"strconv"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type spinner struct {
	stepper   *cocoa.Stepper
	textfield *cocoa.TextField
}

type nativeTypeSpinner = *spinner

func (w *Spinner) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Spinner)
	if !ok {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		if w.ref != nil {
			w.ref.stepper.SetValue(w.Value)
			w.ref.textfield.SetStringValue(fmt.Sprintf("%.0f", w.Value))
		}
	}

	if next.Min != w.Min {
		w.Min = next.Min
		if w.ref != nil {
			w.ref.stepper.SetMinValue(w.Min)
		}
	}

	if next.Max != w.Max {
		w.Max = next.Max
		if w.ref != nil {
			w.ref.stepper.SetMaxValue(w.Max)
		}
	}

	if next.Step != w.Step {
		w.Step = next.Step
		if w.ref != nil {
			w.ref.stepper.SetIncrement(w.Step)
		}
	}

	w.OnValueChanged = next.OnValueChanged

	return true
}

func (w *Spinner) textFieldCallback(value string) {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		w.Value = v
		w.ref.stepper.SetValue(w.Value)
		if w.OnValueChanged != nil {
			w.OnValueChanged(v)
		}
	}
}

func (w *Spinner) stepperCallback(value float64) {
	w.Value = value
	w.ref.textfield.SetStringValue(fmt.Sprintf("%.0f", w.Value))
	if w.OnValueChanged != nil {
		w.OnValueChanged(value)
	}
}

func (w *Spinner) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	spinner := &spinner{}
	stepperWidth := 16
	stepperPadding := 0
	spinner.stepper = cocoa.NewStepper(w.X+w.Width-stepperWidth, w.Y, stepperWidth, w.Height)
	spinner.stepper.SetValue(w.Value)
	spinner.stepper.SetMinValue(w.Min)
	spinner.stepper.SetMaxValue(w.Max)
	spinner.stepper.SetIncrement(w.Step)
	spinner.stepper.SetValueWraps(true)
	spinner.stepper.OnStepperValueChanged(w.stepperCallback)

	spinner.textfield = cocoa.NewTextField(w.X, w.Y, w.Width-stepperWidth-stepperPadding, w.Height)
	spinner.textfield.SetStringValue(fmt.Sprintf("%.0f", w.Value))
	spinner.textfield.OnChange(w.textFieldCallback)

	w.ref = spinner

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddStepper(w.ref.stepper)
		window.ref.AddTextField(w.ref.textfield)
	}

	return w.ref
}
