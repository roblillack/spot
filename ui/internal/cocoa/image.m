#import "image.h"
#include "_cgo_export.h"

ImagePtr Image_NewWithRGBA(int w, int h, unsigned char *rgba) {
  NSBitmapImageRep *bitmap = [[NSBitmapImageRep alloc]
      initWithBitmapDataPlanes:NULL
                    pixelsWide:w
                    pixelsHigh:h
                 bitsPerSample:8
               samplesPerPixel:4
                      hasAlpha:YES
                      isPlanar:NO
                colorSpaceName:NSCalibratedRGBColorSpace
                   bytesPerRow:w * 4
                  bitsPerPixel:32];

  unsigned char *bitmapData = [bitmap bitmapData];
  memcpy(bitmapData, rgba, w * h * 4);

  NSImage *theImage = [[NSImage alloc] init];
  [theImage addRepresentation:bitmap];

  return (ImagePtr)theImage;
}
