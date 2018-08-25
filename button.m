#import "button.h"
#include "_cgo_export.h"

void* Button_New() {
    id button = [[[NSButton alloc] initWithFrame:NSMakeRect(20, 20, 280, 60)] autorelease];
    [button setTarget:NSApp];
    [button setAction:@selector(terminate:)];
    [button setTitle:@"Quit!"];

    [button setButtonType:NSMomentaryLightButton];
    [button setBezelStyle:NSRoundedBezelStyle];
    return button;
}

void Button_SetTitle(void *btnPtr, const char* title) {
    NSButton* button = (NSButton*)btnPtr;
    [button setTitle:[NSString stringWithUTF8String:title]];

    //[[window contentView] addSubview:button];
}