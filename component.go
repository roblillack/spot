package spot

type Component interface {
	Render(ctx *RenderContext) Component
}

type Fragment []Component

func (l Fragment) Render(ctx *RenderContext) Component {
	return l
}

// Control is a component that can be mounted into the UI tree.
type Control interface {
	Component
	Mount(parent Control) any
	Update(next Control) bool
	// Unmount()
}

// Unmountable is a control component that can be unmounted from the UI tree again.
type Unmountable interface {
	Control
	Unmount()
}

// Container is a control component that owns other controls.
type Container interface {
	Control
	ContentWidth() int
	ContentHeight() int
	BuildNode(ctx *RenderContext) Node // BuildNode renders the control and its children into tree of nodes.
}

type RenderFn = func(ctx *RenderContext) Component

type makeRenderable RenderFn

func (r makeRenderable) Render(ctx *RenderContext) Component {
	return r(ctx)
}

var _ Component = makeRenderable(nil)

// Make creates a component from a render function.
func Make(fn RenderFn) Component {
	return makeRenderable(fn)
}

// Build renders a component into a tree of controls. This tree can be mounted
// to display the UI.
func Build(el Component) Node {
	ctx := &RenderContext{
		content: el,
		values:  make(map[int]any),
	}
	rendered := ctx.BuildNode(el)
	ctx.root = rendered
	// fmt.Println("Rendered control tree:")
	// printNodes(rendered, 0)
	return rendered
}

// BuildFn renders a render function into a tree of controls. It is a shortcut
// for `Build(Make(fn))`.
func BuildFn(fn RenderFn) Node {
	return Build(Make(fn))
}

// Mount renders a component into a tree of controls and mounts it into the UI.
// This is a shortcut for `Build(el).Mount()`.
func Mount(el Component) {
	Build(el).Mount()
}

// MountFn creates a component from a render function, builds it into a tree of
// controls, and mounts it into the UI. This is a shortcut for
// `Build(Make(fn)).Mount()`.
func MountFn(fn func(ctx *RenderContext) Component) {
	Build(Make(fn)).Mount()
}
