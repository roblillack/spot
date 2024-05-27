#include "_cgo_export.h"
#import "button.h"

@implementation CustomButton
- (void)mouseDown:(NSEvent *)theEvent {
  NSPoint p = [self convertPoint:theEvent.locationInWindow fromView:nil];
  onCustomButtonClicked([self goButtonID], p.x, p.y,
                        theEvent.buttonNumber != 1);
}
@end

ButtonPtr CustomButton_New(int goButtonID, int x, int y, int w, int h) {
  CustomButton *nsButton =
      [[[CustomButton alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

  [nsButton setGoButtonID:goButtonID];
  [nsButton setButtonType:NSButtonTypeMomentaryLight];
  [nsButton setBezelStyle:NSBezelStyleRounded];

  return (ButtonPtr)nsButton;
}

void Window_AddCustomButton(void *wndPtr, ButtonPtr btnPtr) {
  CustomButton *button = (CustomButton *)btnPtr;
  NSWindow *window = (NSWindow *)wndPtr;
  [[window contentView] addSubview:button];
}
