#import <Cocoa/Cocoa.h>

typedef void *ViewPtr;

void View_SetFrameOrigin(ViewPtr ptr, int x, int y);
void View_SetFrameSize(ViewPtr ptr, int w, int h);
void View_SetFrame(ViewPtr ptr, int x, int y, int w, int h);
void View_Frame(ViewPtr ptr, int *x, int *y, int *w, int *h);