#import "button.h"
#import "image.h"

@interface CustomButton : NSButton {
}
@property(assign) int goButtonID;
- (void)mouseDown:(NSEvent *)theEvent;
@end

ButtonPtr CustomButton_New(int goButtonID, int x, int y, int w, int h);
void Window_AddCustomButton(void *wndPtr, ButtonPtr btnPtr);
