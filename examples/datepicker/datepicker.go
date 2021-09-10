package main

import (
	"fmt"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("DatePicker example", 400, 300)

	datePicker := gocoa.NewDatePicker(25, 125, 150, 150)
	datePicker.SetStyle(gocoa.DatePickerStyleClockAndCalendar)
	datePicker.SetMode(gocoa.DatePickerModeSingle)
	datePicker.SetDateFormat("YYYY-MM-dd")
	datePicker.SetDate("2022-02-03")
	datePicker.SetMinimumDate("2022-02-01")
	datePicker.SetMaximumDate("2022-02-15")
	wnd.AddDatePicker(datePicker)

	datePicker2 := gocoa.NewDatePicker(200, 125, 150, 150)
	datePicker2.SetStyle(gocoa.DatePickerStyleTextField)
	datePicker2.SetDateFormat("dd/MM YYYY")
	datePicker2.SetDate("05/11 2011")
	wnd.AddDatePicker(datePicker2)

	fmt.Println(datePicker.Date())

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
