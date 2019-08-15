#import "window.h"
#include "_cgo_export.h"

void* Window_New(int x, int y, int width, int height, const char* title) 
{
    NSRect windowRect = NSMakeRect(x, y, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window autorelease];
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

void Window_AddTextView(void *wndPtr, pTextView ptv)
{
    NSTextView* textview = (NSTextView*)ptv;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:indicator];
}

void Window_AddTextField(void *wndPtr, pTextField ptf)
{
    NSTextField* textfield = (NSTextField*)ptf;
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