package ui

import "github.com/roblillack/spot"

type Window struct {
	Title     string
	Width     int
	Height    int
	Resizable bool
	Children  []spot.Component
	ref       nativeTypeWindow
}

var _ spot.Component = &Window{}
var _ spot.Control = &Window{}
var _ spot.Container = &Window{}

func (c *Window) Render(ctx *spot.RenderContext) spot.Component {
	return c
}

func (w *Window) BuildNode(ctx *spot.RenderContext) spot.Node {
	kids := []spot.Node{}
	for _, child := range w.Children {
		kid := ctx.BuildNode(child)
		if kid.Content == nil {
			if len(kid.Children) > 0 {
				kids = append(kids, kid.Children...)
			}
			continue
		}
		kids = append(kids, kid)
	}

	return spot.Node{
		Content:  w,
		Children: kids,
	}
}
