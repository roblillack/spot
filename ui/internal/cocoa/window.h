#import <Cocoa/Cocoa.h>

#include "button.h"
#include "combobox.h"
#include "datepicker.h"
#include "imageview.h"
#include "progressIndicator.h"
#include "slider.h"
#include "stepper.h"
#include "tableview.h"
#include "textfield.h"
#include "textview.h"

void *Window_New(int goWindowID, int x, int y, int width, int height,
                 const char *title);
void *Centered_Window_New(int goWindowID, int width, int height,
                          const char *title);
int Screen_Center_X(void *wndPtr);
int Screen_Center_Y(void *wndPtr);
int Screen_X(void *wndPtr);
int Screen_Y(void *wndPtr);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr);
void Window_AddDatePicker(void *wndPtr, DatePickerPtr datePickerPtr);
void Window_AddTextView(void *wndPtr, TextViewPtr tvPtr);
void Window_AddTextField(void *wndPtr, TextFieldPtr tfPtr);
void Window_AddProgressIndicator(void *wndPtr,
                                 ProgressIndicatorPtr progressIndicatorPtr);
void Window_AddImageView(void *wndPtr, ImageViewPtr imageViewPtr);
void Window_AddSlider(void *wndPtr, SliderPtr sliderPtr);
void Window_AddStepper(void *wndPtr, StepperPtr ptr);
void Window_AddComboBox(void *wndPtr, ComboBoxPtr comboBoxPtr);
void Window_AddTableView(void *wndPtr, TableViewPtr tableViewPtr);
void Window_Update(void *wndPtr);
void Window_SetTitle(void *wndPtr, const char *title);
void Window_SetMiniaturizeButtonEnabled(void *wndPtr, int enabled);
void Window_SetZoomButtonEnabled(void *wndPtr, int enabled);
void Window_SetCloseButtonEnabled(void *wndPtr, int enabled);
void Window_SetAllowsResizing(void *wndPtr, int enabled);
void Window_AddDefaultQuitMenu(void *wndPtr);
