package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("TextField example", 400, 300)

	textField := gocoa.NewTextField(25, 250, 350, 25)
	textField.SetFontFamily("Courier New")
	textField.SetFontSize(20)
	textField.SetBackgroundColor("#EDE9C0FF") // #RGBA
	textField.SetColor("#0000FFFF")
	textField.SetBorderColor("#00000030")
	textField.SetBorderWidth(2)
	textField.SetEnabled(true)
	textField.SetEditable(true)
	wnd.AddTextField(textField)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
