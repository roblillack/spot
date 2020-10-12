#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextViewHandler : NSObject

@property (assign) int goTextViewId;

@end

typedef void* TextViewPtr;

TextViewPtr TextView_New(int goTextViewId, int x, int y, int w, int h);
void TextView_SetText(TextViewPtr textViewPtr, const char* text);
