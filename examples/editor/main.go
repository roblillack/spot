package main

import (
	"flag"
	"log"
	"os"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	flag.Parse()

	var initialFilename string
	if flag.NArg() > 0 {
		initialFilename = flag.Arg(0)
	}

	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		filename, _ := spot.UseState(ctx, initialFilename)
		content, setContent := spot.UseState(ctx, "")
		spot.UseEffect(ctx, func() {
			if filename == "" {
				setContent("")
				return
			}
			log.Printf("Loading %s ...\n", filename)
			raw, _ := os.ReadFile(filename)
			setContent(string(raw))
		}, []any{filename})

		filenameMsg := "no file loaded"
		if filename != "" {
			filenameMsg = filename
		}

		return &ui.Window{
			Title: "Spot Edit — " + filenameMsg,
			Width: 800, Height: 600,
			Children: []spot.Component{
				&ui.Label{
					X: 10, Y: 10, Width: 780, Height: 20,
					Value: filenameMsg,
				},

				&ui.TextEditor{
					X: 10, Y: 30, Width: 780, Height: 560,
					Text:     content,
					OnChange: setContent,
				},
			},
		}
	})

	ui.Run()
}
