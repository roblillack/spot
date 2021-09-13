#import "button.h"
#include "_cgo_export.h"

@implementation ButtonHandler
-(void) buttonClicked:(id) sender
{
    onButtonClicked([self goButtonID]);
}
@end

ButtonPtr Button_New(int goButtonID, int x, int y, int w, int h) {
    id nsButton = [[[NSButton alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
    ButtonHandler* handler = [[ButtonHandler alloc] init];
    [handler setGoButtonID:goButtonID];
    [handler autorelease];
    [nsButton setTarget:handler];
    [nsButton setAction:@selector(buttonClicked:)];

    [nsButton setButtonType:NSMomentaryLightButton];
    [nsButton setBezelStyle:NSRoundedBezelStyle];

    return (ButtonPtr)nsButton;
}

void Button_SetTitle(ButtonPtr btnPtr, const char* title) {
    NSButton* button = (NSButton*)btnPtr;
    [button setTitle:[NSString stringWithUTF8String:title]];
}

const char* Button_Title(ButtonPtr btnPtr) {
    NSButton* button = (NSButton*)btnPtr;
    return [[button title] cStringUsingEncoding:NSASCIIStringEncoding];
}

void Button_SetButtonType(ButtonPtr btnPtr, int buttonType) {
    NSButton* button = (NSButton*)btnPtr;
    [button setButtonType:buttonType];
}

void Button_SetBezelStyle(ButtonPtr btnPtr, int bezelStyle) {
    NSButton* button = (NSButton*)btnPtr;
    [button setBezelStyle:bezelStyle];
}

void Button_SetFontFamily(ButtonPtr btnPtr, const char* fontFamily) {
    NSButton* button = (NSButton*)btnPtr;
	button.font = [NSFont fontWithName:[NSString stringWithUTF8String:fontFamily] size:button.font.pointSize];
}

void Button_SetFontSize(ButtonPtr btnPtr, int fontSize) {
    NSButton* button = (NSButton*)btnPtr;
    button.font = [NSFont fontWithName:button.font.fontName size:fontSize];
}

void Button_SetColor(ButtonPtr btnPtr, int r, int g, int b, int a) {
    NSButton* button = (NSButton*)btnPtr;
    NSMutableAttributedString *attrTitle =[[NSMutableAttributedString alloc] initWithString:[NSString stringWithUTF8String:Button_Title(btnPtr)]];
    NSUInteger len = [attrTitle length];
    NSRange range = NSMakeRange(0, len);
    [attrTitle addAttribute:NSForegroundColorAttributeName value:[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f] range:range];
    [attrTitle addAttribute:NSFontAttributeName value:button.font range:range];
    [attrTitle fixAttributesInRange:range];
    [button setAttributedTitle:attrTitle];
}

void Button_SetBackgroundColor(ButtonPtr btnPtr, int r, int g, int b, int a) {
    NSButton* button = (NSButton*)btnPtr;
    [button setBordered:false]; // required, otherwise can't set background
    [[button cell] setBackgroundColor:[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f]];
}

void Button_SetBorderColor(ButtonPtr btnPtr, int r, int g, int b, int a) {
    NSButton* button = (NSButton*)btnPtr;
    button.wantsLayer = true;
	button.layer.borderColor = [[NSColor colorWithCalibratedRed:r/255.f green:g/255.f blue:b/255.f alpha:a/255.f] CGColor];
}

void Button_SetBorderWidth(ButtonPtr btnPtr, int borderWidth) {
    NSButton* button = (NSButton*)btnPtr;
    button.wantsLayer = true;
	button.layer.borderWidth = borderWidth;
}
