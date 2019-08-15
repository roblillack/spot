#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface TextFieldHandler : NSObject

@property (assign) int goTextFieldId;

@end

typedef void* pTextField;

pTextField TextField_New(int goTextFieldId, int x, int y, int w, int h);
