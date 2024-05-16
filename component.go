package spot

import "fmt"

type Component interface {
	Render(ctx *RenderContext) Component
}

type Fragment []Component

func (l Fragment) Render(ctx *RenderContext) Component {
	return l
}

type Unmountable interface {
	Unmount()
}

// Control is a component that can be mounted into the UI tree.
type Control interface {
	Component
	Mount(parent Control) any
	Update(next Control) bool
	// Unmount()
}

type makeRenderable func(ctx *RenderContext) Component

func (r makeRenderable) Render(ctx *RenderContext) Component {
	return r(ctx)
}

var _ Component = makeRenderable(nil)

// func Make(render func(ctx *RenderContext) Element) Component {
// 	return Render(makeRenderable(render))
// }

func Build(el Component) {
	Render(el).Mount(nil)
}

func BuildFn(fn func(ctx *RenderContext) Component) {
	Render(Make(fn)).Mount(nil)
}

func Make(fn func(ctx *RenderContext) Component) Component {
	return makeRenderable(fn)
}

func Render(el Component) Node {
	ctx := &RenderContext{
		root:   el,
		values: make(map[int]any),
	}
	rendered := ctx.RenderElement(el)
	ctx.rendered = rendered
	fmt.Println("Rendered component tree:")
	printNodes(rendered, 0)
	return rendered
}
