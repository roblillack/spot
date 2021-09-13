package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("Button example", 400, 300)

	myButton := gocoa.NewButton(25, 125, 150, 150)
	myButton.SetTitle("My Button")
	myButton.SetBackgroundColor("#FF000033")
	myButton.SetFontFamily("Courier New")
	myButton.SetFontSize(24)
	myButton.SetBorderColor("#FF0000EE")
	myButton.SetBorderWidth(5)
	myButton.SetColor("#0000FF90")
	wnd.AddButton(myButton)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
