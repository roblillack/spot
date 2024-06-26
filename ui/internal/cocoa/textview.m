#import "textview.h"
#include "_cgo_export.h"
#import "textviewdelegate.h"

@implementation TextViewHandler
@end

TextViewPtr TextView_New(int goTextViewId, int x, int y, int w, int h) {
  /* create the NSTextView and add it to the window */
  NSTextView *textView =
      [[[NSTextView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];

  id d = [[TextViewDelegate alloc] init];
  [d setGoTextViewId:goTextViewId];
  [textView setDelegate:d];

  NSScrollView *scrollView =
      [[[NSScrollView alloc] initWithFrame:NSMakeRect(x, y, w, h)] autorelease];
  [scrollView setDocumentView:textView];
  [scrollView setHasVerticalScroller:YES];
  scrollView.translatesAutoresizingMaskIntoConstraints = NO;
  //   tableView.translatesAutoresizingMaskIntoConstraints = NO;
  scrollView.borderType = NSBezelBorder;

  return (TextViewPtr)scrollView;
}

const char *TextView_Text(TextViewPtr ptr) {
  NSTextView *c = ((NSScrollView *)ptr).documentView;
  return
      [[[c textStorage] string] cStringUsingEncoding:NSISOLatin1StringEncoding];
}

void TextView_SetText(TextViewPtr ptr, const char *text) {
  NSTextView *c = ((NSScrollView *)ptr).documentView;
  [c setString:[NSString stringWithUTF8String:text]];
}

void TextView_Remove(TextViewPtr ptr) {
  NSScrollView *c = (NSScrollView *)ptr;
  [c removeFromSuperview];
}

void TextView_SetFontSize(TextViewPtr ptr, int size) {
  NSTextView *c = ((NSScrollView *)ptr).documentView;
  NSFont *font = [NSFont fontWithName:@"Helvetica" size:size];
  [c setFont:font];
}

const int TextView_Editable(TextViewPtr ptr) {
  NSTextView *c = ((NSScrollView *)ptr).documentView;
  return c.editable;
}

void TextView_SetEditable(TextViewPtr ptr, int editable) {
  NSTextView *c = ((NSScrollView *)ptr).documentView;
  [c becomeFirstResponder];
  c.editable = editable ? YES : NO;
  [c resignFirstResponder];
}