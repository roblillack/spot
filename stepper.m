#import "stepper.h"
#include "_cgo_export.h"

@implementation StepperHandler
- (void)stepperValueChanged:(id)sender {
  onStepperValueChanged([self goStepperId]);
}
@end

StepperPtr Stepper_New(int goStepperId, int x, int y, int w, int h) {
  id nsStepper =
      [[[NSStepper alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

  StepperHandler *handler = [[StepperHandler alloc] init];
  [handler setGoStepperId:goStepperId];
  [handler autorelease];
  [nsStepper setTarget:handler];
  [nsStepper setAction:@selector(stepperValueChanged:)];

  return (StepperPtr)nsStepper;
}

void Stepper_SetMaxValue(StepperPtr StepperPtr, double val) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper setMaxValue:val];
}

void Stepper_SetMinValue(StepperPtr StepperPtr, double val) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper setMinValue:val];
}

void Stepper_SetIncrement(StepperPtr StepperPtr, double val) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper setIncrement:val];
}

void Stepper_SetValue(StepperPtr StepperPtr, double val) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper setDoubleValue:val];
}

bool Stepper_ValueWraps(StepperPtr StepperPtr) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  return stepper.valueWraps;
}

void Stepper_SetValueWraps(StepperPtr StepperPtr, bool wraps) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper setValueWraps:wraps];
}

double Stepper_Value(StepperPtr StepperPtr) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  return stepper.doubleValue;
}

void Stepper_Remove(StepperPtr StepperPtr) {
  NSStepper *stepper = (NSStepper *)StepperPtr;
  [stepper removeFromSuperview];
}