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

void* Centered_Window_New(int goWindowID, int width, int height, const char* title) 
{
    NSRect windowRect = NSMakeRect(0, 0, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window autorelease];
    CGFloat xPos = NSWidth([[window screen] frame])/2 - NSWidth([window frame])/2;
    CGFloat yPos = NSHeight([[window screen] frame])/2 - NSHeight([window frame])/2;
    gocoa_windowDelegate = [[WindowDelegate alloc] init];
    [gocoa_windowDelegate setGoWindowID:goWindowID];
    [window setDelegate:gocoa_windowDelegate];
    [window setFrame:NSMakeRect(xPos, yPos, NSWidth([window frame]), NSHeight([window frame])) display:YES];
    
    NSButton *button = [window standardWindowButton:NSWindowMiniaturizeButton];
    [button setEnabled: NO];

    return window;
}

int Screen_Center_X(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat xPos = NSWidth([[window screen] frame])/2 - NSWidth([window frame])/2;
    return (int)(xPos);
}

int Screen_Center_Y(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat yPos = NSHeight([[window screen] frame])/2 - NSHeight([window frame])/2;
    return (int)(yPos);
}

int Screen_X(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat xPos = NSWidth([[window screen] frame]);
    return (int)(xPos);
}

int Screen_Y(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    CGFloat yPos = NSHeight([[window screen] frame]);
    return (int)(yPos);
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

void Window_SetMaximumSize(void *wndPtr, int maxWidth, int maxHeight) {
    NSWindow* window = (NSWindow*)wndPtr;
    [window setMaxSize:NSMakeSize(maxWidth, maxHeight)];
}

void Window_SetMinimumSize(void *wndPtr, int minWidth, int minHeight) {
    NSWindow* window = (NSWindow*)wndPtr;
    [window setMinSize:NSMakeSize(minWidth, minHeight)];
}
