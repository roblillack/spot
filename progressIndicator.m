#import "progressIndicator.h"
#include "_cgo_export.h"

ProgressIndicatorPtr ProgressIndicator_New(int x, int y, int width, int height) {
    id nsProgressIndicator = [[NSProgressIndicator alloc] init];
    [nsProgressIndicator setFrame: NSMakeRect(x, y, width, height)];
    //[nsProgressIndicator setBounds: NSMakeRect(x, y, width, height)];
    [nsProgressIndicator setUsesThreadedAnimation:YES];
    [nsProgressIndicator autorelease];
    return (ProgressIndicatorPtr)nsProgressIndicator;
}

void ProgressIndicator_StartAnimation(ProgressIndicatorPtr progressIndicatorPtr)
{
    dispatch_async(dispatch_get_main_queue(), ^{
        NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
        [progressIndicator startAnimation:nil];
        [progressIndicator setNeedsDisplay:YES];
    });
}

void ProgressIndicator_StopAnimation(ProgressIndicatorPtr progressIndicatorPtr)
{
    dispatch_async(dispatch_get_main_queue(), ^{
        NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
         [progressIndicator stopAnimation:nil];
        [progressIndicator setNeedsDisplay:YES];
    });
}

void ProgressIndicator_SetLimits(ProgressIndicatorPtr progressIndicatorPtr, double minValue, double maxValue)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    [progressIndicator setMinValue:minValue];
    [progressIndicator setMaxValue:maxValue];
}

void ProgressIndicator_SetValue(ProgressIndicatorPtr progressIndicatorPtr, double value)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    dispatch_async(dispatch_get_main_queue(), ^{
        [progressIndicator setDoubleValue:value];
        // NSLog(@"progressIndicator value is %.2f", value);
        [progressIndicator setNeedsDisplay:YES];
    });
}

void ProgressIndicator_IncrementBy(ProgressIndicatorPtr progressIndicatorPtr, double value)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    dispatch_async(dispatch_get_main_queue(), ^{
        [progressIndicator incrementBy:value];
        NSLog(@"increasing progressIndicator value by %.2f", value);
        [progressIndicator setNeedsDisplay:YES];
    });
}

void ProgressIndicator_SetIsIndeterminate(ProgressIndicatorPtr progressIndicatorPtr, int value)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    if(value > 0) {
        [progressIndicator setIndeterminate: YES];
    } else {
        [progressIndicator setIndeterminate: NO];
    }
}

int ProgressIndicator_IsIndeterminate(ProgressIndicatorPtr progressIndicatorPtr)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    if ([progressIndicator isIndeterminate]) {
        return 1;
    }
    return 0;
}

void ProgressIndicator_SetDisplayedWhenStopped(ProgressIndicatorPtr progressIndicatorPtr, int value) {
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    if(value > 0) {
        [progressIndicator setDisplayedWhenStopped: YES];
    } else {
        [progressIndicator setDisplayedWhenStopped: NO];
    }
}

void ProgressIndicator_Show(ProgressIndicatorPtr progressIndicatorPtr)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    dispatch_async(dispatch_get_main_queue(), ^{
        progressIndicator.hidden = NO;
    });
}

void ProgressIndicator_Hide(ProgressIndicatorPtr progressIndicatorPtr){
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    [NSThread sleepForTimeInterval:1.0];
    dispatch_async(dispatch_get_main_queue(), ^{
        progressIndicator.hidden = YES;
    });
}

void ProgressIndicator_Remove(ProgressIndicatorPtr progressIndicatorPtr)
{
    NSProgressIndicator* progressIndicator = (NSProgressIndicator*)progressIndicatorPtr;
    [progressIndicator removeFromSuperview];
}
