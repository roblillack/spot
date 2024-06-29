#import "interactiveview.h"
#include "_cgo_export.h"

@implementation InteractiveView
- (void)mouseDown:(NSEvent *)theEvent {
  NSPoint p = [self convertPoint:theEvent.locationInWindow fromView:nil];
  onInteractiveViewClicked((InteractiveViewPtr)self, p.x, p.y, false);
}
- (void)rightMouseDown:(NSEvent *)theEvent {
  NSPoint p = [self convertPoint:theEvent.locationInWindow fromView:nil];
  onInteractiveViewClicked((InteractiveViewPtr)self, p.x, p.y, true);
}
@end

InteractiveViewPtr InteractiveView_New(int x, int y, int w, int h) {
  InteractiveView *control = [[[InteractiveView alloc]
      initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
  [control setBordered:NO];
  [control setTitle:@""];
  // [control setBezelStyle:NSBezelStyleShadowlessSquare];
  // [control setButtonType:NSMomentaryChangeButton];
  [control setImageScaling:NSImageScaleAxesIndependently];
  // [nsButton setImageScaling:NSImageScaleProportionallyDown];

  return (InteractiveViewPtr)control;
}

void InteractiveView_SetImage(InteractiveViewPtr ptr, ImagePtr imagePtr) {
  NSImage *theImage = (NSImage *)imagePtr;
  InteractiveView *control = (InteractiveView *)ptr;
  [control setImage:theImage];
}

void InteractiveView_Remove(InteractiveViewPtr ptr) {
  InteractiveView *control = (InteractiveView *)ptr;
  [control removeFromSuperview];
}

void Window_AddInteractiveView(void *wndPtr, InteractiveViewPtr ptr) {
  InteractiveView *control = (InteractiveView *)ptr;
  NSWindow *window = (NSWindow *)wndPtr;
  [[window contentView] addSubview:control];
}
