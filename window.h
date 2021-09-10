#import <Cocoa/Cocoa.h>

#include "button.h"
#include "imageview.h"
#include "textview.h"
#include "textfield.h"
#include "progressIndicator.h"

void* Window_New(int goWindowID, int x, int y, int width, int height, const char* title);
void* Centered_Window_New(int goWindowID, int width, int height, const char* title);
int Screen_Center_X(void *wndPtr);
int Screen_Center_Y(void *wndPtr);
int Screen_X(void *wndPtr);
int Screen_Y(void *wndPtr);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr);
void Window_AddTextView(void *wndPtr, TextViewPtr tvPtr);
void Window_AddTextField(void *wndPtr, TextFieldPtr tfPtr);
void Window_AddProgressIndicator(void *wndPtr, ProgressIndicatorPtr progressIndicatorPtr);
void Window_AddImageView(void *wndPtr, ImageViewPtr imageViewPtr);
void Window_Update(void *wndPtr);
void Window_SetTitle(void *wndPtr, const char* title);
void Window_AddDefaultQuitMenu(void *wndPtr);
