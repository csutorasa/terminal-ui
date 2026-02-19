package document

import (
	"sync"

	"github.com/csutorasa/terminal-ui/ansi"
)

var Default *Document = NewDocument()

// Dom document
type Document struct {
	root     *Element
	focus    *Element
	elements map[Component]*Element
	parents  map[*Element]*Element
	mutex    sync.RWMutex
}

// Creates a new document.
func NewDocument() *Document {
	return &Document{
		elements: map[Component]*Element{},
		parents:  map[*Element]*Element{},
	}
}

// Triggers a new event.
func (d *Document) Trigger(e any) {
	if d.root == nil {
		panic(ErrDocumentRootMissing)
	}
	keyEvent, ok := e.(*KeyboardEvent)
	if ok {
		if keyEvent.Key == ansi.KeyTab {
			d.focusNext()
			return
		}
	}
	event := NewEvent(e)
	if d.focus == nil {
		d.root.onEvent(event)
	} else {
		parents := d.findParents(d.focus)
		if parents == nil {
			d.SetFocus(nil)
			d.root.onEvent(event)
		} else {
			for _, e := range parents {
				e.onEvent(event)
				if !event.propagate {
					break
				}
				e = d.parents[e]
			}
		}
	}
}

// Renders the document, and returns the output.
// It also returns if the out has changed since the last call.
func (d *Document) Render(rc RenderContext) (RenderOutput, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if d.root == nil {
		panic(ErrDocumentRootMissing)
	}
	layoutContext := NewLayoutContext(rc)
	layoutContext.Add(d.root, rc)
	d.root.renderLayout(layoutContext)
	layoutContext.apply()
	changed := d.root.render()
	return d.root.Render(), changed
}

// Sets the document root.
func (d *Document) SetRoot(root *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.root = root
	d.focus = nil
}

// Sets the focus to an element.
func (d *Document) SetFocus(element *Element) {
	if element.document != d {
		panic(ErrDocumentMismatch)
	}
	if d.focus == element {
		return
	}
	if d.focus != nil {
		d.focus.properties.focus.Set(false)
		focusLostEvent := NewEvent(&OnFocusLostEvent{})
		d.Trigger(focusLostEvent)
	}
	d.setFocus(element)
	if d.focus != nil {
		d.focus.properties.focus.Set(true)
		focusEvent := NewEvent(&OnFocusEvent{})
		d.Trigger(focusEvent)
	}
}

func (d *Document) setFocus(element *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if element.document != d {
		panic(ErrDocumentMismatch)
	}
	if d.root == nil {
		panic(ErrDocumentRootMissing)
	}
	if !d.isInTree(element) {
		panic(ErrCorruptDom)
	}
	if !element.component.Focusable() {
		return
	}
	d.focus = element
}

func (d *Document) focusNext() {
	if d.root == nil {
		panic(ErrDocumentRootMissing)
	}
	if d.focus == nil {
		foundStart := true
		recurse([]*Element{d.root}, &foundStart, nil, nil, func(e *Element) bool {
			return e.component.Focusable()
		}, func(e *Element) {
			d.SetFocus(e)
		})
		return
	}
	curr := d.focus
	foundStart := false
	found := recurse([]*Element{d.root}, &foundStart, curr, nil, func(e *Element) bool {
		return e.component.Focusable()
	}, func(e *Element) {
		d.SetFocus(e)
	})
	if found {
		return
	}
	foundStart = true
	recurse([]*Element{d.root}, &foundStart, nil, curr, func(e *Element) bool {
		return e.component.Focusable()
	}, func(e *Element) {
		d.SetFocus(e)
	})
}

func recurse(elements []*Element, foundStart *bool, from *Element, to *Element, condition func(*Element) bool, process func(*Element)) bool {
	for _, e := range elements {
		if !*foundStart && e == from {
			*foundStart = true
			continue
		}
		if to != nil && to == e {
			return false
		}
		if *foundStart && condition(e) {
			process(e)
			return true
		}
		c := recurse(e.properties.children.Value(), foundStart, from, to, condition, process)
		if c {
			return true
		}
	}
	return false
}

func (d *Document) findParents(e *Element) []*Element {
	// TODO
	/*d.mutex.RLock()
	defer d.mutex.RUnlock()*/
	if e.document != d {
		panic(ErrDocumentMismatch)
	}
	if !d.isInTree(e) {
		return nil
	}
	parents := []*Element{}
	for e != nil {
		parents = append(parents, e)
		e = d.parents[e]
	}
	return parents
}

func (d *Document) addChild(parent *Element, child *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if child.document != d {
		panic(ErrDocumentMismatch)
	}
	_, found := d.parents[child]
	if found {
		panic(ErrCorruptDom)
	}
	d.parents[child] = parent
}

func (d *Document) removeChild(parent *Element, child *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	p, found := d.parents[child]
	if !found {
		panic(ErrCorruptDom)
	}
	if p != parent {
		panic(ErrCorruptDom)
	}
	delete(d.parents, child)
}

func (d *Document) isInTree(element *Element) bool {
	if element.document != d {
		panic(ErrDocumentMismatch)
	}
	for element != nil {
		if element == d.root {
			return true
		}
		element = d.parents[element]
	}
	return false
}
