package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
)

type RichText struct {
	*SimpleComponent
	text *document.Property[ansi.FormattedText]
}

func NewRichText(element *document.Element) *RichText {
	return &RichText{
		SimpleComponent: NewSimpleComponent(element),
		text:            document.AppendState(element, document.NewPropertyFunc(ansi.FormattedText{}, ansi.FormattedTextEquals)),
	}
}

func (c *Creator) NewRichText() *RichText {
	return Create(c, NewRichText)
}

func (t *RichText) Focusable() bool {
	return false
}

func (t *RichText) SetText(text ansi.FormattedText) *RichText {
	t.text.Set(text)
	return t
}

func (t *RichText) Render(c *document.RenderWriter) {
	text := t.text.Value()
	for line := range text.SplitSeq("\n") {
		if !c.WriteLineFormattedText(line) {
			return
		}
	}
}

func (t *RichText) OnEvent(e *document.Event) {

}
