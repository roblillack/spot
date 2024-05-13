package ui

import "journey/spot"

var _ spot.Component = &Spinner{}

func (b *Spinner) Equals(other spot.Component) bool {
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
