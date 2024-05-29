#import <Cocoa/Cocoa.h>

@interface TableViewController
    : NSViewController <NSTableViewDataSource, NSTableViewDelegate>
@property(assign) int goTableViewId;
@property(strong) NSArray *dataList;
@property(strong) NSTableView *tableView;
@end

typedef void *TableViewPtr;

TableViewPtr TableView_New(int goTableViewId, int x, int y, int w, int h);
void TableView_Remove(TableViewPtr tableViewPtr);

const int TableView_IsEnabled(TableViewPtr tableViewPtr);
void TableView_SetEnabled(TableViewPtr tableViewPtr, int enabled);

const int TableView_AllowsMultipleSelection(TableViewPtr tableViewPtr);
void TableView_SetAllowsMultipleSelection(TableViewPtr tableViewPtr,
                                          int enabled);

void TableView_Clear(TableViewPtr tableViewPtr);
void TableView_Add(TableViewPtr tableViewPtr, const char *text);

void TableView_DeselectAll(TableViewPtr tableViewPtr);
void TableView_SelectRowIndex(TableViewPtr tableViewPtr, int row);
int TableView_NumberOfRows(TableViewPtr tableViewPtr);
const int TableView_IsRowSelected(TableViewPtr tableViewPtr, int row);

void TableView_ScrollRowToVisible(TableViewPtr tableViewPtr, int row);
