#import <Cocoa/Cocoa.h>
#include <stdint.h>

void InitSharedApplication();
void RunApplication();
void TerminateApplication();
void RunOnMainLoop(uintptr_t h);