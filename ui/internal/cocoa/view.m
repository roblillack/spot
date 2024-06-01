#import "view.h"
#include "_cgo_export.h"

void View_SetFrameOrigin(ViewPtr ptr, int x, int y) {
  NSView *control = (NSView *)ptr;
  NSRect frame = control.frame;
  frame.origin.x = x;
  frame.origin.y = y;
  control.frame = frame;
}

void View_SetFrameSize(ViewPtr ptr, int w, int h) {
  NSView *control = (NSView *)ptr;
  NSRect frame = control.frame;
  frame.size.width = w;
  frame.size.height = h;
  control.frame = frame;
}

void View_SetFrame(ViewPtr ptr, int x, int y, int w, int h) {
  NSView *control = (NSView *)ptr;
  control.frame = NSMakeRect(x, y, w, h);
}

void View_Frame(ViewPtr ptr, int *x, int *y, int *w, int *h) {
  NSView *control = (NSView *)ptr;
  NSRect frame = control.frame;
  *x = frame.origin.x;
  *y = frame.origin.y;
  *w = frame.size.width;
  *h = frame.size.height;
}
