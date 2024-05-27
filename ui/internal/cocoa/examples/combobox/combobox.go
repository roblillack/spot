package main

import (
	"fmt"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("ComboBox example", 400, 300)

	combobox := gocoa.NewComboBox(25, 125, 150, 25)
	combobox.AddItem("hello")
	combobox.AddItem("world")
	combobox.SetSelectedIndex(0)
	combobox.SetEditable(false)
	combobox.OnSelectionDidChange(func() {
		fmt.Println(combobox.SelectedText())
		fmt.Println(combobox.SelectedIndex())
	})
	wnd.AddComboBox(combobox)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
