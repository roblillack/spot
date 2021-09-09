#import "application.h"
#import "appdelegate.h"
#include "_cgo_export.h"

AppDelegate *gocoa_applicationDelegate = nil;

// InitSharedApplication calls NSApplication.sharedApplication() which initializes the 
// global application instance NSApp.
void InitSharedApplication() {
    static bool hasBeenInitialized = false; // false first time function is called
    if (hasBeenInitialized)
        return;
    [NSApplication sharedApplication];
    gocoa_applicationDelegate = [[AppDelegate alloc] init];
    [NSApp setDelegate:gocoa_applicationDelegate];
    hasBeenInitialized = true;
}

void releaseSharedApplication() {
    if (gocoa_applicationDelegate != nil) {
        [gocoa_applicationDelegate release];
    }
}

void RunApplication() {
    @autoreleasepool {

        NSString *gocoaBundleIdentifier = @"gocoa";
        NSString *bundleIdentifier = [[NSBundle mainBundle] bundleIdentifier];
        bundleIdentifier = gocoaBundleIdentifier;

        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [NSApp activateIgnoringOtherApps:YES];
        [NSApp deactivate];
        [NSApp activateIgnoringOtherApps:YES];
        [NSApp run];

        releaseSharedApplication();
    }
}

void TerminateApplication() {
    [NSApp terminate:nil];
}