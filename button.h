#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface ButtonHandler : NSObject

@property (assign) int goButtonID;
-(void) buttonClicked:(id) sender;

@end

typedef void* ButtonPtr;

ButtonPtr Button_New(int goButtonID, int x, int y, int w, int h);
void Button_SetTitle(ButtonPtr btnPtr, const char* title);
