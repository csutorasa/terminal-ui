package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Border struct {
	*WrapperComponent
	background *document.Property[ansi.FormattedRune]
	border     *document.Property[style.Directional[ansi.FormattedRune]]
	thickness  *document.Property[style.Directional[uint8]]
}

func NewBorder(element *document.Element) *Border {
	return &Border{
		WrapperComponent: NewWrapperComponent(element),
		border:           document.AppendState(element, document.NewPropertyFunc(style.NewDirectional(ansi.NewSimpleRune('*')), style.DirectionalEqualsFunc(ansi.FormattedRuneEquals))),
		thickness:        document.AppendState(element, document.NewPropertyFunc(style.NewDirectional(uint8(1)), style.DirectionalEquals)),
		background:       document.AppendState(element, document.NewPropertyFunc(ansi.NewSimpleRune(' '), ansi.FormattedRuneEquals)),
	}
}

func (c *Creator) NewBorder() *Border {
	return CreateWithTheme(c, NewBorder)
}

func (b *Border) ApplyTheme(t *style.Theme) {
	b.SetBorder(style.NewDirectional(t.Border))
	b.SetBackground(t.Background)
}

func (b *Border) SetBorder(border style.Directional[ansi.FormattedRune]) *Border {
	b.border.Set(border)
	return b
}

func (b *Border) SetBackground(background ansi.FormattedRune) *Border {
	b.background.Set(background)
	return b
}

func (b *Border) SetThickness(thickness style.Directional[uint8]) *Border {
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
	leftBorder, topBorder, rightBorder, bottomBorder := border.Get()
	thickness := b.thickness.Value()
	leftThickness, topThickness, rightThickness, bottomThickness := thickness.Get()
	cc := b.Child().RenderFill(b.background.Value())
	for range topThickness {
		c.WriteLineFormattedString(topBorder.Repeat(w))
	}
	for line := range cc.Lines() {
		left := leftBorder.Repeat(int(leftThickness)).ToText()
		right := rightBorder.Repeat(int(rightThickness)).ToText()
		c.WriteLineFormattedText(left.Concat(line).Concat(right))
	}
	for range bottomThickness {
		c.WriteLineFormattedString(bottomBorder.Repeat(w))
	}
}

func (b *Border) Layout(c document.LayoutContext) {
	w, h := c.Size()
	thickness := b.thickness.Value()
	leftThickness, topThickness, rightThickness, bottomThickness := thickness.Get()
	hThickness := int(leftThickness) + int(rightThickness)
	vThickness := int(topThickness) + int(bottomThickness)
	if w <= hThickness || h <= vThickness {
		c.Add(b.Child(), document.NewEmptyRenderContext())
		return
	}
	c.Add(b.Child(), document.NewRenderContext(w-hThickness, h-vThickness))
}

func (b *Border) OnEvent(e *document.Event) {

}
