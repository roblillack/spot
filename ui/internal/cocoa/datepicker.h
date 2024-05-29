#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface DatePickerHandler : NSObject

@property (assign) int goDatePickerID;

@end

typedef void* DatePickerPtr;

DatePickerPtr DatePicker_New(int goDatePickerID, int x, int y, int w, int h);
void DatePicker_SetStyle(DatePickerPtr datePickerPtr, int style);
void DatePicker_SetMode(DatePickerPtr datePickerPtr, int mode);
void DatePicker_SetDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat);
void DatePicker_SetMinimumDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat);
void DatePicker_SetMaximumDate(DatePickerPtr datePickerPtr, const char* date, const char* dateFormat);
const char* DatePicker_Date(DatePickerPtr datePickerPtr, const char* dateFormat);
void DatePicker_Remove(DatePickerPtr datePickerPtr);
