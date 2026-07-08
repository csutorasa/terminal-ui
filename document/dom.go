package document

import (
	"sync"
)

// Dom document
type Dom struct {
	document *Document
	root     *Element
	focus    *Element
	elements map[Component]*Element
	parents  map[*Element]*Element
	mutex    sync.RWMutex
}

// Creates a new [Dom].
func NewDom(document *Document) *Dom {
	return &Dom{
		document: document,
		elements: map[Component]*Element{},
		parents:  map[*Element]*Element{},
	}
}

// Sets the [Dom] root.
func (d *Dom) SetRoot(root *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.root = root
	d.focus = nil
}

// Gets the [Dom] root.
func (d *Dom) Root() *Element {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.validateHasRoot()
}

// Sets the focus to an element.
func (d *Dom) SetFocus(element *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.validateHasRoot()
	if element != nil {
		d.validateDocument(element)
		d.validateInTree(element)
	}
	d.setFocus(element)
}

// Cycles the focus to the next [Element].
func (d *Dom) FocusNext() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	root := d.validateHasRoot()
	if d.focus == nil {
		foundStart := true
		d.recurse([]*Element{root}, &foundStart, nil, nil)
		return
	}
	curr := d.focus
	foundStart := false
	found := d.recurse([]*Element{root}, &foundStart, curr, nil)
	if found {
		return
	}
	foundStart = true
	d.recurse([]*Element{root}, &foundStart, nil, curr)
}

func (d *Dom) addChild(parent *Element, child *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.validateDocument(parent)
	d.validateDocument(child)
	_, found := d.parents[child]
	if found {
		panic(ErrCorruptDom)
	}
	d.parents[child] = parent
}

func (d *Dom) removeChild(parent *Element, child *Element) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.validateDocument(parent)
	d.validateDocument(child)
	p, found := d.parents[child]
	if !found {
		panic(ErrCorruptDom)
	}
	if p != parent {
		panic(ErrCorruptDom)
	}
	delete(d.parents, child)
	if d.focus == child {
		d.focus = nil
		d.FocusNext()
	}
}

func (d *Dom) validateDocument(element *Element) {
	if element.document != d.document {
		panic(ErrDocumentMismatch)
	}
}

func (d *Dom) validateHasRoot() *Element {
	if d.root == nil {
		panic(ErrDocumentRootMissing)
	}
	return d.root
}

func (d *Dom) validateInTree(element *Element) bool {
	for element != nil {
		if element == d.root {
			return true
		}
		element = d.parents[element]
	}
	return false
}

func (d *Dom) setFocus(element *Element) {
	focus := d.focus
	if focus == element {
		return
	}
	if focus != nil {
		focus.properties.focus.Set(false)
		d.document.Dispatch(func() {
			focusLostEvent := NewEvent(&OnFocusLostEvent{})
			d.document.propagateEvent(focus, focusLostEvent)
		})
		d.focus = nil
	}
	if element != nil && element.component.Focusable() {
		d.focus = element
		d.focus.properties.focus.Set(true)
		d.document.Dispatch(func() {
			focusEvent := NewEvent(&OnFocusEvent{})
			d.document.propagateEvent(element, focusEvent)
		})

	}
}

func (d *Dom) findParents(e *Element) []*Element {
	parents := []*Element{}
	for e != nil {
		parents = append(parents, e)
		e = d.parents[e]
	}
	return parents
}

func (d *Dom) recurse(elements []*Element, foundStart *bool, from *Element, to *Element) bool {
	for _, e := range elements {
		if !*foundStart && e == from {
			*foundStart = true
			continue
		}
		if to != nil && to == e {
			return false
		}
		if *foundStart && e.component.Focusable() {
			d.setFocus(e)
			return true
		}
		c := d.recurse(e.properties.children.Value(), foundStart, from, to)
		if c {
			return true
		}
	}
	return false
}
