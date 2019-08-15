#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextViewHandler : NSObject

@property (assign) int goTextViewId;

@end

typedef void* pTextView;

pTextView TextView_New(int goTextViewId, int x, int y, int w, int h);
