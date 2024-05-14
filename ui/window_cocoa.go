//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type Window struct {
	Title     string
	Width     int
	Height    int
	X         int
	Y         int
	Resizable bool
	Children  []spot.Component
	ref       *gocoa.Window
}

func (w *Window) Equals(other spot.Component) bool {
	next, ok := other.(*Window)
	if !ok {
		return false
	}

	if next == nil && w != nil || next != nil && w == nil {
		return false
	}

	if len(next.Children) != len(w.Children) {
		return false
	}

	for i, child := range w.Children {
		if !child.Equals(next.Children[i]) {
			return false
		}
	}

	return next.Title == w.Title &&
		next.Width == w.Width &&
		next.Height == w.Height &&
		next.X == w.X &&
		next.Y == w.Y &&
		next.Resizable == w.Resizable
}

func (w *Window) mountChild(child spot.Component) spot.Component {
	if child == nil {
		return nil
	}

	switch c := child.(type) {
	case spot.ComponentList:
		for _, cc := range c {
			w.mountChild(cc)
		}
	case *Button:
		w.ref.AddButton(c.Mount().(*gocoa.Button))
	case *Checkbox:
		w.ref.AddButton(c.Mount().(*gocoa.Button))
	case *ComboBox:
		w.ref.AddComboBox(c.Mount().(*gocoa.ComboBox))
	case *Dial:
		w.ref.AddSlider(c.Mount().(*gocoa.Slider))
	case *Label:
		w.ref.AddTextView(c.Mount().(*gocoa.TextView))
	case *ProgressIndicator:
		w.ref.AddProgressIndicator(c.Mount().(*gocoa.ProgressIndicator))
	case *Slider:
		w.ref.AddSlider(c.Mount().(*gocoa.Slider))
	case *Spinner:
		w.ref.AddTextField(c.Mount().(*gocoa.TextField))
	case *TextField:
		w.ref.AddTextField(c.Mount().(*gocoa.TextField))
	case *TextView:
		w.ref.AddTextView(c.Mount().(*gocoa.TextView))
	default:
		panic("Unknown component type")
	}

	return child
}

func (w *Window) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetTitle(w.Title)
		}
	}

	if next.Width != w.Width || next.Height != w.Height {
		w.Width = next.Width
		w.Height = next.Height
	}

	if next.X != w.X || next.Y != w.Y {
		w.X = next.X
		w.Y = next.Y
	}

	if next.Resizable != w.Resizable {
		w.Resizable = next.Resizable
		if w.ref != nil {
			w.ref.SetAllowsResizing(w.Resizable)
		}
	}

	// TODO: We should introduce a "key" concept to avoid
	// re-building the whole tree
	if len(next.Children) != len(w.Children) {
		w.Children = next.Children
		for _, child := range w.Children {
			w.mountChild(child)
		}
	} else {
		for i, child := range w.Children {
			// TODO: Missing functionality to remove
			// a child from the tree
			if child == nil && next.Children[i] == nil {
				continue
			}
			if child == nil {
				child = next.Children[i]
				w.mountChild(child)
			}
			child.Update(next.Children[i])
		}
	}

	return true
}

func (w *Window) Mount() any {
	w.ref = gocoa.NewCenteredWindow(w.Title, w.Width, w.Height)
	for _, child := range w.Children {
		w.mountChild(child)
	}
	w.ref.SetAllowsResizing(w.Resizable)

	w.ref.MakeKeyAndOrderFront()
	w.ref.AddDefaultQuitMenu()
	return w.ref
}
