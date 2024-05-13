# Spot

Spot is a simple, cross-platform, reactive GUI toolkit for Go using native
widgets where available. It is designed to be easy to use and to provide a
consistent API across different platforms.

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	ui.Init()

	root := spot.Make(func(ctx *spot.RenderContext) spot.Component {
		counter, setCounter := spot.UseState[int](ctx, 0)

		buttonTitle := "Click me!"
		if counter > 0 {
			buttonTitle = fmt.Sprintf("Clicked %d times!", counter)
		}

		return &ui.Window{
			Title:  "Hello World!",
			Width:  200,
			Height: 125,
			Children: []spot.Component{
				&ui.Button{
					X: 25, Y: 50, Width: 150, Height: 25,
					Title: buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
			},
		}
	})

	root.Mount()
	ui.Run()
}
```

## Features

## FAQs

#### What does "reactive" mean?

In the context of Spot, _reactive_ means that the UI is automatically updated
when the state of the application changes. This is achieved by re-building an
immutable component tree upon state changes which can quickly be compared to
the previous state in order to determine what UI controls need to be updated.
In the web world, this idea is often called a "virtual DOM" and Spot actually
started as an experiment to bring this concept to Go by implementing a
React-like GUI library for the desktop.

By using a reactive model, the developer does not need to worry about updating
the UI manually. Instead, the developer can focus on the application logic and
let Spot take care of updating the UI.

#### What are the "native widgets" that Spot uses?

Currently, Spot uses a Cocoa backend on macOS and a FLTK-based one on all other
platforms. Optionally, FLTK can be used on the Mac, too. Better support for
Windows is planned for the future.

## Features, Spot does not have right now

- Automatic layouting
- Multiple windows
- Modal dialogs
- Resizable windows
- Menu bars
- Custom widgets
- Access to native widgets
- Drag and drop
- Internationalization

## List of supported UI controls

| Name      | Description                                              | Status                             | Cocoa                                         | FLTK                                                                        | GTK | Win32 |
| --------- | -------------------------------------------------------- | ---------------------------------- | --------------------------------------------- | --------------------------------------------------------------------------- | --- | ----- |
| Window    | Control representing a (top-level) window on the screen. | Done                               | NSWindow                                      | [Fl_Window](https://www.fltk.org/doc-1.4/classFl__Window.html)              | TBD | TBD   |
| Button    | Simple button to initiate an action.                     | Done                               | NSButton                                      | [Fl_Button](https://www.fltk.org/doc-1.4/classFl__Button.html)              | TBD | TBD   |
| Label     | A label control                                          | ?                                  | NSTextField                                   | [Fl_Output](https://www.fltk.org/doc-1.4/classFl__Output.html)              | TBD | TBD   |
| Text      | A text control                                           | ?                                  | NSTextView                                    | Text                                                                        | TBD | TBD   |
| TextInput | Control for single-line text input                       | ?                                  | NSTextField                                   | [Fl_Input](https://www.fltk.org/doc-1.4/classFl__Input.html)                | TBD | TBD   |
| Slider    | Horizontal slider input control                          | Done                               | NSSlider                                      | [Fl_Slider](https://www.fltk.org/doc-1.4/classFl__Slider.html)              | TBD | TBD   |
| Dial      | Circular status control                                  | Currently not supported by go-fltk | NSProgressIndicator (with `NSCircular` style) | [Fl_Dial](https://www.fltk.org/doc-1.4/classFl__Dial.html)                  | TBD | TBD   |
| Image     | An image control                                         | Not started                        | NSImageView                                   | Image                                                                       | TBD | TBD   |
| Dropdown  | A dropdown control                                       | Done                               | NSPopUpButton                                 | [Fl_Choice](https://www.fltk.org/doc-1.4/classFl__Choice.html)              | TBD | TBD   |
| ComboBox  | A combo box control                                      | Not started                        | NSComboBox                                    | ComboBox                                                                    | TBD | TBD   |
| Progress  | A progress control                                       | Done                               | NSProgress                                    | Progress                                                                    | TBD | TBD   |
| List      | A list control                                           | Not started                        | NSTableView                                   | List                                                                        | TBD | TBD   |
| Spinner   | Number input control with up/down buttons                | FLTK only                          | NSStepper                                     | [Fl_Spinner](https://www.fltk.org/doc-1.4/classFl__Spinner.html)            | TBD | TBD   |
| Checkbox  | A checkbox control                                       | Done                               | NSButton                                      | [Fl_Check_Button](https://www.fltk.org/doc-1.4/classFl__Check__Button.html) | TBD | TBD   |

## Potential future backends to look at

- Native Windows controls: https://github.com/rodrigocfd/windigo
