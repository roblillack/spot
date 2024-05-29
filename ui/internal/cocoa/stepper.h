#import <Cocoa/Cocoa.h>

@interface StepperHandler : NSObject

@property(assign) int goStepperId;
- (void)stepperValueChanged:(id)sender;

@end

typedef void *StepperPtr;

StepperPtr Stepper_New(int goStepperId, int x, int y, int w, int h);
void Stepper_SetMaxValue(StepperPtr StepperPtr, double val);
void Stepper_SetMinValue(StepperPtr StepperPtr, double val);
void Stepper_SetIncrement(StepperPtr StepperPtr, double val);
void Stepper_SetValue(StepperPtr StepperPtr, double val);
double Stepper_Value(StepperPtr StepperPtr);
void Stepper_Remove(StepperPtr StepperPtr);
bool Stepper_ValueWraps(StepperPtr StepperPtr);
void Stepper_SetValueWraps(StepperPtr StepperPtr, bool wraps);