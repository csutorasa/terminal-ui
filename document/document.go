package document

import (
	"github.com/csutorasa/terminal-ui/ansi"
)

var Default *Document = NewDocument()

// Dom document
type Document struct {
	dom        *Dom
	dispatcher *Dispatcher
}

// Creates a new document.
func NewDocument() *Document {
	document := &Document{
		dispatcher: NewDispatcher(),
	}
	document.dom = NewDom(document)
	return document
}

// Dispatches a new event.
func (d *Document) Dispatch(callback DispatcherCallback) {
	d.dispatcher.Dispatch(callback)
}

// Triggers a new event.
func (d *Document) Trigger(e any) {
	switch ev := e.(type) {
	case *KeyboardEvent:
		if ev.Key == ansi.KeyTab {
			d.dom.FocusNext()
			return
		}
	case *DispatchEvent:
		d.dispatcher.Run()
		return
	}
	event := NewEvent(e)
	d.triggerEvent(event)
	d.dispatcher.Run()
}

// Renders the document, and returns the output.
// It also returns if the out has changed since the last call.
func (d *Document) Render(rc RenderContext) (RenderOutput, bool) {
	d.dom.mutex.Lock()
	defer d.dom.mutex.Unlock()
	root := d.dom.validateHasRoot()
	layoutContext := NewLayoutContext(rc)
	layoutContext.add(root, rc)
	root.renderLayout(layoutContext)
	layoutContext.apply()
	changed := root.render()
	return root.Render(), changed
}

// Sets the document root.
func (d *Document) SetRoot(root *Element) {
	d.dom.SetRoot(root)
}

// Sets the focus to an element.
func (d *Document) SetFocus(element *Element) {
	d.dom.SetFocus(element)
}

// Cycles the focus to the next [Element].
func (d *Document) FocusNext() {
	d.dom.FocusNext()
}

func (d *Document) triggerEvent(event *Event) {
	d.dom.mutex.Lock()
	defer d.dom.mutex.Unlock()
	root := d.dom.validateHasRoot()
	focus := d.dom.focus
	if focus == nil {
		d.propagateEvent(root, event)
	} else {
		d.propagateEvent(focus, event)
	}
}

func (d *Document) propagateEvent(element *Element, event *Event) {
	parents := d.dom.findParents(element)
	for _, e := range parents {
		e.onEvent(event)
		if !event.propagating() {
			return
		}
	}
}
