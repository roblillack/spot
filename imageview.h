#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface ImageViewHandler : NSObject

@property (assign) int goImageViewID;

@end

typedef void* ImageViewPtr;

ImageViewPtr ImageView_New(int goImageViewID, int x, int y, int w, int h, const char* url);
void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle);
void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr, int imageAlignment);
void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling);
