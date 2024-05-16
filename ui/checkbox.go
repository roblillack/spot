package ui

import "github.com/roblillack/spot"

var _ spot.Control = &Checkbox{}

func (c *Checkbox) Render(ctx *spot.RenderContext) spot.Component {
	return c
}

func (b *Checkbox) Equals(other spot.Control) bool {
	next, ok := other.(*Checkbox)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Label == b.Label
}
