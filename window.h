#import <Cocoa/Cocoa.h>

#include "button.h" // because we have specific Button code for now

void* Window_New(int x, int y, int width, int height, const char* title);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_AddButton(void *wndPtr, ButtonPtr btnPtr);
