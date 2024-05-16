//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type Window struct {
	Title     string
	Width     int
	Height    int
	Resizable bool
	Children  []spot.Component
	ref       *goFltk.Window
}

var _ spot.HostComponent = &Window{}

// var _ spot.ComponentContainer = &Window{}

// func (w *Window) GetChildComponents() []spot.HostComponent {
// 	return w.children
// }

var _ spot.ToNode = &Window{}

func (w *Window) ToNode(ctx *spot.RenderContext) spot.Node {
	kids := []spot.Node{}
	for _, child := range w.Children {
		kid := ctx.RenderElement(child)
		if kid.HostComponent == nil {
			if len(kid.Children) > 0 {
				kids = append(kids, kid.Children...)
			}
			continue
		}
		kids = append(kids, kid)
	}

	return spot.Node{
		HostComponent: w,
		Children:      kids,
	}
}

func (w *Window) Equals(other spot.HostComponent) bool {
	return false
	// next, ok := other.(*Window)
	// if !ok {
	// 	return false
	// }

	// if next == nil && w != nil || next != nil && w == nil {
	// 	return false
	// }

	// if len(next.children) != len(w.children) {
	// 	return false
	// }

	// for i, child := range w.children {
	// 	if !child.Equals(next.children[i]) {
	// 		return false
	// 	}
	// }

	// return next.Title == w.Title &&
	// 	next.Width == w.Width &&
	// 	next.Height == w.Height &&
	// 	next.Resizable == w.Resizable
}

// func (w *Window) mountChild(child spot.HostComponent) spot.HostComponent {
// 	if child == nil {
// 		return nil
// 	}

// 	if list, ok := child.(spot.HostComponentList); ok {
// 		for _, cc := range list {
// 			w.mountChild(cc)
// 		}
// 		return child
// 	}

// 	ref := child.Mount()
// 	if widget, ok := ref.(goFltk.Widget); ok {
// 		w.ref.Add(widget)
// 		return child
// 	}

// 	panic("Unknown component type")
// }

func (w *Window) Update(nextComponent spot.HostComponent) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetLabel(w.Title)
		}
	}

	if next.Width != w.Width || next.Height != w.Height {
		w.Width = next.Width
		w.Height = next.Height
	}

	if next.Resizable != w.Resizable {
		w.Resizable = next.Resizable
		if w.ref != nil {
			// w.ref.SetSizeRange(w.Resizable)
		}
	}

	// TODO: We should introduce a "key" concept to avoid
	// re-building the whole tree
	// if len(next.children) != len(w.children) {
	// 	w.children = next.children
	// 	for _, child := range w.children {
	// 		w.mountChild(child)
	// 	}
	// } else {
	// 	for i, child := range w.children {
	// 		// TODO: Missing functionality to remove
	// 		// a child from the tree
	// 		if child == nil && next.children[i] == nil {
	// 			continue
	// 		}
	// 		if child == nil {
	// 			child = next.children[i]
	// 			w.mountChild(child)
	// 		}
	// 		child.Update(next.children[i])
	// 	}
	// }

	return true
}

func (w *Window) Mount(parent spot.HostComponent) any {
	w.ref = goFltk.NewWindow(w.Width, w.Height, w.Title)
	// for _, child := range w.children {
	// 	w.mountChild(child)
	// }
	// w.ref.SetAllowsResizing(w.Resizable)

	// w.ref.MakeKeyAndOrderFront()
	// w.ref.AddDefaultQuitMenu()
	w.ref.End()
	w.ref.Show()
	return w.ref
}
