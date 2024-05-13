package ui

import "journey/spot"

var _ spot.Component = &Button{}

func (b *Button) Equals(other spot.Component) bool {
	next, ok := other.(*Button)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Title == b.Title
}
