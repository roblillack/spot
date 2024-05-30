#import "image.h"
#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface ButtonHandler : NSObject

@property(assign) int goButtonID;
- (void)buttonClicked:(id)sender;

@end

typedef void *ButtonPtr;

ButtonPtr Button_New(int goButtonID, int x, int y, int w, int h);
void Button_SetTitle(ButtonPtr btnPtr, const char *title);
void Button_Remove(ButtonPtr btnPtr);
const char *Button_Title(ButtonPtr btnPtr);
void Button_SetButtonType(ButtonPtr btnPtr, int buttonType);
void Button_SetBezelStyle(ButtonPtr btnPtr, int bezelStyle);
void Button_SetFontFamily(ButtonPtr btnPtr, const char *fontFamily);
void Button_SetFontSize(ButtonPtr btnPtr, int fontSize);
void Button_SetColor(ButtonPtr btnPtr, int r, int g, int b, int a);
void Button_SetBackgroundColor(ButtonPtr btnPtr, int r, int g, int b, int a);
void Button_SetBorderColor(ButtonPtr btnPtr, int r, int g, int b, int a);
void Button_SetBorderWidth(ButtonPtr btnPtr, int borderWidth);
void Button_SetState(ButtonPtr btnPtr, int state);
int Button_State(ButtonPtr btnPtr);
void Button_SetImage(ButtonPtr btnPtr, ImagePtr imagePtr);