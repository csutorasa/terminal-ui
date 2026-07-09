package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Text struct {
	*SimpleComponent
	text       *document.Property[string]
	textFormat *document.Property[*ansi.Format]
}

func NewText(element *document.Element) *Text {
	return &Text{
		SimpleComponent: NewSimpleComponent(element),
		text:            document.AppendState(element, document.NewProperty("")),
		textFormat:      document.AppendState(element, document.NewPropertyFunc(new(ansi.Format), ansi.FormatEquals)),
	}
}

func (c *Creator) NewText() *Text {
	return CreateWithTheme(c, NewText)
}

func (b *Text) ApplyTheme(t *style.Theme) {
	b.SetTextFormat(t.Default)
}

func (t *Text) Focusable() bool {
	return false
}

func (t *Text) SetText(text string) *Text {
	t.text.Set(text)
	return t
}

func (t *Text) SetTextFormat(f *ansi.Format) *Text {
	t.textFormat.Set(f)
	return t
}

func (t *Text) Render(c *document.RenderWriter) {
	text := t.text.Value()
	format := t.textFormat.Value()
	formattedText := ansi.NewFormattedText([]rune(text), format)
	for line := range formattedText.SplitSeq("\n") {
		if !c.WriteLineFormattedText(line) {
			return
		}
	}
}

func (t *Text) OnEvent(e *document.Event) {

}
