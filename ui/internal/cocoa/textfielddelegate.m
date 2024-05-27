#import "textfielddelegate.h"
#include "_cgo_export.h"

@implementation TextFieldDelegate

- (void)dealloc {
  [super dealloc];
}

- (void)controlTextDidChange:(NSNotification *)aNotification {
  //   NSTextField *textField = [aNotification object];
  //   const char *text = [[textField stringValue]
  //   cStringUsingEncoding:NSUTF8StringEncoding];
  onTextFieldDidChange([self goTextFieldId]);
}

@end
