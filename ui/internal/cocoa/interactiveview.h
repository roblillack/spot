#import "image.h"
#import <Cocoa/Cocoa.h>

@interface InteractiveView : NSButton
- (void)mouseDown:(NSEvent *)theEvent;
@end

typedef void *InteractiveViewPtr;

InteractiveViewPtr InteractiveView_New(int x, int y, int w, int h);
void InteractiveView_SetImage(InteractiveViewPtr ptr, ImagePtr imagePtr);
void InteractiveView_Remove(InteractiveViewPtr ptr);

void Window_AddInteractiveView(void *wndPtr, InteractiveViewPtr ptr);
