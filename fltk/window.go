package fltk

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
)

type Window struct {
	Title     string
	Width     int
	Height    int
	X         int
	Y         int
	Resizable bool
	Children  []spot.Component
	ref       *goFltk.Window
}

var _ spot.Component = &Window{}

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

	if list, ok := child.(spot.ComponentList); ok {
		for _, cc := range list {
			w.mountChild(cc)
		}
		return child
	}

	ref := child.Mount()
	if widget, ok := ref.(goFltk.Widget); ok {
		w.ref.Add(widget)
		return child
	}

	panic("Unknown component type")
}

func (w *Window) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetLabel(w.Title)
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
			// w.ref.SetSizeRange(w.Resizable)
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
	w.ref = goFltk.NewWindow(w.Width, w.Height, w.Title)
	for _, child := range w.Children {
		w.mountChild(child)
	}
	// w.ref.SetAllowsResizing(w.Resizable)

	// w.ref.MakeKeyAndOrderFront()
	// w.ref.AddDefaultQuitMenu()
	w.ref.End()
	w.ref.Show()
	return w.ref
}
