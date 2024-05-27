#import "imageview.h"
#include "_cgo_export.h"
#import "image.h"

ImageViewPtr ImageView_New(int goButtonID, int x, int y, int w, int h) {
  id nsImageView =
      [[[NSImageView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

  return (ImageViewPtr)nsImageView;
}

ImageViewPtr ImageView_NewWithContentsOfURL(int goButtonID, int x, int y, int w,
                                            int h, const char *url) {
  NSImage *theImage = [[NSImage alloc]
      initWithContentsOfURL:[NSURL
                                URLWithString:[NSString
                                                  stringWithUTF8String:url]]];

  id nsImageView =
      [[[NSImageView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
  [nsImageView setImage:theImage];

  return (ImageViewPtr)nsImageView;
}

void ImageView_SetFrameStyle(ImageViewPtr imageViewPtr, int frameStyle) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setImageFrameStyle:frameStyle];
}

void ImageView_SetImageAlignment(ImageViewPtr imageViewPtr,
                                 int imageAlignment) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setImageAlignment:imageAlignment];
}

void ImageView_SetImageScaling(ImageViewPtr imageViewPtr, int imageScaling) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setImageScaling:imageScaling];
}

void ImageView_SetAnimates(ImageViewPtr imageViewPtr, int animates) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setAnimates:animates];
}

void ImageView_SetContentTintColor(ImageViewPtr imageViewPtr, int r, int g,
                                   int b, int a) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setContentTintColor:[NSColor colorWithCalibratedRed:r / 255.f
                                                             green:g / 255.f
                                                              blue:b / 255.f
                                                             alpha:a / 255.f]];
}

void ImageView_SetEditable(ImageViewPtr imageViewPtr, int editable) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setEditable:editable];
}

void ImageView_Remove(ImageViewPtr imageViewPtr) {
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView removeFromSuperview];
}

void ImageView_SetImage(ImageViewPtr imageViewPtr, ImagePtr imagePtr) {
  NSImage *theImage = (NSImage *)imagePtr;
  NSImageView *nsImageView = (NSImageView *)imageViewPtr;
  [nsImageView setImage:theImage];
}
