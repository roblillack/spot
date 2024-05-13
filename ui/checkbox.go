package ui

import "github.com/roblillack/spot"

var _ spot.Component = &Checkbox{}

func (b *Checkbox) Equals(other spot.Component) bool {
	next, ok := other.(*Checkbox)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Label == b.Label
}
