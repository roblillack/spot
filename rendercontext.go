package spot

import (
	"fmt"
	"strings"
)

type RenderContext struct {
	root    Component
	render  func(ctx *RenderContext) Component
	changed bool
	values  map[int]any
	count   int
}

func (ctx *RenderContext) Make(render func(*RenderContext) Component) Component {
	subContext, _ := UseState(ctx, &RenderContext{
		render: render,
		values: make(map[int]any),
	})
	subContext.count = 0
	root := render(subContext)
	subContext.root = root
	return root
}

func printComponentTree(c Component, indent int) {
	w, ok := c.(Container)
	if !ok {
		fmt.Printf("%s<%T/>\n", strings.Repeat("  ", indent), c)
		return
	}

	fmt.Printf("%s<%T>\n", strings.Repeat("  ", indent), c)
	for _, child := range w.GetChildren() {
		printComponentTree(child, indent+1)
	}
	fmt.Printf("%s</%T>\n", strings.Repeat("  ", indent), c)
}

func (ctx *RenderContext) TriggerUpdate() {
	if ctx.root == nil {
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
		oldTree := ctx.root
		fmt.Println("**** RENDER STARTING ****")
		newTree := ctx.render(ctx)
		fmt.Println("**** RENDER DONE ****")

		fmt.Printf("[%v] Old tree: %+v\n", ctx, oldTree)
		printComponentTree(oldTree, 0)
		fmt.Printf("[%v] New tree: %+v\n", ctx, newTree)
		printComponentTree(newTree, 0)

		if !oldTree.Equals(newTree) {
			// fmt.Printf("[%v] Updating tree!\n", ctx)
			// fmt.Printf("[%v] On main thread here.\n", ctx)
			oldTree.Update(newTree)
			// fmt.Printf("[%v] Updating tree done.\n", ctx)
		}

		ctx.changed = false
	})
}
