//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeDropdown = *cocoa.ComboBox

func (c *Dropdown) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	c.ref = cocoa.NewComboBox(c.X, c.Y, c.Width, c.Height)
	for _, item := range c.Items {
		c.ref.AddItem(item)
	}
	c.ref.SetEditable(c.Editable)
	c.ref.SetSelectedIndex(c.SelectedIndex)
	c.ref.OnSelectionDidChange(func() {
		if c.OnSelectionDidChange != nil {
			c.OnSelectionDidChange(c.ref.SelectedIndex())
		}
	})

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddComboBox(c.ref)
	}

	return c.ref
}

func (c *Dropdown) Update(next spot.Control) bool {
	nextDropdown, ok := next.(*Dropdown)
	if !ok {
		return false
	}

	if len(c.Items) != len(nextDropdown.Items) {
		c.Items = nextDropdown.Items
		// w.ref.ClearItems()
		// for _, item := range w.Items {
		// 	w.ref.AddItem(item)
		// }
	}

	if c.SelectedIndex != nextDropdown.SelectedIndex {
		c.SelectedIndex = nextDropdown.SelectedIndex
		c.ref.SetSelectedIndex(c.SelectedIndex)
	}

	if c.Editable != nextDropdown.Editable {
		c.Editable = nextDropdown.Editable
		c.ref.SetEditable(c.Editable)
	}

	return true
}
