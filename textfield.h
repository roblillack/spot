#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextFieldHandler : NSObject

@property (assign) int goTextFieldId;

@end

typedef void* TextFieldPtr;

TextFieldPtr TextField_New(int goTextFieldId, int x, int y, int w, int h);

const char* TextField_StringValue(TextFieldPtr textFieldPtr);
void TextField_SetStringValue(TextFieldPtr textFieldPtr, const char* text);
const int TextField_Enabled(TextFieldPtr textFieldPtr);
void TextField_SetEnabled(TextFieldPtr textFieldPtr, int enabled);
const int TextField_Editable(TextFieldPtr textFieldPtr);
void TextField_SetEditable(TextFieldPtr textFieldPtr, int editable);
void TextField_SetFontFamily(TextFieldPtr textFieldPtr, const char* fontFamily);
void TextField_SetFontSize(TextFieldPtr textFieldPtr, const int fontSize);
void TextField_SetColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBackgroundColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBorderColor(TextFieldPtr textFieldPtr, const int r, const int g, const int b, const int a);
void TextField_SetBorderWidth(TextFieldPtr textFieldPtr, const int borderWidth);
