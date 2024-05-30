#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface SliderHandler : NSObject

@property(assign) int goSliderID;
- (void)sliderValueChanged:(id)sender;

@end

typedef void *SliderPtr;

SliderPtr Slider_New(int goSliderID, int x, int y, int w, int h);
void Slider_SetMaximumValue(SliderPtr sliderPtr, double val);
void Slider_SetMinimumValue(SliderPtr sliderPtr, double val);
void Slider_SetValue(SliderPtr sliderPtr, double val);
void Slider_SetSliderType(SliderPtr sliderPtr, int sliderType);
double Slider_Value(SliderPtr sliderPtr);
void Slider_Remove(SliderPtr sliderPtr);