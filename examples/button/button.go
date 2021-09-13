package main

import (
	"fmt"

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

	checkbox := gocoa.NewButton(200, 40, 150, 150)
	checkbox.SetTitle("check me!")
	checkbox.SetButtonType(gocoa.ButtonTypeSwitch)
	checkbox.OnClick(func() {
		if checkbox.State() == gocoa.ButtonStateValueOn {
			fmt.Println("checked")
		} else {
			fmt.Println("unchecked")
		}
	})
	wnd.AddButton(checkbox)

	radiobutton := gocoa.NewButton(200, 125, 150, 150)
	radiobutton.SetTitle("radiobutton 1")
	radiobutton.SetState(gocoa.ButtonStateValueOn)
	radiobutton.SetButtonType(gocoa.ButtonTypeRadio)
	radiobutton.OnClick(func() {
		if radiobutton.State() == gocoa.ButtonStateValueOn {
			fmt.Println("checked")
		} else {
			fmt.Println("unchecked")
		}
	})
	wnd.AddButton(radiobutton)

	radiobutton2 := gocoa.NewButton(200, 100, 150, 150)
	radiobutton2.SetTitle("radiobutton 2")
	radiobutton2.SetButtonType(gocoa.ButtonTypeRadio)
	radiobutton2.OnClick(func() {
		if radiobutton2.State() == gocoa.ButtonStateValueOn {
			fmt.Println("checked")
		} else {
			fmt.Println("unchecked")
		}
	})
	wnd.AddButton(radiobutton2)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
