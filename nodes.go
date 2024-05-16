package spot

type Node struct {
	HostComponent HostComponent
	Children      []Node
}

type ToNode interface {
	ToNode(ctx *RenderContext) Node
}

func (n Node) Mount(parent HostComponent) {
	if n.HostComponent != nil {
		n.HostComponent.Mount(parent)
	}
	for _, child := range n.Children {
		child.Mount(n.HostComponent)
	}
}

func (n Node) updateChild(idx int, new HostComponent) {
	old := n.Children[idx].HostComponent
	if old == nil && new == nil {
		return
	}

	if old != nil && new == nil {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].HostComponent = nil
		return
	}

	if old == nil && new != nil {
		n.Children[idx].HostComponent = new
		new.Mount(n.HostComponent)
		return
	}

	ok := old.Update(new)
	if !ok {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].HostComponent = new
		new.Mount(n.HostComponent)
	}
}

func (n Node) Update(other Node, parent HostComponent) {
	if n.HostComponent != nil && other.HostComponent == nil {
		if unmountable, ok := n.HostComponent.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.HostComponent = nil
	} else if n.HostComponent == nil && other.HostComponent != nil {
		n.HostComponent = other.HostComponent
		n.HostComponent.Mount(parent)
	} else if n.HostComponent != nil && other.HostComponent != nil {
		ok := n.HostComponent.Update(other.HostComponent)
		if !ok {
			if unmountable, ok := n.HostComponent.(Unmountable); ok {
				unmountable.Unmount()
			}
			n.HostComponent = other.HostComponent
			n.HostComponent.Mount(parent)
		}
	}

	if len(n.Children) != len(other.Children) {
		for idx := range n.Children {
			n.updateChild(idx, nil)
		}
		n.Children = make([]Node, len(other.Children))
		for idx := range n.Children {
			n.updateChild(idx, other.Children[idx].HostComponent)
		}
		return
	}

	for idx := range n.Children {
		n.updateChild(idx, other.Children[idx].HostComponent)
	}
}
