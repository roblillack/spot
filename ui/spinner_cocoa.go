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

const stepperWidth = 16
const stepperPadding = 0

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

func (c *Spinner) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	spinner := &spinner{}
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	spinner.stepper = cocoa.NewStepper(x+w-stepperWidth, y, stepperWidth, h)
	spinner.stepper.SetValue(c.Value)
	spinner.stepper.SetMinValue(c.Min)
	spinner.stepper.SetMaxValue(c.Max)
	spinner.stepper.SetIncrement(c.Step)
	spinner.stepper.SetValueWraps(true)
	spinner.stepper.OnStepperValueChanged(c.stepperCallback)

	spinner.textfield = cocoa.NewTextField(x, y, w-stepperWidth-stepperPadding, h)
	spinner.textfield.SetStringValue(fmt.Sprintf("%.0f", c.Value))
	spinner.textfield.OnChange(c.textFieldCallback)

	c.ref = spinner

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddStepper(c.ref.stepper)
		window.ref.AddTextField(c.ref.textfield)
	}

	return c.ref
}

func (c *Spinner) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.stepper.Remove()
	c.ref.textfield.Remove()
	c.ref = nil
}

func (c *Spinner) Layout(ctx *spot.RenderContext, parent spot.Control) {
	if c.ref == nil {
		return
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.stepper.SetFrame(x+w-stepperWidth, y, stepperWidth, h)
	c.ref.textfield.SetFrame(x, y, w-stepperWidth-stepperPadding, h)
}
