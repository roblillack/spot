package spot

import (
	"fmt"
	"strings"
)

type RenderContext struct {
	root     Component
	rendered Node
	// render   func(ctx *RenderContext) Component
	// changed  bool
	values map[int]any
	count  int
}

// RenderElement recursively renders an element and its children into a tree
// of HostComponents.
func (ctx *RenderContext) RenderElement(el Component) Node {
	if el == nil {
		return Node{}
	}

	if l, ok := el.(Fragment); ok {
		list := []Node{}
		for _, e := range l {
			childNode := ctx.RenderElement(e)
			if childNode.Content == nil {
				if len(childNode.Children) == 0 {
					continue
				} else {
					list = append(list, childNode.Children...)
					continue
				}
			}
			list = append(list, childNode)
		}
		return Node{Children: list}
	}

	if container, ok := el.(ToNode); ok {
		return container.ToNode(ctx)
	}

	if c, ok := el.(Control); ok {
		return Node{Content: c}
	}

	if r, ok := el.(Component); ok {
		return ctx.RenderElement(r.Render(ctx))
	}

	panic(fmt.Sprintf("Unknown element type: %T", el))
}

func (ctx *RenderContext) Make(render func(*RenderContext) Component) Node {
	subContext, _ := UseState(ctx, &RenderContext{
		root:   makeRenderable(render),
		values: make(map[int]any),
	})
	subContext.count = 0
	root := ctx.RenderElement(subContext.root)
	subContext.rendered = root
	return root
}

func printNodes(node Node, indent int) {
	if len(node.Children) == 0 {
		fmt.Printf("%s<%T/>\n", strings.Repeat("  ", indent), node.Content)
		return
	}

	fmt.Printf("%s<%T>\n", strings.Repeat("  ", indent), node.Content)
	for _, child := range node.Children {
		printNodes(child, indent+1)
	}
	fmt.Printf("%s</%T>\n", strings.Repeat("  ", indent), node.Content)
}

func (ctx *RenderContext) TriggerUpdate() {
	if ctx.rendered.Content == nil {
		// fmt.Printf("[%v] Root is nil, returning.\n", ctx)
		return
	}

	fmt.Println("STATE VALUES ******")
	for i := 0; i < ctx.count; i++ {
		fmt.Printf("%02d -> %v\n", i, ctx.values[i])
	}
	fmt.Println("*******************")

	// We need to make sure we're running on the main loop
	// for two reasons:
	//
	// 1. We're going to be updating the UI, which is only
	//    allowed on the main thread.
	// 2. We want to ensure not to trigger multiple renders
	//    at the same time, which could lead to weird jumps
	//    in the state of the UI.
	//
	// For reason 2, we could also use a mutex, but this is
	// a simpler solution, as we need to be on the main loop
	// anyway.
	RunOnMainLoop(func() {
		// fmt.Printf("[%v] RENDER TRIGGERED!\n", ctx)
		ctx.count = 0
		oldTree := ctx.rendered
		fmt.Println("**** RENDER STARTING ****")
		// newTree := ctx.RenderElement(ctx.root)
		newTree := ctx.RenderElement(ctx.root)
		fmt.Println("**** RENDER DONE ****")

		fmt.Printf("[%v] Old tree: %+v\n", ctx, oldTree)
		printNodes(oldTree, 0)
		fmt.Printf("[%v] New tree: %+v\n", ctx, newTree)
		printNodes(newTree, 0)

		// if !oldTree.Equals(newTree) {
		// fmt.Printf("[%v] Updating tree!\n", ctx)
		// fmt.Printf("[%v] On main thread here.\n", ctx)
		oldTree.Update(newTree, nil)
		// fmt.Printf("[%v] Updating tree done.\n", ctx)
		// }

		// ctx.changed = false
	})
}
