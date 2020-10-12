#import "textfield.h"
#include "_cgo_export.h"

@implementation TextFieldHandler
@end

TextFieldPtr TextField_New(int goTextFieldId, int x, int y, int w, int h) {
	/* create the NSTextField and add it to the window */
	id nsTextField = [[[NSTextField alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (TextFieldPtr)nsTextField;
}

const char* TextField_StringValue(TextFieldPtr textFieldPtr) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	return [[textField stringValue] cStringUsingEncoding:NSISOLatin1StringEncoding];
}

void TextField_SetStringValue(TextFieldPtr textFieldPtr, const char* text) {
	NSTextField* textField = (NSTextField*)textFieldPtr;
	[textField setStringValue:[NSString stringWithUTF8String:text]];
}