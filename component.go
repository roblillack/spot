package spot

import "fmt"

type Element interface{}

type Component interface {
	Render(ctx *RenderContext) Element
}

// type ComponentList  []Component
// func (l ComponentList) Render(ctx *RenderContext) Element {
// 	return l
// }

type Unmountable interface {
	Unmount()
}

type HostComponent interface {
	// Component
	Update(next HostComponent) bool
	Equals(other HostComponent) bool
	Mount(parent HostComponent) any
	// Unmount()
}

type makeRenderable func(ctx *RenderContext) Element

func (r makeRenderable) Render(ctx *RenderContext) Element {
	return r(ctx)
}

var _ Component = makeRenderable(nil)

// func Make(render func(ctx *RenderContext) Element) Component {
// 	return Render(makeRenderable(render))
// }

func Build(el Element) {
	Render(el).Mount(nil)
}

func BuildFn(fn func(ctx *RenderContext) Element) {
	Render(Make(fn)).Mount(nil)
}

func Make(fn func(ctx *RenderContext) Element) Element {
	return makeRenderable(fn)
}

func Render(el Element) Node {
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
