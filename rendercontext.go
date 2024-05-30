package spot

import (
	"fmt"
	"strings"
	"sync"
)

type RenderContext struct {
	content Component
	root    Node
	values  map[int]any
	count   int
	mutex   sync.Mutex
}

// BuildNode recursively renders a component and its children into a tree
// of UI controls.
func (ctx *RenderContext) BuildNode(component Component) Node {
	if component == nil {
		return Node{}
	}

	switch c := component.(type) {
	case Fragment:
		list := []Node{}
		for _, e := range c {
			childNode := ctx.BuildNode(e)
			if childNode.Content == nil {
				if len(childNode.Children) != 0 {
					list = append(list, childNode.Children...)
				}
				continue
			}
			list = append(list, childNode)
		}
		return Node{Children: list}

	case Container:
		return c.BuildNode(ctx)
	case Control:
		return Node{Content: c}
	default:
		return ctx.BuildNode(component.Render(ctx))
	}
}

func (ctx *RenderContext) Make(render func(*RenderContext) Component) Node {
	subContext, _ := UseState(ctx, &RenderContext{
		content: makeRenderable(render),
		values:  make(map[int]any),
	})
	subContext.count = 0
	root := subContext.BuildNode(subContext.content)
	subContext.root = root
	return root
}

func printNodes(node *Node, indent int) {
	if len(node.Children) == 0 {
		fmt.Printf("%s<%T/>\n", strings.Repeat("  ", indent), node.Content)
		return
	}

	fmt.Printf("%s<%T>\n", strings.Repeat("  ", indent), node.Content)
	for _, child := range node.Children {
		printNodes(&child, indent+1)
	}
	fmt.Printf("%s</%T>\n", strings.Repeat("  ", indent), node.Content)
}

func (ctx *RenderContext) TriggerUpdate() {
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
		if ctx.root.Content == nil {
			// fmt.Printf("[%v] Root is nil, returning.\n", ctx)
			return
		}

		// fmt.Println("STATE VALUES ******")
		// for i := 0; i < ctx.count; i++ {
		// 	fmt.Printf("%02d -> %v\n", i, ctx.values[i])
		// }
		// fmt.Println("*******************")

		// fmt.Printf("[%v] RENDER TRIGGERED!\n", ctx)
		ctx.count = 0
		// fmt.Println("**** RENDER STARTING ****")
		newTree := ctx.BuildNode(ctx.content)
		// log.Printf("render time: %s\n", time.Now().Sub(now))
		// fmt.Println("**** RENDER DONE ****")

		// fmt.Printf("[%v] Old tree: %+v\n", ctx, ctx.root)
		// printNodes(&ctx.root, 0)
		// fmt.Printf("[%v] New tree: %+v\n", ctx, newTree)
		// printNodes(&newTree, 0)

		ctx.root.Update(newTree, nil)
	})
}
