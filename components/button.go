package components

import (
	"strings"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type ButtonActionHandler = func()

type Button struct {
	*SimpleComponent
	keyBinding   ButtonKeyBinding
	text         *document.Property[string]
	textFormat   *document.Property[*ansi.Format]
	cursorFormat *document.Property[*ansi.Format]
	onAction     ButtonActionHandler
}

func NewButton(element *document.Element) *Button {
	return &Button{
		SimpleComponent: NewSimpleComponent(element),
		keyBinding:      DefaultButtonKeyBinding,
		text:            document.AppendState(element, document.NewProperty("")),
		textFormat:      document.AppendState(element, document.NewPropertyFunc(new(ansi.Format), ansi.FormatEquals)),
		cursorFormat:    document.AppendState(element, document.NewPropertyFunc(new(ansi.Format).BackgroundColor(ansi.FormatColorGreen), ansi.FormatEquals)),
	}
}

func (c *Creator) NewButton() *Button {
	return CreateWithTheme(c, NewButton)
}

func (b *Button) ApplyTheme(t *style.Theme) {
	b.SetTextFormat(t.Input)
	b.SetCursorFormat(t.Cursor)
}

func (b *Button) SetTextFormat(f *ansi.Format) *Button {
	b.textFormat.Set(f)
	return b
}

func (b *Button) SetCursorFormat(f *ansi.Format) *Button {
	b.cursorFormat.Set(f)
	return b
}

func (b *Button) SetText(text string) *Button {
	b.text.Set(text)
	return b
}

func (b *Button) SetOnAction(handler ButtonActionHandler) *Button {
	b.onAction = handler
	return b
}

func (b *Button) Focusable() bool {
	return true
}

func (b *Button) Render(c *document.RenderWriter) {
	text := b.text.Value()
	for line := range strings.SplitSeq(text, "\n") {
		if b.Focused() {
			if !c.WriteLineFormattedString(ansi.NewFormattedString([]rune(line), b.cursorFormat.Value())) {
				return
			}
		} else {
			if !c.WriteLineFormattedString(ansi.NewFormattedString([]rune(line), b.textFormat.Value())) {
				return
			}
		}
	}
}

func (b *Button) OnEvent(e *document.Event) {
	switch e := e.Event().(type) {
	case *document.KeyboardEvent:
		if b.keyBinding.Action.Matches(e.Key) {
			if b.onAction != nil {
				b.onAction()
			}
		}
	}
}

type ButtonKeyBinding struct {
	Action *KeyBinding
}

var DefaultButtonKeyBinding = ButtonKeyBinding{
	Action: NewKeyBinding(ansi.KeyFromRune(' '), ansi.KeyEnter),
}
