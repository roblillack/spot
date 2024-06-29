package spot

type Node struct {
	Content  Component
	Children []Node
}

func (n Node) Mount(ctx *RenderContext) {
	n.mount(ctx, nil)
}

func (n Node) mount(ctx *RenderContext, parent Mountable) {
	closestMountable := parent
	if mountable, ok := n.Content.(Mountable); ok && mountable != nil {
		mountable.Mount(ctx, parent)
		closestMountable = mountable
	}
	for _, child := range n.Children {
		child.mount(ctx, closestMountable)
	}
}

// func (n *Node) updateChild(ctx *RenderContext, idx int, new Component, parent Mountable) {
// 	old := n.Children[idx].Content
// 	if old == nil && new == nil {
// 		return
// 	}

// 	if old != nil && new == nil {
// 		if oldMountable, ok := old.(Mountable); ok {
// 			oldMountable.Unmount()
// 		}
// 		n.Children[idx].Content = nil
// 		return
// 	}

// 	if old == nil && new != nil {
// 		n.Children[idx].Content = new
// 		if newMountable, ok := new.(Mountable); ok {
// 			newMountable.Mount(ctx, n.Content)
// 		}
// 		return
// 	}

// 	ok := old.Update(new)
// 	if !ok {
// 		if unmountable, ok := old.(Unmountable); ok {
// 			unmountable.Unmount()
// 		}
// 		n.Children[idx].Content = new
// 		new.Mount(ctx, n.Content)
// 	}
// }

func (n Node) Unmount() {
	for _, child := range n.Children {
		child.Unmount()
	}

	if m, ok := n.Content.(Mountable); ok {
		m.Unmount()
	}
}

func (n *Node) Update(ctx *RenderContext, other Node, parent Mountable) {
	if n.Content != nil && other.Content == nil {
		n.Unmount()
		n.Content = nil
		return
	} else if n.Content == nil && other.Content != nil {
		n.Content = other.Content
		n.mount(ctx, parent)
		return
	} else if n.Content != nil && other.Content != nil {
		updated := false
		this, thisOk := n.Content.(Mountable)
		other, otherOk := other.Content.(Mountable)
		if thisOk && otherOk {
			updated = this.Update(other)
			n.Content = other
		}
		if !updated {
			n.Unmount()
			n.Content = other
			n.mount(ctx, parent)
			return
		}
	}

	closestMountable := parent
	if m, ok := n.Content.(Mountable); ok {
		closestMountable = m
	}

	// ---

	if len(n.Children) != len(other.Children) {
		for _, child := range n.Children {
			child.Unmount()
		}
		n.Children = other.Children
		for _, child := range n.Children {
			child.mount(ctx, closestMountable)
		}
		return
	}

	for idx, child := range n.Children {
		child.Update(ctx, other.Children[idx], closestMountable)
	}
}
