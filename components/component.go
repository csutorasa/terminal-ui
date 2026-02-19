package components

import (
	"iter"

	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type ApplyThemer interface {
	document.Component
	// Applies the theme to the component.
	ApplyTheme(*style.Theme)
}

// Base struct for a component.
type ComponentBase struct {
	element *document.Element
}

func NewComponentBase(element *document.Element) *ComponentBase {
	return &ComponentBase{
		element: element,
	}
}

// Returns the [document.Element].
func (c *ComponentBase) Element() *document.Element {
	return c.element
}

// Gets if the element is focused.
func (c *ComponentBase) Focused() bool {
	return c.element.Properties().Focused()
}

// Component with no children.
type Component struct {
	*ComponentBase
}

func NewComponent(element *document.Element) *Component {
	return &Component{
		ComponentBase: NewComponentBase(element),
	}
}

// Defaults to empty Layout, as it has no children.
func (c *Component) Layout(rc document.LayoutContext) {

}

// Component with a single child.
type WrapperComponent struct {
	*ComponentBase
}

func NewWrapperComponent(element *document.Element) *WrapperComponent {
	return &WrapperComponent{
		ComponentBase: NewComponentBase(element),
	}
}

// Gets the child.
func (wc *WrapperComponent) Child() *document.Element {
	children := wc.element.Properties().Children().Value()
	if len(children) == 0 {
		return nil
	}
	return children[0]
}

// Sets the child.
func (wc *WrapperComponent) SetChild(e *document.Element) {
	children := wc.element.Properties().Children()
	children.RemoveAll()
	if e != nil {
		children.Add(e)
	}
}

// Component with children.
type ContainerComponent struct {
	*ComponentBase
}

func NewContainerComponent(element *document.Element) *ContainerComponent {
	return &ContainerComponent{
		ComponentBase: NewComponentBase(element),
	}
}

// Adds children.
func (cc *ContainerComponent) AddChildren(elements ...*document.Element) {
	cc.element.Properties().Children().Add(elements...)
}

// Removes children.
func (cc *ContainerComponent) RemoveChildren(elements ...*document.Element) {
	cc.element.Properties().Children().Remove(elements...)
}

// Gets the children.
func (cc *ContainerComponent) Children() iter.Seq[*document.Element] {
	return func(yield func(*document.Element) bool) {
		for _, c := range cc.element.Properties().Children().Value() {
			if !yield(c) {
				return
			}
		}
	}
}
