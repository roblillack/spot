package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("Label example", 400, 300)

	label := gocoa.NewLabel(25, 125, 300, 150)
	label.SetStringValue("Labels")
	label.SetFontFamily("Helvetica")
	label.SetFontSize(24)
	wnd.AddLabel(label)

	label2 := gocoa.NewLabel(25, 85, 300, 150)
	label2.SetStringValue("All your labels are belong to us!")
	label2.SetFontSize(12)
	wnd.AddLabel(label2)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
