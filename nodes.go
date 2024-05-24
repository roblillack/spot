package spot

import "fmt"

type Node struct {
	Content  Control
	Children []Node
}

func (n Node) Mount() {
	n.mount(nil)
}

func (n Node) mount(parent Control) {
	if n.Content != nil {
		fmt.Printf("Mounting %T[%p]\n", n.Content, n.Content)
		n.Content.Mount(parent)
	}
	for _, child := range n.Children {
		fmt.Printf("Mounting child %T[%p] of %T[%p]\n", child.Content, child.Content, n.Content, n.Content)
		child.mount(n.Content)
	}
}

func (n Node) updateChild(idx int, other Node) {
	n.Children[idx].Update(other, n.Content)
	return
	new := other.Content

	old := n.Children[idx].Content
	oldEmpty := isEmpty(old)
	newEmpty := isEmpty(new)

	fmt.Printf("Updating child #%d, %T[%p] <- %T[%p]\n", idx, old, old, new, new)
	if oldEmpty && newEmpty {
		fmt.Println("Both empty, nothing to do")
		return
	}

	if !oldEmpty && newEmpty {
		fmt.Println("Unmounting old child")
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].Content = nil
		return
	}

	if oldEmpty && !newEmpty {
		fmt.Println("Mounting new child")
		n.Children[idx].Content = new
		n.Children[idx].mount(n.Content)
		return
	}

	fmt.Println("Updating child")
	ok := old.Update(new)
	if !ok {
		if unmountable, ok := old.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Children[idx].Content = new
		n.Children[idx].mount(n.Content)
	}
}

func (n Node) Update(other Node, parent Control) {
	if n.Content != nil && other.Content == nil {
		if unmountable, ok := n.Content.(Unmountable); ok {
			unmountable.Unmount()
		}
		n.Content = nil
	} else if n.Content == nil && other.Content != nil {
		fmt.Printf("Will replace %T[%p] with %T[%p]\n", n.Content, n.Content, other.Content, other.Content)
		n.Content = other.Content
		fmt.Printf("Mounting new Node %T[%p] with parent %T[%p])\n", n.Content, n.Content, parent, parent)
		n.mount(parent)
	} else if n.Content != nil && other.Content != nil {
		fmt.Printf("Will update %T[%p] with new data\n", n.Content, n.Content)
		ok := n.Content.Update(other.Content)
		fmt.Printf("Update returned %v\n", ok)
		if !ok {
			if unmountable, ok := n.Content.(Unmountable); ok {
				unmountable.Unmount()
			}
			n.Content = other.Content
			n.mount(parent)
		}
	}

	if len(n.Children) != len(other.Children) {
		for idx := range n.Children {
			n.updateChild(idx, Node{})
		}
		n.Children = make([]Node, len(other.Children))
		for idx := range n.Children {
			n.updateChild(idx, other.Children[idx])
		}
		return
	}

	for idx := range n.Children {
		n.updateChild(idx, other.Children[idx])
	}
}

func (n Node) Unmount() {
	for _, child := range n.Children {
		child.Unmount()
	}

	if n.Content != nil {
		if unmountable, ok := n.Content.(Unmountable); ok {
			unmountable.Unmount()
		}
	}
}
