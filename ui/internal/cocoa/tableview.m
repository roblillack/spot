#import "tableview.h"
#include "_cgo_export.h"

@implementation TableViewController
- (NSInteger)numberOfRowsInTableView:(NSTableView *)tableView {
  return self.dataList.count;
}

- (void)tableViewSelectionDidChange:(NSNotification *)notification {
  onTableViewSelectionDidChange(self.goTableViewId);
}

- (id)tableView:(NSTableView *)tableView
    objectValueForTableColumn:(NSTableColumn *)tableColumn
                          row:(NSInteger)row {
  if (self.dataList == nil || row < 0 || row >= self.dataList.count) {
    return nil;
  }

  return self.dataList[row];
}

- (BOOL)tableView:(NSTableView *)tableView
    shouldEditTableColumn:(NSTableColumn *)tableColumn
                      row:(NSInteger)row {
  return NO;
}

- (double)tableView:(NSTableView *)tableView heightOfRow:(long)row {
  return 18;
}

@end

TableViewPtr TableView_New(int goTableViewId, int x, int y, int w, int h) {
  NSTableView *tableView =
      [[[NSTableView alloc] initWithFrame:NSMakeRect(0, 0, w, h)] autorelease];

  NSScrollView *scrollView =
      [[[NSScrollView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
  [scrollView setDocumentView:tableView];
  [scrollView setHasVerticalScroller:YES];
  scrollView.translatesAutoresizingMaskIntoConstraints = NO;
  //   tableView.translatesAutoresizingMaskIntoConstraints = NO;
  scrollView.borderType = NSBezelBorder;

  NSTableColumn *column =
      [[[NSTableColumn alloc] initWithIdentifier:@"Column"] autorelease];
  column.width = w;
  [tableView addTableColumn:column];

  TableViewController *d = [[[TableViewController alloc] init] autorelease];
  d.goTableViewId = goTableViewId;

  // This line removes the header view entirely
  tableView.headerView = nil;
  tableView.delegate = d;
  tableView.dataSource = d;

  // things that don't seem to do a thing:
  tableView.rowSizeStyle = NSTableViewRowSizeStyleSmall;
  tableView.intercellSpacing = NSZeroSize;

  return (TableViewPtr)scrollView;
}

void TableView_Remove(TableViewPtr tableViewPtr) {
  NSScrollView *x = (NSScrollView *)tableViewPtr;
  [x removeFromSuperview];
  [[x delegate] release];
}

const int TableView_IsEnabled(TableViewPtr tableViewPtr) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  return tableView.enabled ? 1 : 0;
}

void TableView_SetEnabled(TableViewPtr tableViewPtr, int enabled) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  [tableView becomeFirstResponder];
  tableView.enabled = enabled ? YES : NO;
  [tableView resignFirstResponder];
}

const int TableView_AllowsMultipleSelection(TableViewPtr tableViewPtr) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  return tableView.allowsMultipleSelection;
}

void TableView_SetAllowsMultipleSelection(TableViewPtr tableViewPtr,
                                          int allowed) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  [tableView becomeFirstResponder];
  tableView.allowsMultipleSelection = allowed ? YES : NO;
  [tableView resignFirstResponder];
}

void TableView_Clear(TableViewPtr tableViewPtr) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  ((TableViewController *)tableView.dataSource).dataList = @[];
  [tableView reloadData];
}

void TableView_Add(TableViewPtr tableViewPtr, const char *text) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  ((TableViewController *)tableView.dataSource).dataList =
      [((TableViewController *)tableView.dataSource).dataList
          arrayByAddingObject:[NSString stringWithUTF8String:text]];
  [tableView reloadData];
}

void TableView_DeselectAll(TableViewPtr tableViewPtr) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  [tableView deselectAll:nil];
}

void TableView_SelectRowIndex(TableViewPtr tableViewPtr, int row) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  [tableView selectRowIndexes:[NSIndexSet indexSetWithIndex:row]
         byExtendingSelection:NO];
}

int TableView_NumberOfRows(TableViewPtr tableViewPtr) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  return (int)[tableView numberOfRows];
}

const int TableView_IsRowSelected(TableViewPtr tableViewPtr, int row) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  return [tableView isRowSelected:row] ? 1 : 0;
}

void TableView_ScrollRowToVisible(TableViewPtr tableViewPtr, int row) {
  NSTableView *tableView = ((NSScrollView *)tableViewPtr).documentView;
  [tableView scrollRowToVisible:row];
}