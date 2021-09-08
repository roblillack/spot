#import "imageview.h"
#include "_cgo_export.h"

ButtonPtr ImageView_New(int goButtonID, int x, int y, int w, int h, const char* url) {
    NSImage *theImage = [[NSImage alloc] initWithContentsOfURL:[NSURL URLWithString:[NSString stringWithUTF8String:url]]];

    id nsImageView = [[[NSImageView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
    [nsImageView setImage:theImage];

    return (ImageViewPtr)nsImageView;
}

void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageFrameStyle:frameStyle];
}

void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr, int imageAlignment) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageAlignment:imageAlignment];
}

void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling) {
    NSImageView* nsImageView = (NSImageView*)imageViewPtr;
    [nsImageView setImageScaling:imageScaling];
}
