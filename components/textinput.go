package components

import (
	"slices"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
)

type TextInput struct {
	*Input
}

func NewTextInput(element *document.Element) *TextInput {
	return &TextInput{
		Input: NewInput(element),
	}
}

func (c *Creator) NewTextInput() *TextInput {
	return CreateWithTheme(c, NewTextInput)
}

func (t *TextInput) Render(c *document.RenderWriter) {
	w, _ := c.Size()
	width := min(w, 80)
	var text []rune
	str := t.Input.str.Value()
	if width > len(str) {
		text = append(str, slices.Repeat([]rune{' '}, width-len(str))...)
	} else {
		text = str[:min(len(str), width)]
	}
	if !t.Focused() {
		formattedStr := ansi.NewFormattedString(text, t.background.Value(), t.foreground.Value())
		c.WriteLineFormattedString(formattedStr)
		return
	}
	cursor := t.cursor.Value()
	if cursor < width {
		formattedText := ansi.FormattedText{
			ansi.NewFormattedString(text[:cursor], t.background.Value(), t.foreground.Value()),
			ansi.NewFormattedRune(text[cursor], t.cursorColor.Value()...).ToString(),
			ansi.NewFormattedString(text[cursor+1:], t.background.Value(), t.foreground.Value()),
		}
		c.WriteLineFormattedText(formattedText)
	} else {
		formattedStr := ansi.NewFormattedString(text, t.background.Value(), t.foreground.Value())
		c.WriteLineFormattedString(formattedStr)
	}
}
