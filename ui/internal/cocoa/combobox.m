#import "combobox.h"
#include "_cgo_export.h"

@implementation ComboBoxHandler
- (void)comboBoxSelectionDidChange:(id)sender {
  onSelectionDidChange([self goComboBoxID]);
}
@end

ComboBoxPtr ComboBox_New(int goComboBoxID, int x, int y, int w, int h) {
  id nsComboBox =
      [[[NSComboBox alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
  ComboBoxHandler *handler = [[ComboBoxHandler alloc] init];
  [handler setGoComboBoxID:goComboBoxID];
  [nsComboBox setDelegate:handler];
  [nsComboBox setTarget:handler];
  [nsComboBox setAction:@selector(comboBoxSelectionDidChange:)];
  return (ComboBoxPtr)nsComboBox;
}

void ComboBox_AddItem(ComboBoxPtr comboBoxPtr, const char *item) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox addItemWithObjectValue:[NSString stringWithUTF8String:item]];
}

void ComboBox_SetEditable(ComboBoxPtr comboBoxPtr, int editable) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox setEditable:editable];
}

int ComboBox_SelectedIndex(ComboBoxPtr comboBoxPtr) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  return (int)[comboBox indexOfSelectedItem];
}

const char *ComboBox_SelectedText(ComboBoxPtr comboBoxPtr) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  return [[comboBox itemObjectValueAtIndex:[comboBox indexOfSelectedItem]]
      cStringUsingEncoding:NSISOLatin1StringEncoding];
}

void ComboBox_SetSelectedIndex(ComboBoxPtr comboBoxPtr, int selectedIndex) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox selectItemAtIndex:selectedIndex];
}

void ComboBox_SetSelectedText(ComboBoxPtr comboBoxPtr,
                              const char *selectedText) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox
      selectItemWithObjectValue:[NSString stringWithUTF8String:selectedText]];
}

void ComboBox_SetStringValue(ComboBoxPtr comboBoxPtr, const char *stringValue) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox setStringValue:[NSString stringWithUTF8String:stringValue]];
}

const char *ComboBox_StringValue(ComboBoxPtr comboBoxPtr) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  return
      [[comboBox stringValue] cStringUsingEncoding:NSISOLatin1StringEncoding];
}

void ComboBox_Remove(ComboBoxPtr comboBoxPtr) {
  NSComboBox *comboBox = (NSComboBox *)comboBoxPtr;
  [comboBox removeFromSuperview];
  [[comboBox target] release];
}