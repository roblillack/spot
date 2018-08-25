#import "button.h"
#include "_cgo_export.h"

void* Button_New(int x, int y, int w, int h) {
    id button = [[[NSButton alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
    [button setTarget:NSApp];
    [button setAction:@selector(terminate:)];

    [button setButtonType:NSMomentaryLightButton];
    [button setBezelStyle:NSRoundedBezelStyle];
    return button;
}

void Button_SetTitle(void *btnPtr, const char* title) {
    NSButton* button = (NSButton*)btnPtr;
    [button setTitle:[NSString stringWithUTF8String:title]];
}