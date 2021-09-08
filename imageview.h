#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

typedef void* ImageViewPtr;

ImageViewPtr ImageView_New(int goImageViewID, int x, int y, int w, int h, const char* url);
void ImageView_SetAnimates(ImageViewPtr imageViewPtr, int animates);
void ImageView_SetContentTintColor(ImageViewPtr imageViewPtr, int r, int g, int b, int a);
void ImageView_SetEditable(ImageViewPtr imageViewPtr, int editable);
void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle);
void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr, int imageAlignment);
void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling);
