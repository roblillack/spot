#import "textview.h"
#include "_cgo_export.h"

@implementation TextViewHandler
@end

pTextView TextView_New(int goTextViewId, int x, int y, int w, int h) {
	/* create the NSTextView and add it to the window */
	NSTextView *nsTextView = [[NSTextView alloc] initWithFrame:NSMakeRect(x, y, w, h)];

	return (pTextView)nsTextView;
}
