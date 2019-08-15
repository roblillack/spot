#import "textfield.h"
#include "_cgo_export.h"

@implementation TextFieldHandler
@end

pTextField TextField_New(int goTextFieldId, int x, int y, int w, int h) {
	/* create the NSTextField and add it to the window */
	NSTextField *nsTextField = [[NSTextField alloc] initWithFrame:NSMakeRect(x, y, w, h)];

	return (pTextField)nsTextField;
}
