package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Border struct {
	*WrapperComponent
	background *document.Property[ansi.FormatCode]
	border     *document.Property[ansi.FormattedRune]
	thickness  *document.Property[int]
}

func NewBorder(element *document.Element) *Border {
	return &Border{
		WrapperComponent: NewWrapperComponent(element),
		border:           document.AppendState(element, document.NewPropertyFunc(ansi.NewFormattedRune('*'), ansi.FormattedRuneEquals)),
		thickness:        document.AppendState(element, document.NewProperty(1)),
		background:       document.AppendState(element, document.NewProperty(ansi.FormatCodeDefaultBackground)),
	}
}

func (c *Creator) NewBorder() *Border {
	return CreateWithTheme(c, NewBorder)
}

func (b *Border) ApplyTheme(t *style.Theme) {
	b.SetBorder(t.Border)
	b.SetBackground(t.DefaultBackground)
}

func (b *Border) SetBorder(border ansi.FormattedRune) *Border {
	b.border.Set(border)
	return b
}

func (b *Border) SetBackground(background ansi.FormatCode) *Border {
	b.background.Set(background)
	return b
}

func (b *Border) SetThickness(thickness int) *Border {
	if thickness < 1 {
		panic("thickness must be positive")
	}
	b.thickness.Set(thickness)
	return b
}

func (b *Border) Focusable() bool {
	return false
}

func (b *Border) Render(c *document.RenderWriter) {
	w, h := c.Size()
	if w < 1 || h < 1 {
		return
	}
	border := b.border.Value()
	thickness := b.thickness.Value()
	if w <= 2*thickness || h <= 2*thickness {
		for range h {
			c.WriteLineFormattedString(border.Repeat(w))
		}
		return
	}
	cc := b.Child().RenderFill(ansi.NewFormattedRune(' ', b.background.Value()))
	for range thickness {
		c.WriteLineFormattedString(border.Repeat(w))
	}
	for line := range cc.Lines() {
		b := border.Repeat(thickness).ToText()
		c.WriteLineFormattedText(b.Concat(line).Concat(b))
	}
	for range thickness {
		c.WriteLineFormattedString(border.Repeat(w))
	}
}

func (b *Border) Layout(c document.LayoutContext) {
	w, h := c.Size()
	thickness := b.thickness.Value()
	if w <= 2*thickness || h <= 2*thickness {
		c.Add(b.Child(), document.NewEmptyRenderContext())
		return
	}
	c.Add(b.Child(), document.NewRenderContext(w-2*thickness, h-2*thickness))
}

func (b *Border) OnEvent(e *document.Event) {

}
