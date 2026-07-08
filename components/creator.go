package components

import (
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

// A helper for creating components.
type Creator struct {
	d *document.Document
	t *style.Theme
}

func NewCreator(d *document.Document, t *style.Theme) *Creator {
	return &Creator{
		d: d,
		t: t,
	}
}

func NewDefaultCreator() *Creator {
	return &Creator{
		d: document.Default,
		t: style.DefaultTheme,
	}
}

// Sets the document root.
func (c *Creator) SetRoot(root Component) {
	c.d.SetRoot(root.Element())
}

// Gets the document.
func (c *Creator) Document() *document.Document {
	return c.d
}

// Gets the theme.
func (c *Creator) Theme() *style.Theme {
	return c.t
}

// Creates a new component.
func Create[T document.Component](c *Creator, new func(element *document.Element) T) T {
	return document.NewComponent(c.d, new)
}

// Creates a new component and applies the theme.
func CreateWithTheme[T ApplyThemer](c *Creator, new func(element *document.Element) T) T {
	element := Create(c, new)
	element.ApplyTheme(c.t)
	return element
}
