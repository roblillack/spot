//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"fmt"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeComboBox = *goFltk.Choice

func (w *ComboBox) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewChoice(w.X, w.Y, w.Width, w.Height)
	for idx, item := range w.Items {
		idx := idx
		item := item
		w.ref.Add(item, func() {
			fmt.Printf("Selected item: %d/%s\n", idx, item)
			if w.OnSelectionDidChange != nil {
				fmt.Printf("Firing for item: %d/%s\n", idx, item)
				w.OnSelectionDidChange(idx)
			}
		})
	}
	w.ref.SetValue(w.SelectedIndex)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
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

	w.OnSelectionDidChange = nextComboBox.OnSelectionDidChange

	if w.SelectedIndex != nextComboBox.SelectedIndex {
		w.SelectedIndex = nextComboBox.SelectedIndex
		w.ref.SetValue(w.SelectedIndex)
		w.ref.Redraw()
	}

	return true
}
