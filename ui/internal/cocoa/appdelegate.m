#import "appdelegate.h"
#include "_cgo_export.h"

@implementation AppDelegate

- (void)dealloc {
  [super dealloc];
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
  [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
  [NSApp activateIgnoringOtherApps:YES];

  callOnApplicationDidFinishLaunchingHandler();
}

@end
