package ui

import "github.com/roblillack/spot"

var _ spot.Control = &Spinner{}

func (c *Spinner) Render(ctx *spot.RenderContext) spot.Component {
	return c
}

func (b *Spinner) Equals(other spot.Control) bool {
	next, ok := other.(*Spinner)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Max == b.Max && next.Min == b.Min &&
		next.Value == b.Value &&
		next.Step == b.Step
}
