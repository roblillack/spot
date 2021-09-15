#import <Cocoa/Cocoa.h>

// typedef void (*callback)(void);

@interface ComboBoxHandler : NSObject<NSComboBoxDelegate>

@property (assign) int goComboBoxID;
-(void) comboBoxSelectionDidChange:(id) sender;

@end

typedef void* ComboBoxPtr;

ComboBoxPtr ComboBox_New(int goComboBoxID, int x, int y, int w, int h);
void ComboBox_AddItem(ComboBoxPtr comboBoxPtr, const char* item);
void ComboBox_SetEditable(ComboBoxPtr comboBoxPtr, int editable);
void ComboBox_SetSelectedIndex(ComboBoxPtr comboBoxPtr, int selectedIndex);
void ComboBox_SetSelectedText(ComboBoxPtr comboBoxPtr, const char* selectedText);
void ComboBox_SetStringValue(ComboBoxPtr comboBoxPtr, const char* stringValue);
int ComboBox_SelectedIndex(ComboBoxPtr comboBoxPtr);
const char* ComboBox_SelectedText(ComboBoxPtr comboBoxPtr);
const char* ComboBox_StringValue(ComboBoxPtr comboBoxPtr);
