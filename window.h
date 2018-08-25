#import <Cocoa/Cocoa.h>

void* Window_New(int x, int y, int width, int height, const char* title);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_AddButton(void *wndPtr, void *btnPtr);
