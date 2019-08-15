#import <Cocoa/Cocoa.h>

#include "button.h"
#include "textview.h"
#include "textfield.h"
#include "progressIndicator.h"

void* Window_New(int goWindowID, int x, int y, int width, int height, const char* title);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr);
void Window_AddTextView(void *wndPtr, pTextView ptv);
void Window_AddTextField(void *wndPtr, pTextField ptf);
void Window_AddProgressIndicator(void *wndPtr, ProgressIndicatorPtr progressIndicatorPtr);
void Window_Update(void *wndPtr);
