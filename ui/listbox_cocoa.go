//go:build !fltk && (darwin || cocoa)

package ui

import (
	"slices"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeListBox = *cocoa.TableView

func (c *ListBox) getSelection() []int {
	if c.ref == nil {
		return nil
	}

	if !c.Multiselect {
		for i := 0; i < c.ref.NumberOfRows(); i++ {
			if c.ref.IsRowSelected(i) {
				return []int{i}
			}
		}

		return []int{}
	}

	var selection []int
	for i := 0; i < c.ref.NumberOfRows(); i++ {
		if c.ref.IsRowSelected(i) {
			selection = append(selection, i)
		}
	}

	return selection
}

func (c *ListBox) setValues(values []string) {
	c.Values = values

	if c.ref == nil {
		return
	}

	c.ref.Clear()
	for _, v := range values {
		c.ref.Add(v)
	}
}

func (c *ListBox) setSelection(selection []int) {
	c.Selection = selection

	if c.ref == nil {
		return
	}

	c.ref.DeselectAll()

	for _, i := range selection {
		c.ref.SelectRowIndex(i)
		c.ref.ScrollRowToVisible(i)
	}
}

func (c *ListBox) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*ListBox)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	if !slices.Equal(c.Values, next.Values) {
		c.setValues(next.Values)
	}

	if !slices.Equal(c.Selection, next.Selection) {
		c.setSelection(next.Selection)
	}

	c.OnSelect = next.OnSelect
	c.ref.OnSelectionDidChange(c.OnSelect)

	if c.Multiselect != next.Multiselect {
		c.Multiselect = next.Multiselect
		c.ref.SetAllowsMultipleSelection(c.Multiselect)
	}

	return true
}

func (c *ListBox) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	c.ref = cocoa.NewTableView(c.X, c.Y, c.Width, c.Height)
	c.ref.SetAllowsMultipleSelection(c.Multiselect)

	c.setValues(c.Values)
	c.setSelection(c.Selection)
	c.ref.OnSelectionDidChange(c.OnSelect)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTableView(c.ref)
	}

	return c.ref
}
