package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
)

type Text struct {
	*Component
	text *document.Property[ansi.FormattedText]
}

func NewText(element *document.Element) *Text {
	return &Text{
		Component: NewComponent(element),
		text:      document.AppendState(element, document.NewPropertyFunc(ansi.FormattedText{}, ansi.FormattedTextEquals)),
	}
}

func (c *Creator) NewText() *Text {
	return Create(c, NewText)
}

func (t *Text) Focusable() bool {
	return false
}

func (t *Text) SetText(text ansi.FormattedText) *Text {
	t.text.Set(text)
	return t
}

func (t *Text) Render(c *document.RenderWriter) {
	text := t.text.Value()
	for line := range text.SplitSeq("\n") {
		if !c.WriteLineFormattedText(line) {
			return
		}
	}
}

func (t *Text) OnEvent(e *document.Event) {

}
