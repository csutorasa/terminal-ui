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

// Creates a new [Component].
func NewComponent[T Component](document *Document, creator func(element *Element) T) T {
	if document == nil {
		document = Default
	}
	element := &Element{
		document: document,
	}
	children := NewUniqueSlicePropertyOnChange([]*Element{}, func(e *Element) {
		document.dom.addChild(element, e)
	}, func(e *Element) {
		document.dom.removeChild(element, e)
	})
	elementProperty := newElementProperties(children)
	element.properties = elementProperty
	c := creator(element)
	element.component = c
	return c
}

// Get the [Component].
func (e *Element) Component() Component {
	return e.component
}

// Gets the properties of the [Element].
func (e *Element) Properties() *ElementProperties {
	return e.properties
}

// Renders the [Element].
func (e *Element) Render() RenderOutput {
	return e.cachedOutput
}

// Renders the [Element] and fills the missing output with the given rune.
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
	shouldRender := e.cachedOutput == nil || e.properties.Changed() || childrenChanged
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
	c.setCurrentParent(e)
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
