#import "window.h"
#import "windowdelegate.h"
#include "_cgo_export.h"

WindowDelegate *gocoa_windowDelegate = nil;

void* Window_New(int goWindowID, int x, int y, int width, int height, const char* title) 
{
    NSRect windowRect = NSMakeRect(x, y, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window autorelease];

    gocoa_windowDelegate = [[WindowDelegate alloc] init];
    [gocoa_windowDelegate setGoWindowID:goWindowID];
    [window setDelegate:gocoa_windowDelegate];
    return window;
}

void Window_MakeKeyAndOrderFront(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window makeKeyAndOrderFront:nil];
}

void Window_AddButton(void *wndPtr, ButtonPtr btnPtr) 
{
    NSButton* button = (NSButton*)btnPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:button];
}

void Window_AddTextView(void *wndPtr, TextViewPtr tvPtr)
{
    NSTextView* textview = (NSTextView*)tvPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:textview];
}

void Window_AddTextField(void *wndPtr, TextFieldPtr tfPtr)
{
    NSTextField* textfield = (NSTextField*)tfPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:textfield];
}

void Window_AddProgressIndicator(void *wndPtr, ProgressIndicatorPtr progressIndicatorPtr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)progressIndicatorPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:indicator];
}

void Window_Update(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] setNeedsDisplay:YES];
}

void Window_SetTitle(void *wndPtr, const char* title)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setTitle:[NSString stringWithUTF8String:title]];
}
