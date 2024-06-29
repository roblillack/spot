//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"slices"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeListBox = *goFltk.Browser

func (c *ListBox) getSelection() []int {
	if c.ref == nil {
		return nil
	}

	if !c.Multiselect {
		if c.ref.Value() == 0 {
			return nil
		}

		return []int{c.ref.Value() - 1}
	}

	var selection []int
	for i := 0; i < c.ref.Size(); i++ {
		if c.ref.IsSelected(i + 1) {
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

	for i := 0; i < c.ref.Size(); i++ {
		c.ref.SetSelected(i+1, false)
	}

	for _, i := range selection {
		c.ref.SetSelected(i+1, true)
	}
}

func (c *ListBox) callback() {
	if c.ref == nil {
		return
	}

	oldSelection := c.Selection

	c.Selection = c.getSelection()
	// fmt.Printf("LISTBOX[%p] callback, multi: %v, selection: %v\n", c, c.Multiselect, c.Selection)

	if !c.Multiselect {
		c.setSelection(c.Selection)
	}

	if c.OnSelect != nil && !slices.Equal(oldSelection, c.Selection) {
		c.OnSelect(c.Selection)
	}
}

func (c *ListBox) Update(nextComponent spot.Mountable) bool {
	// fmt.Printf("LISTBOX[%p] Update\n", c)
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

	if c.Multiselect != next.Multiselect {
		c.Multiselect = next.Multiselect
		if c.ref != nil {
			parent := c.ref.Parent()
			c.ref.Destroy()
			c.ref = nil
			c.Mount(nil, nil)
			parent.Add(c.ref)
		}
	}

	return true
}

func (c *ListBox) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	if c.Multiselect {
		ref := goFltk.NewMultiBrowser(x, y, w, h)
		browser := ref.Browser
		c.ref = &browser
	} else {
		ref := goFltk.NewSelectBrowser(x, y, w, h)
		browser := ref.Browser
		c.ref = &browser
	}
	c.ref.SetCallback(c.callback)
	c.setValues(c.Values)
	c.setSelection(c.Selection)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *ListBox) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *ListBox) Layout(ctx *spot.RenderContext, parent spot.Container) {
	if c.ref == nil {
		return
	}

	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
