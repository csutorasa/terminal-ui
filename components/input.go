package components

import (
	"slices"
	"unicode"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Input struct {
	*SimpleComponent
	str          *document.Property[[]rune]
	cursor       *document.Property[int]
	textFormat   *document.Property[*ansi.Format]
	cursorFormat *document.Property[*ansi.Format]
}

func NewInput(element *document.Element) *Input {
	return &Input{
		SimpleComponent: NewSimpleComponent(element),
		cursor:          document.AppendState(element, document.NewProperty(0)),
		str:             document.AppendState(element, document.NewSliceProperty([]rune{})),
		textFormat:      document.AppendState(element, document.NewPropertyFunc(new(ansi.Format), ansi.FormatEquals)),
		cursorFormat:    document.AppendState(element, document.NewPropertyFunc(new(ansi.Format).BackgroundColor(ansi.FormatColorGreen), ansi.FormatEquals)),
	}
}

func (i *Input) ApplyTheme(t *style.Theme) {
	i.SetTextFormat(t.Input)
	i.SetCursorFormat(t.Cursor)
}

func (i *Input) SetTextFormat(f *ansi.Format) *Input {
	i.textFormat.Set(f)
	return i
}

func (i *Input) SetCursorFormat(f *ansi.Format) *Input {
	i.cursorFormat.Set(f)
	return i
}

func (i *Input) Focusable() bool {
	return true
}

func (i *Input) Input() []rune {
	return i.str.Value()
}

func (i *Input) OnEvent(e *document.Event) {
	switch e := e.Event().(type) {
	case *document.KeyboardEvent:
		switch e.Key {
		case ansi.KeyBackspace, ansi.KeyBackspace2:
			i.str.Map(func(oldVal []rune) []rune {
				cursor := i.cursor.Value()
				if cursor <= 0 {
					return oldVal
				}
				return slices.Delete(oldVal, cursor-1, cursor)
			})
			i.cursor.Map(func(oldVal int) int {
				if oldVal > len(i.str.Value()) {
					return len(i.str.Value())
				}
				if oldVal <= 0 {
					return oldVal
				}
				return oldVal - 1
			})
		case ansi.KeyDelete, ansi.KeyCtrlD:
			i.str.Map(func(oldVal []rune) []rune {
				cursor := i.cursor.Value()
				if cursor >= len(oldVal) {
					return oldVal
				}
				return slices.Delete(oldVal, cursor, cursor+1)
			})
			i.cursor.Map(func(oldVal int) int {
				if oldVal > len(i.str.Value()) {
					return len(i.str.Value())
				}
				return oldVal
			})
		case ansi.KeyCtrlArrowRight, ansi.KeyAltB:
			i.cursor.Map(func(oldVal int) int {
				str := i.str.Value()
				return startOfNextWordIndex(str, oldVal)
			})
		case ansi.KeyCtrlArrowLeft, ansi.KeyAltF:
			i.cursor.Map(func(oldVal int) int {
				str := i.str.Value()
				return startOfPreviousWordIndex(str, oldVal)
			})
		case ansi.KeyArrowRight, ansi.KeyCtrlF:
			i.cursor.Map(func(oldVal int) int {
				if oldVal < len(i.str.Value()) {
					return oldVal + 1
				}
				return oldVal
			})
		case ansi.KeyArrowLeft, ansi.KeyCtrlB:
			i.cursor.Map(func(oldVal int) int {
				if oldVal > 0 {
					return oldVal - 1
				}
				return oldVal
			})
		case ansi.KeyHome, ansi.KeyCtrlA:
			i.cursor.Set(0)
		case ansi.KeyEnd, ansi.KeyCtrlE:
			str := i.str.Value()
			i.cursor.Set(len(str))
		case ansi.KeyCtrlK:
			i.str.Map(func(oldVal []rune) []rune {
				cursor := i.cursor.Value()
				if cursor == len(oldVal) {
					return oldVal
				}
				return oldVal[:cursor]
			})
		case ansi.KeyCtrlU:
			i.str.Set([]rune{})
			i.cursor.Set(0)
		case ansi.KeyAltD:
			i.str.Map(func(oldVal []rune) []rune {
				oldCursor := i.cursor.Value()
				end := startOfNextWordIndex(oldVal, oldCursor)
				return slices.Delete(oldVal, oldCursor, end)
			})
		case ansi.KeyCtrlW:
			oldCursor := i.cursor.Value()
			i.cursor.Map(func(oldVal int) int {
				str := i.str.Value()
				return startOfPreviousWordIndex(str, oldVal)
			})
			newCursor := i.cursor.Value()
			i.str.Map(func(oldVal []rune) []rune {
				return slices.Delete(oldVal, newCursor, oldCursor)
			})
		case ansi.KeyEnter:
		default:
			if !e.Key.IsEscape() && e.Key >= 0x20 {
				str := i.str.Value()
				cursor := i.cursor.Value()
				if len(str) == cursor {
					i.str.Set(append(str, rune(e.Key)))
				} else {
					i.str.Set(slices.Insert(str, cursor, rune(e.Key)))
				}
				i.cursor.Set(cursor + 1)
			}
		}
	}
}

func startOfPreviousWordIndex(str []rune, currentIndex int) int {
	if currentIndex == 0 {
		return 0
	}
	newVal := currentIndex - 1
	foundRune := false
	for i := newVal; i >= 0; i-- {
		if isLetter(str[i]) {
			newVal = i
			foundRune = true
		} else if foundRune {
			break
		}
	}
	return newVal
}

func startOfNextWordIndex(str []rune, currentIndex int) int {
	if currentIndex == len(str) {
		return currentIndex
	}
	foundOther := false
	newVal := currentIndex
	for i := currentIndex; i < len(str); i++ {
		if isLetter(str[i]) {
			if foundOther {
				newVal = i
				break
			}
		} else {
			foundOther = true
		}
	}
	if newVal == currentIndex {
		return len(str)
	}
	return newVal
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}
