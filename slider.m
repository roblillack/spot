#import "slider.h"
#include "_cgo_export.h"

@implementation SliderHandler
-(void) sliderValueChanged:(id) sender
{
    onSliderValueChanged([self goSliderID]);
}
@end

SliderPtr Slider_New(int goSliderID, int x, int y, int w, int h) {
	id nsSlider = [[[NSSlider alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

    SliderHandler* handler = [[SliderHandler alloc] init];
    [handler setGoSliderID:goSliderID];
    [nsSlider setTarget:handler];
    [nsSlider setAction:@selector(sliderValueChanged:)];
    // [handler autorelease];

	return (SliderPtr)nsSlider;
}

void Slider_SetMaximumValue(SliderPtr sliderPtr, double val) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    [slider setMaxValue:val];
}

void Slider_SetMinimumValue(SliderPtr sliderPtr, double val) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    [slider setMinValue:val];
}

void Slider_SetValue(SliderPtr sliderPtr, double val) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    [slider setDoubleValue:val];
}

double Slider_Value(SliderPtr sliderPtr) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    return slider.doubleValue;
}

void Slider_SetSliderType(SliderPtr sliderPtr, int sliderType) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    [slider setSliderType:sliderType];
}

void Slider_Remove(SliderPtr sliderPtr) {
    NSSlider* slider = (NSSlider*)sliderPtr;
    [slider removeFromSuperview];
}