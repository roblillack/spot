#import "windowdelegate.h"
#include "_cgo_export.h"

const int DID_RESIZE_EVENT = 0;
const int DID_MOVE_EVENT = 1;
const int DID_MINIATURIZE_EVENT = 2;
const int DID_DEMINIATURIZE_EVENT = 3;
const int SHOULD_CLOSE_EVENT = 4;

@implementation WindowDelegate

- (void)dealloc {
  [super dealloc];
}

- (void)windowDidResize:(NSNotification *)notification {
  NSWindow *movedWindow = notification.object;
  triggerEvent([self goWindowID], movedWindow, @"windowDidResize",
               DID_RESIZE_EVENT);
}

- (void)windowDidMove:(NSNotification *)notification {
  NSWindow *movedWindow = notification.object;
  triggerEvent([self goWindowID], movedWindow, @"windowDidMove",
               DID_MOVE_EVENT);
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
  NSWindow *movedWindow = notification.object;
  triggerEvent([self goWindowID], movedWindow, @"windowDidMiniaturize",
               DID_MINIATURIZE_EVENT);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
  NSWindow *movedWindow = notification.object;
  triggerEvent([self goWindowID], movedWindow, @"windowDidDeminiaturize",
               DID_DEMINIATURIZE_EVENT);
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
  // TODO: In the future, we should send the event to Go and let Go decide
  //       and allow the Go side to optionally trigger some alternative
  //       behavior for secondary windows.
  // triggerEvent([self goWindowID], sender, @"windowShouldClose",
  //              SHOULD_CLOSE_EVENT);
  [NSApp stop:sender];
  return YES;
}

@end

void triggerEvent(int goWindowID, NSWindow *movedWindow, NSString *eventTitle,
                  const int eventId) {
  if ([movedWindow isKeyWindow]) {
    NSRect rect = movedWindow.frame;
    int x = (int)(rect.origin.x);
    int y = (int)(rect.origin.y);
    int w = (int)(rect.size.width);
    int h = (int)(rect.size.height);
    // NSLog(@"%@ %@", eventTitle, movedWindow);
    onWindowEvent(goWindowID, eventId, x, y, w, h);
  }
}