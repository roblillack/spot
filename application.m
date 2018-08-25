#import "application.h"
#include "_cgo_export.h"

// InitSharedApplication calls NSApplication.sharedApplication() which initializes the 
// global application instance NSApp.
void InitSharedApplication() {
    [NSApplication sharedApplication];
}

void RunApplication() {
    @autoreleasepool {
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [NSApp activateIgnoringOtherApps:YES];
        [NSApp run];
    }
}
