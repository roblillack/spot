//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"fmt"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeDropdown = *goFltk.Choice

func (c *Dropdown) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	c.ref = goFltk.NewChoice(c.X, c.Y, c.Width, c.Height)
	for idx, item := range c.Items {
		idx := idx
		item := item
		c.ref.Add(item, func() {
			fmt.Printf("Selected item: %d/%s\n", idx, item)
			if c.OnSelectionDidChange != nil {
				fmt.Printf("Firing for item: %d/%s\n", idx, item)
				c.OnSelectionDidChange(idx)
			}
		})
	}
	c.ref.SetValue(c.SelectedIndex)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
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

	c.OnSelectionDidChange = nextDropdown.OnSelectionDidChange

	if c.SelectedIndex != nextDropdown.SelectedIndex {
		c.SelectedIndex = nextDropdown.SelectedIndex
		c.ref.SetValue(c.SelectedIndex)
		c.ref.Redraw()
	}

	return true
}
