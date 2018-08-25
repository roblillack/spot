#import "application.h"
#include "_cgo_export.h"

// GetSharedApplication() Returns the shared application object, creating it if necessary.
// TODO: What about lifetime/deconstruction/drain?
// TODO: Should be some kind of singleton perhaps?
void*
GetSharedApplication() {
    return [NSApplication sharedApplication];
}

void
App_Run(void* nsApp) {
    NSApplication* app = (NSApplication*)nsApp;

    @autoreleasepool {
        [app setActivationPolicy:NSApplicationActivationPolicyRegular];
        [app activateIgnoringOtherApps:YES];
        [app run];
    }
}
