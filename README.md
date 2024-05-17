<p style="text-align:center"><img src="./resources/2024-05-13-demo-video.gif" width="50%" height="50%" /></p>

# Spot

[![Go Reference](https://pkg.go.dev/badge/github.com/roblillack/spot.svg)](https://pkg.go.dev/github.com/roblillack/spot)
[![Go Report Card](https://goreportcard.com/badge/github.com/roblillack/spot)](https://goreportcard.com/report/github.com/roblillack/spot)

Spot is a simple, cross-platform, reactive GUI toolkit for Go using native
widgets where available. It is designed to be easy to use and to provide a
consistent API across different platforms.

## Example

```go
package main

import (
	"fmt"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
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

	ui.Run()
}
```

## Features

- **Simple**: You can add Spot as a simple dependency to your project and start
  building your UI right away. No need to use additional tools or code
  generation steps. Just write Go code and get a native GUI application as a
  self-contained binary.
- **Cross-platform**: Spot uses native widgets where available and
  automatically selects the best backend for the platform you are running on
  at compile time. Currently, two backend implementations are provided: one
  based on [FLTK](https://fltk.org) using
  [go-fltk](https://github.com/pwiecz/go-fltk) and one based on Cocoa using
  ([a modified version of](https://github.com/roblillack/gocoa))
  [gocoa](https://github.com/mojbro/gocoa).
- **Reactive**: Spot automatically updates the UI when the state of the
  application changes. You just provide side-effect free rendering functions
  and manage the state of your application using the [`UseState`](https://pkg.go.dev/github.com/roblillack/spot#UseState) hook.
- **Broad widget support**: Spot provides a wide range of UI controls out of
  the box, including buttons, labels, text inputs, sliders, dropdowns, and
  more. See the full list:
  [List of supported UI controls](#list-of-supported-ui-controls).

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

#### Can I implement my own hooks?

Yes, just like in React, you can implement your own hooks. Just create a
function which takes a `*spot.RenderContext` as first argument and use this to
"hook" into the Spot lifecycle by calling `spot.UseState`, `spot.UseEffect`,
etc. Convention here is to prefix the function with `Use…`.

#### How do I write custom components?

There are a few different ways to separate your UI into components in Spot;
for some ideas, check out the `custom-components` example. The main way to
write custom components is to create a struct that implements the
`spot.Component` interface. This interface has a single method,
`Render(ctx *spot.RenderContext) spot.Component`, which is called to render
the component. Components created like this can be used in the same way as
the built-in ones.

Look at the `BlinkingButton` component in the example to see how this is done.

#### Can I use Spot with a completely different widget library than the provided one?

Yes, you can. You just need to create some structs that implement the
`spot.Component` interface and which take care of managing the native widgets.

#### Can I use `spot/ui`, but with a different backend than Cocoa or FLTK?

Currently, these are the only backends that are supported. But feel free to
create a PR if you want to add support for another backend. _\*hint hint\*_

#### What's the difference between `spot/ui` and `spot`?

`spot` is the core package that provides the reactive model and the rendering
functionality. It is backend-agnostic and can be used with any set of controls
which implement the `spot.Control` interface.

`spot/ui` is a package that provides a set of pre-built cross-platform GUI
controls that which can be used with `spot`.

#### What's the difference between a “component” and a “control”?

In Spot, a _component_ is a logical unit of the application that contains
business logic and state. Any component is made out of other componens and
can ultimately be rendered down to a single or multiple "controls".

A _control_ is special kind component is mounted to the UI tree and represents
a visual element on the screen. Usually a control is backed by a native
implementation of the GUI backend, like a button, a label, or a text input.

#### What do the terms ”make”, “render”, “build”, “mount”, and “update” mean in the context of Spot?

- _Make_: The process of creating a new component instance. This is done by
  creating a reference to an instance of a struct that implements the
  `spot.Component` interface or by calling `spot.Make` with a render function.

- _Render_: The process of applying a component's state to its building blocks
  and hereby returning another component instance. This is done by calling the
  `Render` method on a component instance.

- _Build_: The process of creating a new UI tree from a component instance.
  This is done by _recursively_ rendering a component to create a tree of
  controls. This can be done by calling `spot.Build` with a component instance
  or `spot.BuildFn` with a render function.

- _Mount_: The process of creating real UI controls from a (virtual) tree of
  controls. This is done by calling `Mount` on a tree node or `spot.Mount` with
  a component instance or `spot.MountFn` with a render function.

- _Update_: The process of updating a tree of (mounted) controls. This is done
  by calling `Update` on a tree node.

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

Explanation of the status column: \
❓ Not implemented /
🚧 Work in progress /
⚠️ Partially implemented /
✅ Done

| Name                                                                        | Description                                                                   | Native controls used                                                                                          | Status      |
| --------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ----------- |
| [Button](https://pkg.go.dev/github.com/roblillack/spot/ui#Button)           | Simple button to initiate an action                                           | [Fl_Button](https://www.fltk.org/doc-1.4/classFl__Button.html) <br> NSButton                                  | ✅          |
| [Checkbox](https://pkg.go.dev/github.com/roblillack/spot/ui#Checkbox)       | Control offering the user a choice between two mutually exclusive options     | [Fl_Check_Button](https://www.fltk.org/doc-1.4/classFl__Check__Button.html) <br> NSButton                     | ✅          |
| ComboBox                                                                    | A combined dropdown menu with text input                                      | ComboBox <br> NSComboBox                                                                                      | Not started |
| [Dial](https://pkg.go.dev/github.com/roblillack/spot/ui#Dial)               | Circular status control                                                       | [Fl_Dial](https://www.fltk.org/doc-1.4/classFl__Dial.html) <br> NSProgressIndicator (with `NSCircular` style) | ⚠️          |
| [Dropdown](https://pkg.go.dev/github.com/roblillack/spot/ui#Dropdown)       | Drop-down menu to select a single item out of multiple options                | [Fl_Choice](https://www.fltk.org/doc-1.4/classFl__Choice.html) <br> NSPopUpButton                             | ✅          |
| Image                                                                       | An image control                                                              | Image <br> NSImageView                                                                                        | Not started |
| [Label](https://pkg.go.dev/github.com/roblillack/spot/ui#Label)             | Simple, non-editable text label                                               | [Fl_Output](https://www.fltk.org/doc-1.4/classFl__Output.html) <br> NSTextField                               | ✅          |
| [ProgressBar](https://pkg.go.dev/github.com/roblillack/spot/ui#ProgressBar) | Progress bar control to visualize the progression of a long-running operation | Progress <br> NSProgress                                                                                      | ✅          |
| ListBox                                                                     | A list control                                                                | List <br> NSTableView                                                                                         | 🚧          |
| [Slider](https://pkg.go.dev/github.com/roblillack/spot/ui#Slider)           | Horizontal slider input control                                               | [Fl_Slider](https://www.fltk.org/doc-1.4/classFl__Slider.html) <br> NSSlider                                  | ✅          |
| [Spinner](https://pkg.go.dev/github.com/roblillack/spot/ui#Spinner)         | Number input control with up/down buttons                                     | [Fl_Spinner](https://www.fltk.org/doc-1.4/classFl__Spinner.html) <br> NSStepper                               | ⚠️          |
| TextView/TextEditor                                                         | General-purpose text box to view/edit multi-line text content                 | Text <br> NSTextView                                                                                          | 🚧          |
| [TextField](https://pkg.go.dev/github.com/roblillack/spot/ui#TextField)     | Control for single-line text input                                            | [Fl_Input](https://www.fltk.org/doc-1.4/classFl__Input.html) <br> NSTextField                                 | ✅          |
| [Window](https://pkg.go.dev/github.com/roblillack/spot/ui#Window)           | Control representing a (top-level) window on the screen                       | [Fl_Window](https://www.fltk.org/doc-1.4/classFl__Window.html) <br> NSWindow                                  | ✅          |

## Potential future backends to look at

- Native Windows controls: https://github.com/rodrigocfd/windigo
