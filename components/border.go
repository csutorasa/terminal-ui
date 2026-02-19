package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Border struct {
	*WrapperComponent
	border     *document.Property[ansi.FormattedRune]
	background *document.Property[ansi.FormatCode]
}

func NewBorder(element *document.Element) *Border {
	return &Border{
		WrapperComponent: NewWrapperComponent(element),
		border:           document.AppendState(element, document.NewPropertyFunc(ansi.NewFormattedRune('*'), ansi.FormattedRuneEquals)),
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

func (b *Border) Focusable() bool {
	return false
}

func (b *Border) Render(c *document.RenderWriter) {
	w, h := c.Size()
	if w < 1 || h < 1 {
		return
	}
	border := b.border.Value()
	if w < 3 {
		for range h {
			c.WriteLineFormattedString(border.Repeat(w))
		}
		return
	}
	if h < 3 {
		for range h {
			c.WriteLineFormattedString(border.Repeat(w))
		}
		return
	}
	cc := b.Child().RenderFill(ansi.NewFormattedRune(' ', b.background.Value()))
	c.WriteLineFormattedString(border.Repeat(w))
	for line := range cc.Lines() {
		c.WriteLineFormattedText(border.ToText().Concat(line).Concat(border.ToText()))
	}
	c.WriteLineFormattedString(border.Repeat(w))
}

func (b *Border) Layout(c document.LayoutContext) {
	w, h := c.Size()
	if w < 3 || h < 3 {
		c.Add(b.Child(), document.NewEmptyRenderContext())
		return
	}
	c.Add(b.Child(), document.NewRenderContext(w-2, h-2))
}

func (b *Border) OnEvent(e *document.Event) {

}
