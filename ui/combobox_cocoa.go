//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeComboBox = *gocoa.ComboBox

func (w *ComboBox) Mount(parent spot.Control) any {
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

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddComboBox(w.ref)
	}

	return w.ref
}

func (w *ComboBox) Update(next spot.Control) bool {
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
