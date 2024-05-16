package main

import (
	"fmt"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func StateLessComponent() spot.Component {
	return &ui.Button{
		X: 10, Y: 25, Width: 230, Height: 25,
		Title: "Stateless button",
	}
}

func StatefulNoProps(ctx *spot.RenderContext) spot.Component {
	counter, setCounter := spot.UseState[int](ctx, 0)

	title := "Stateful button"
	if counter > 0 {
		title = fmt.Sprintf("Clicked Stateful %dx", counter)
	}

	return &ui.Button{
		X: 10, Y: 65, Width: 230, Height: 25,
		Title:   title,
		OnClick: func() { setCounter(counter + 1) },
	}
}

func StatefulWithProps(ctx *spot.RenderContext, initialTitle string) spot.Component {
	counter, setCounter := spot.UseState[int](ctx, 0)

	title := initialTitle
	if counter > 0 {
		title = fmt.Sprintf("Clicked %s %dx", initialTitle, counter)
	}

	return &ui.Button{
		X: 10, Y: 105, Width: 230, Height: 25,
		Title:   title,
		OnClick: func() { setCounter(counter + 1) },
	}
}

type StructComponent struct {
	X, Y, Width, Height int
	Title               string
}

func (r *StructComponent) Render(ctx *spot.RenderContext) spot.Component {
	fmt.Println("Rendering StructComponent")
	counter, setCounter := spot.UseState[int](ctx, 0)

	title := r.Title
	if counter > 0 {
		title = fmt.Sprintf("Clicked %s %dx", r.Title, counter)
	}
	return &ui.Button{
		X: r.X, Y: r.Y, Width: r.Width, Height: r.Height,
		Title: title, OnClick: func() { setCounter(counter + 1) },
	}
}

func main() {
	ui.Init()

	root := spot.Render(spot.Make(func(ctx *spot.RenderContext) spot.Component {
		return &ui.Window{
			Title:  "Custom components in Spot",
			Width:  250,
			Height: 200,
			Children: []spot.Component{
				StateLessComponent(),
				StatefulNoProps(ctx),
				StatefulWithProps(ctx, "Stateful w/ props"),
				&StructComponent{
					X: 10, Y: 145, Width: 230, Height: 25,
					Title: "Struct component",
				},
			},
		}
	}))

	fmt.Println("Attempting to mount.")

	root.Mount(nil)

	ui.Run()
}
