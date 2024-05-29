package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "tableview.h"
import "C"
import "unsafe"

// TableView -Button represents a button control that can trigger actions.
type TableView struct {
	tableViewPtr C.TableViewPtr
	callback     func(indexes []int)
}

var tableViews []*TableView

// NewTableView - This func is not thread safe.
func NewTableView(x int, y int, width int, height int) *TableView {
	tableViewID := len(tableViews)
	tableViewPtr := C.TableView_New(C.int(tableViewID), C.int(x), C.int(y), C.int(width), C.int(height))

	tf := &TableView{
		tableViewPtr: tableViewPtr,
	}
	tableViews = append(tableViews, tf)
	return tf
}

// Remove - removes a Text Field from the parent view
func (tableView *TableView) Remove() {
	C.TableView_Remove(tableView.tableViewPtr)
}

func (tableView *TableView) IsEnabled() bool {
	return C.TableView_IsEnabled(tableView.tableViewPtr) == 1
}

func (tableView *TableView) SetEnabled(enabled bool) {
	if enabled {
		C.TableView_SetEnabled(tableView.tableViewPtr, 1)
	} else {
		C.TableView_SetEnabled(tableView.tableViewPtr, 0)
	}
}

func (tableView *TableView) AllowsMultipleSelection() bool {
	return C.TableView_AllowsMultipleSelection(tableView.tableViewPtr) == 1
}

func (tableView *TableView) SetAllowsMultipleSelection(enabled bool) {
	if enabled {
		C.TableView_SetAllowsMultipleSelection(tableView.tableViewPtr, 1)
	} else {
		C.TableView_SetAllowsMultipleSelection(tableView.tableViewPtr, 0)
	}
}

func (tableView *TableView) Clear() {
	C.TableView_Clear(tableView.tableViewPtr)
}

func (tableView *TableView) Add(row string) {
	cRow := C.CString(row)
	defer C.free(unsafe.Pointer(cRow))
	C.TableView_Add(tableView.tableViewPtr, cRow)
}

func (tableView *TableView) DeselectAll() {
	C.TableView_DeselectAll(tableView.tableViewPtr)
}

func (tableView *TableView) SelectRowIndex(index int) {
	C.TableView_SelectRowIndex(tableView.tableViewPtr, C.int(index))
}

func (tableView *TableView) NumberOfRows() int {
	return int(C.TableView_NumberOfRows(tableView.tableViewPtr))
}

func (tableView *TableView) IsRowSelected(index int) bool {
	return C.TableView_IsRowSelected(tableView.tableViewPtr, C.int(index)) == 1
}

func (tableView *TableView) OnSelectionDidChange(fn func(indexes []int)) {
	tableView.callback = fn
}

func (tableView *TableView) ScrollRowToVisible(index int) {
	C.TableView_ScrollRowToVisible(tableView.tableViewPtr, C.int(index))
}

func (tableView *TableView) selectedRows() []int {
	var indexes []int
	for i := 0; i < int(C.TableView_NumberOfRows(tableView.tableViewPtr)); i++ {
		if (C.TableView_IsRowSelected(tableView.tableViewPtr, C.int(i))) == 1 {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

//export onTableViewSelectionDidChange
func onTableViewSelectionDidChange(id C.int) {
	tableViewID := int(id)
	if tableViewID < len(tableViews) && tableViews[tableViewID].callback != nil {
		x := tableViews[tableViewID]
		if x.callback != nil {
			x.callback(x.selectedRows())
		}
	}
}
