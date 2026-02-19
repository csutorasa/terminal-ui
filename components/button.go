package components

import (
	"strings"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type ButtonActionHandler = func()

type Button struct {
	*Component
	keyBinding  ButtonKeyBinding
	text        *document.Property[string]
	foreground  *document.Property[ansi.FormatCode]
	background  *document.Property[ansi.FormatCode]
	cursorColor *document.Property[[]ansi.FormatCode]
	onAction    ButtonActionHandler
}

func NewButton(element *document.Element) *Button {
	return &Button{
		Component:   NewComponent(element),
		keyBinding:  DefaultButtonKeyBinding,
		text:        document.AppendState(element, document.NewProperty("")),
		foreground:  document.AppendState(element, document.NewProperty(ansi.FormatCodeDefaultForeground)),
		background:  document.AppendState(element, document.NewProperty(ansi.FormatCodeDefaultBackground)),
		cursorColor: document.AppendState(element, document.NewSliceProperty([]ansi.FormatCode{ansi.FormatCodeGreenBackground})),
	}
}

func (c *Creator) NewButton() *Button {
	return CreateWithTheme(c, NewButton)
}

func (b *Button) ApplyTheme(t *style.Theme) {
	b.SetForeground(t.InputForeground)
	b.SetBackground(t.InputBackground)
	b.SetCursorColor(t.Cursor)
}

func (b *Button) SetForeground(c ansi.FormatCode) *Button {
	b.foreground.Set(c)
	return b
}

func (b *Button) SetBackground(c ansi.FormatCode) *Button {
	b.background.Set(c)
	return b
}

func (b *Button) SetCursorColor(c []ansi.FormatCode) *Button {
	b.cursorColor.Set(c)
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
			if !c.WriteLineFormattedString(ansi.NewFormattedString([]rune(line), b.cursorColor.Value()...)) {
				return
			}
		} else {
			if !c.WriteLineFormattedString(ansi.NewFormattedString([]rune(line), b.background.Value(), b.foreground.Value())) {
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
