package document

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/debugging"
)

// DOM Element.
type Element struct {
	document      *Document
	component     Component
	properties    *ElementProperties
	cachedContext RenderContext
	cachedOutput  RenderOutput
}

// DOM Element properties.
type ElementProperties struct {
	children      *UniqueSliceProperty[*Element]
	focus         *Property[bool]
	renderContext *Property[RenderContext]
	state         *MultiState
}

func newElementProperties(children *UniqueSliceProperty[*Element]) *ElementProperties {
	state := NewMultiState()
	return &ElementProperties{
		children:      appendState(state, children),
		focus:         appendState(state, NewProperty(false)),
		renderContext: appendState(state, NewProperty(RenderContext{})),
		state:         state,
	}
}

// Gets the property to read and write the children of the element.
func (ep *ElementProperties) Children() *UniqueSliceProperty[*Element] {
	return ep.children
}

// Gets if the element has focus.
func (ep *ElementProperties) Focused() bool {
	return ep.focus.Value()
}

func (ep *ElementProperties) changed() bool {
	return ep.state.changed()
}

func (ep *ElementProperties) resetChanged() {
	ep.state.resetChanged()
}

// Appends a new state to the state of the element and returns it.
func AppendState[S State](element *Element, s S) S {
	element.properties.state.Append(s)
	return s
}

// Creates a new [Component].
func NewComponent[T Component](document *Document, creator func(element *Element) T) T {
	if document == nil {
		document = Default
	}
	element := &Element{
		document: document,
	}
	children := NewUniqueSlicePropertyOnChange(func(e *Element) {
		document.addChild(element, e)
	}, func(e *Element) {
		document.removeChild(element, e)
	})
	elementProperty := newElementProperties(children)
	element.properties = elementProperty
	c := creator(element)
	element.component = c
	return c
}

// Gets the properties of the element.
func (e *Element) Properties() *ElementProperties {
	return e.properties
}

// Renders the element.
func (e *Element) Render() RenderOutput {
	return e.cachedOutput
}

// Renders the element and fills the missing output with the given rune.
func (e *Element) RenderFill(r ansi.FormattedRune) RenderOutput {
	return e.cachedOutput.Fill(e.cachedContext, r)
}

func (e *Element) render() bool {
	c := e.properties.renderContext.Value()
	childrenChanged := false
	for _, child := range e.properties.children.Value() {
		changed := child.render()
		if changed {
			childrenChanged = true
		}
	}
	shouldRender := e.cachedOutput == nil || e.properties.changed() || childrenChanged
	if !shouldRender {
		return false
	}
	debugging.LogWithType("render", e.component)
	e.cachedOutput = render(c, e.component)
	e.cachedContext = c
	e.properties.resetChanged()
	return true
}

func (e *Element) renderLayout(c LayoutContext) {
	e.component.Layout(c)
	for _, child := range e.properties.children.Value() {
		cl, ok := c.layout[child]
		if !ok {
			panic("not found")
		}
		ctx := NewLayoutContext(cl)
		ctx.layout = c.layout
		child.renderLayout(ctx)
	}
}

func (e *Element) onEvent(ev *Event) {
	e.component.OnEvent(ev)
}
