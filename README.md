# Spot UI

## What Spot does not offer at the moment

- Automatic layouting
- Custom widgets
- Drag and drop
- Internationalization

## List of supported UI controls

| Name      | Description                               | Status                             | Cocoa                                         | FLTK                                                                        | GTK | Win32 |
| --------- | ----------------------------------------- | ---------------------------------- | --------------------------------------------- | --------------------------------------------------------------------------- | --- | ----- |
| Window    | A window control                          | Done                               | NSWindow                                      | [Fl_Window](https://www.fltk.org/doc-1.4/classFl__Window.html)              | TBD | TBD   |
| Button    | A button control                          | Done                               | NSButton                                      | [Fl_Button](https://www.fltk.org/doc-1.4/classFl__Button.html)              | TBD | TBD   |
| Label     | A label control                           | ?                                  | NSTextField                                   | [Fl_Output](https://www.fltk.org/doc-1.4/classFl__Output.html)              | TBD | TBD   |
| Text      | A text control                            | ?                                  | NSTextView                                    | Text                                                                        | TBD | TBD   |
| TextInput | Control for single-line text input        | ?                                  | NSTextField                                   | [Fl_Input](https://www.fltk.org/doc-1.4/classFl__Input.html)                | TBD | TBD   |
| Slider    | Horizontal slider input control           | Done                               | NSSlider                                      | [Fl_Slider](https://www.fltk.org/doc-1.4/classFl__Slider.html)              | TBD | TBD   |
| Dial      | Circular status control                   | Currently not supported by go-fltk | NSProgressIndicator (with `NSCircular` style) | [Fl_Dial](https://www.fltk.org/doc-1.4/classFl__Dial.html)                  | TBD | TBD   |
| Image     | An image control                          | Not started                        | NSImageView                                   | Image                                                                       | TBD | TBD   |
| Dropdown  | A dropdown control                        | Done                               | NSPopUpButton                                 | [Fl_Choice](https://www.fltk.org/doc-1.4/classFl__Choice.html)              | TBD | TBD   |
| ComboBox  | A combo box control                       | Not started                        | NSComboBox                                    | ComboBox                                                                    | TBD | TBD   |
| Progress  | A progress control                        | Done                               | NSProgress                                    | Progress                                                                    | TBD | TBD   |
| List      | A list control                            | Not started                        | NSTableView                                   | List                                                                        | TBD | TBD   |
| Spinner   | Number input control with up/down buttons | FLTK only                          | NSStepper                                     | [Fl_Spinner](https://www.fltk.org/doc-1.4/classFl__Spinner.html)            | TBD | TBD   |
| Checkbox  | A checkbox control                        | FTLK only                          | NSButton                                      | [Fl_Check_Button](https://www.fltk.org/doc-1.4/classFl__Check__Button.html) | TBD | TBD   |

## Potential future backends to look at

- Native Windows controls: https://github.com/rodrigocfd/windigo
