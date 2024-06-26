#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextViewHandler : NSObject

@property(assign) int goTextViewId;

@end

typedef void *TextViewPtr;

TextViewPtr TextView_New(int goTextViewId, int x, int y, int w, int h);
const char *TextView_Text(TextViewPtr textViewPtr);
void TextView_SetText(TextViewPtr textViewPtr, const char *text);
void TextView_Remove(TextViewPtr textViewPtr);
void TextView_SetFontSize(TextViewPtr textViewPtr, int size);
const int TextView_Editable(TextViewPtr ptr);
void TextView_SetEditable(TextViewPtr ptr, int editable);