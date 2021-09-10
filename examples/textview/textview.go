package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("TextView example", 400, 300)

	textView := gocoa.NewTextView(25, 125, 150, 150)
	textView.SetText("hello world")
	wnd.AddTextView(textView)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
