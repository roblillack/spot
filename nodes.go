package spot

import "log"

type Node struct {
	Content  Control
	Children []Node
}

func (n Node) Mount(ctx *RenderContext) {
	n.mount(ctx, nil)
}

func (n Node) mount(ctx *RenderContext, parent Control) {
	if n.Content != nil {
		n.Content.Mount(ctx, parent)
	}
	for _, child := range n.Children {
		child.mount(ctx, n.Content)
	}
}

func (n *Node) updateChild(ctx *RenderContext, idx int, new Control) {
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
		new.Mount(ctx, n.Content)
		return
	}

	ok := old.Update(new)
	if !ok {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].Content = new
		new.Mount(ctx, n.Content)
	}
}

func (n *Node) Update(ctx *RenderContext, other Node, parent Control) {
	if n.Content != nil && other.Content == nil {
		if unmountable, ok := n.Content.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Content = nil
	} else if n.Content == nil && other.Content != nil {
		n.Content = other.Content
		n.Content.Mount(ctx, parent)
	} else if n.Content != nil && other.Content != nil {
		ok := n.Content.Update(other.Content)
		if !ok {
			if unmountable, ok := n.Content.(Unmountable); ok {
				unmountable.Unmount()
			}
			n.Content = other.Content
			n.Content.Mount(ctx, parent)
		}
	}

	if len(n.Children) != len(other.Children) {
		for idx := range n.Children {
			n.updateChild(ctx, idx, nil)
		}
		n.Children = make([]Node, len(other.Children))
		for idx := range n.Children {
			n.updateChild(ctx, idx, other.Children[idx].Content)
		}
		return
	}

	for idx := range n.Children {
		n.updateChild(ctx, idx, other.Children[idx].Content)
	}
}

func (n *Node) Layout(ctx *RenderContext, parent Control) {
	log.Printf("Layouting %T\n", n.Content)
	if n.Content != nil {
		if layoutable, ok := n.Content.(Layoutable); ok {
			log.Printf("-> Layoutable\n")
			layoutable.Layout(ctx, parent)
		}
	}

	for _, child := range n.Children {
		child.Layout(ctx, n.Content)
	}
}
