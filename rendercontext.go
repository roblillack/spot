package spot

import (
	"fmt"
	"strings"
)

type RenderContext struct {
	content Component
	root    Node
	values  map[int]any
	count   int
}

// BuildNode recursively renders a component and its children into a tree
// of UI controls.
func (ctx *RenderContext) BuildNode(component Component) Node {
	if component == nil {
		return Node{}
	}

	if l, ok := component.(Fragment); ok {
		list := []Node{}
		for _, e := range l {
			childNode := ctx.BuildNode(e)
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

	if container, ok := component.(Container); ok {
		return container.BuildNode(ctx)
	}

	if c, ok := component.(Control); ok {
		return Node{Content: c}
	}

	if r, ok := component.(Component); ok {
		return ctx.BuildNode(r.Render(ctx))
	}

	panic(fmt.Sprintf("Unknown component type: %T", component))
}

func (ctx *RenderContext) Make(render func(*RenderContext) Component) Node {
	subContext, _ := UseState(ctx, &RenderContext{
		content: makeRenderable(render),
		values:  make(map[int]any),
	})
	subContext.count = 0
	root := ctx.BuildNode(subContext.content)
	subContext.root = root
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
	if ctx.root.Content == nil {
		// fmt.Printf("[%v] Root is nil, returning.\n", ctx)
		return
	}

	// fmt.Println("STATE VALUES ******")
	// for i := 0; i < ctx.count; i++ {
	// 	fmt.Printf("%02d -> %v\n", i, ctx.values[i])
	// }
	// fmt.Println("*******************")

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
		oldTree := ctx.root
		// fmt.Println("**** RENDER STARTING ****")
		newTree := ctx.BuildNode(ctx.content)
		// fmt.Println("**** RENDER DONE ****")

		// fmt.Printf("[%v] Old tree: %+v\n", ctx, oldTree)
		// printNodes(oldTree, 0)
		// fmt.Printf("[%v] New tree: %+v\n", ctx, newTree)
		// printNodes(newTree, 0)

		oldTree.Update(newTree, nil)
	})
}
