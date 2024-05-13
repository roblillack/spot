package spot

type Component interface {
	Update(next Component) bool
	Equals(other Component) bool
	Mount() any
}

func Make(render func(ctx *RenderContext) Component) Component {
	ctx := &RenderContext{
		render: render,
		values: make(map[int]any),
	}
	root := render(ctx)
	ctx.root = root
	// root.Mount()
	// root.Update(root)
	return root
}

type ComponentList []Component

var _ Component = ComponentList{}

func (s ComponentList) Update(next Component) bool {
	nexts, ok := next.(ComponentList)
	if !ok {
		return false
	}

	if len(s) != len(nexts) {
		return false
	}

	for i, c := range s {
		if !c.Update(nexts[i]) {
			return false
		}
	}

	return true
}

func (s ComponentList) Equals(other Component) bool {
	nexts, ok := other.(ComponentList)
	if !ok {
		return false
	}

	if len(s) != len(nexts) {
		return false
	}

	for i, c := range s {
		if !c.Equals(nexts[i]) {
			return false
		}
	}

	return true
}

func (s ComponentList) Mount() any {
	for _, c := range s {
		c.Mount()
	}
	return nil
}
