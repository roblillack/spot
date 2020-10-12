#import "textview.h"
#include "_cgo_export.h"

@implementation TextViewHandler
@end

TextViewPtr TextView_New(int goTextViewId, int x, int y, int w, int h) {
	/* create the NSTextView and add it to the window */
	id nsTextView = [[[NSTextView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (TextViewPtr)nsTextView;
}

void TextView_SetText(TextViewPtr textViewPtr, const char* text) {
	NSTextView* tv = (NSTextView*)textViewPtr;
	[tv setString:[NSString stringWithUTF8String:text]];
}