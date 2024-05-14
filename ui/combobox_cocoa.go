//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type ComboBox struct {
	X                    int
	Y                    int
	Width                int
	Height               int
	Items                []string
	SelectedIndex        int
	Editable             bool
	OnSelectionDidChange func(index int)
	ref                  *gocoa.ComboBox
}

var _ spot.Component = &ComboBox{}

func (w *ComboBox) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewComboBox(w.X, w.Y, w.Width, w.Height)
	for _, item := range w.Items {
		w.ref.AddItem(item)
	}
	w.ref.SetEditable(w.Editable)
	w.ref.SetSelectedIndex(w.SelectedIndex)
	w.ref.OnSelectionDidChange(func() {
		if w.OnSelectionDidChange != nil {
			w.OnSelectionDidChange(w.ref.SelectedIndex())
		}
	})
	return w.ref
}

func (w *ComboBox) Equals(other spot.Component) bool {
	next, ok := other.(*ComboBox)
	if !ok {
		return false
	}

	if w == nil && next != nil || w != nil && next == nil {
		return false
	}

	if len(w.Items) != len(next.Items) {
		return false
	}

	for i, item := range w.Items {
		if item != next.Items[i] {
			return false
		}
	}

	return next.SelectedIndex == w.SelectedIndex && next.Editable == w.Editable
}

func (w *ComboBox) Update(next spot.Component) bool {
	nextComboBox, ok := next.(*ComboBox)
	if !ok {
		return false
	}

	if len(w.Items) != len(nextComboBox.Items) {
		w.Items = nextComboBox.Items
		// w.ref.ClearItems()
		// for _, item := range w.Items {
		// 	w.ref.AddItem(item)
		// }
	}

	if w.SelectedIndex != nextComboBox.SelectedIndex {
		w.SelectedIndex = nextComboBox.SelectedIndex
		w.ref.SetSelectedIndex(w.SelectedIndex)
	}

	if w.Editable != nextComboBox.Editable {
		w.Editable = nextComboBox.Editable
		w.ref.SetEditable(w.Editable)
	}

	return true
}
