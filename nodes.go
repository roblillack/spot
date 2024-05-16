package spot

type Node struct {
	Content  Control
	Children []Node
}

func (n Node) Mount() {
	n.mount(nil)
}

func (n Node) mount(parent Control) {
	if n.Content != nil {
		n.Content.Mount(parent)
	}
	for _, child := range n.Children {
		child.mount(n.Content)
	}
}

func (n Node) updateChild(idx int, new Control) {
	old := n.Children[idx].Content
	if old == nil && new == nil {
		return
	}

	if old != nil && new == nil {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].Content = nil
		return
	}

	if old == nil && new != nil {
		n.Children[idx].Content = new
		new.Mount(n.Content)
		return
	}

	ok := old.Update(new)
	if !ok {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].Content = new
		new.Mount(n.Content)
	}
}

func (n Node) Update(other Node, parent Control) {
	if n.Content != nil && other.Content == nil {
		if unmountable, ok := n.Content.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Content = nil
	} else if n.Content == nil && other.Content != nil {
		n.Content = other.Content
		n.Content.Mount(parent)
	} else if n.Content != nil && other.Content != nil {
		ok := n.Content.Update(other.Content)
		if !ok {
			if unmountable, ok := n.Content.(Unmountable); ok {
				unmountable.Unmount()
			}
			n.Content = other.Content
			n.Content.Mount(parent)
		}
	}

	if len(n.Children) != len(other.Children) {
		for idx := range n.Children {
			n.updateChild(idx, nil)
		}
		n.Children = make([]Node, len(other.Children))
		for idx := range n.Children {
			n.updateChild(idx, other.Children[idx].Content)
		}
		return
	}

	for idx := range n.Children {
		n.updateChild(idx, other.Children[idx].Content)
	}
}
