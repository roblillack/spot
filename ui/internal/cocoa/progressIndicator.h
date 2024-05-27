#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

typedef void* ProgressIndicatorPtr;

ProgressIndicatorPtr ProgressIndicator_New(int x, int y, int width, int height);

void ProgressIndicator_StartAnimation(ProgressIndicatorPtr progressIndicatorPtr);

void ProgressIndicator_StopAnimation(ProgressIndicatorPtr progressIndicatorPtr);

void ProgressIndicator_SetLimits(ProgressIndicatorPtr progressIndicatorPtr, double minValue, double maxValue);

void ProgressIndicator_SetValue(ProgressIndicatorPtr progressIndicatorPtr, double value);

void ProgressIndicator_IncrementBy(ProgressIndicatorPtr progressIndicatorPtr, double value);

void ProgressIndicator_SetIsIndeterminate(ProgressIndicatorPtr progressIndicatorPtr, int value);

int ProgressIndicator_IsIndeterminate(ProgressIndicatorPtr progressIndicatorPtr);

void ProgressIndicator_SetDisplayedWhenStopped(ProgressIndicatorPtr progressIndicatorPtr, int value);

void ProgressIndicator_Show(ProgressIndicatorPtr progressIndicatorPtr);

void ProgressIndicator_Hide(ProgressIndicatorPtr progressIndicatorPtr);

void ProgressIndicator_Remove(ProgressIndicatorPtr progressIndicatorPtr);
