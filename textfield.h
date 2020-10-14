#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextFieldHandler : NSObject

@property (assign) int goTextFieldId;

@end

typedef void* TextFieldPtr;

TextFieldPtr TextField_New(int goTextFieldId, int x, int y, int w, int h);

const char* TextField_StringValue(TextFieldPtr textFieldPtr);
void TextField_SetStringValue(TextFieldPtr textFieldPtr, const char* text);
