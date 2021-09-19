package main

import (
	"fmt"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()
	gocoa.OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd := gocoa.NewCenteredWindow("Window", 300, 200)
	fmt.Printf("created window: %v\n", wnd)
	fmt.Printf("screen size %v\n", wnd.GetScreen())
	wnd.OnDidMove(func(uwnd *gocoa.Window) {
		fmt.Printf("old: %v\nnew: %v\n", wnd, uwnd)
	})

	wnd.SetCloseButtonEnabled(false)
	wnd.SetZoomButtonEnabled(false)
	wnd.SetMiniaturizeButtonEnabled(false)
	wnd.SetAllowsResizing(false)

	// Quit button
	quitButton := gocoa.NewButton(75, 25, 150, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() {
		gocoa.TerminateApplication()
	})
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()

	gocoa.RunApplication()

}
