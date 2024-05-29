#import "textfield.h"
#include "_cgo_export.h"

@implementation DatePickerHandler
@end

DatePickerPtr DatePicker_New(int goDatePickerId, int x, int y, int w, int h) {
	id nsDatePicker = [[[NSDatePicker alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

	return (DatePickerPtr)nsDatePicker;
}

void DatePicker_SetStyle(DatePickerPtr datePickerPtr, int style) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    [datePicker setDatePickerStyle:style];
}

void DatePicker_SetMode(DatePickerPtr datePickerPtr, int mode) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    [datePicker setDatePickerMode:mode];
}

void DatePicker_SetDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    NSDateFormatter *dateFormatter = [[NSDateFormatter alloc] init];
    dateFormatter.dateFormat = [NSString stringWithUTF8String:dateFormat];
    NSDate *formattedDate = [dateFormatter dateFromString: [NSString stringWithUTF8String:date]];

    [datePicker setDateValue:formattedDate];
}

void DatePicker_SetMinimumDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    NSDateFormatter *dateFormatter = [[NSDateFormatter alloc] init];
    dateFormatter.dateFormat = [NSString stringWithUTF8String:dateFormat];
    NSDate *formattedDate = [dateFormatter dateFromString: [NSString stringWithUTF8String:date]];

    [datePicker setMinDate:formattedDate];
}

void DatePicker_SetMaximumDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    NSDateFormatter *dateFormatter = [[NSDateFormatter alloc] init];
    dateFormatter.dateFormat = [NSString stringWithUTF8String:dateFormat];
    NSDate *formattedDate = [dateFormatter dateFromString: [NSString stringWithUTF8String:date]];

    [datePicker setMaxDate:formattedDate];
}

const char* DatePicker_Date(DatePickerPtr datePickerPtr, const char* dateFormat) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    NSDateFormatter *dateFormatter = [[NSDateFormatter alloc] init];
    dateFormatter.dateFormat = [NSString stringWithUTF8String:dateFormat];

    NSDate *selectedDate = datePicker.dateValue;
    NSString *stringFromDate = [dateFormatter stringFromDate:selectedDate];
    return [stringFromDate cStringUsingEncoding:NSASCIIStringEncoding];
}

void DatePicker_Remove(DatePickerPtr datePickerPtr) {
    NSDatePicker* datePicker = (NSDatePicker*)datePickerPtr;
    [datePicker removeFromSuperview];
}