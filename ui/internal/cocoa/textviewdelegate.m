#import "textviewdelegate.h"
#include "_cgo_export.h"

@implementation TextViewDelegate

- (void)dealloc {
  [super dealloc];
}

- (void)textDidChange:(NSNotification *)aNotification {
  //   NSTextField *textField = [aNotification object];
  //   const char *text = [[textField stringValue]
  //   cStringUsingEncoding:NSUTF8StringEncoding];
  onTextViewDidChange([self goTextViewId]);
}

@end
