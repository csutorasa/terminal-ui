package components

import (
	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
)

type Scroll struct {
	*WrapperComponent
	keyBinding ScrollKeyBinding
	top        *document.Property[int]
	left       *document.Property[int]
}

func NewScroll(element *document.Element) *Scroll {
	return &Scroll{
		WrapperComponent: NewWrapperComponent(element),
		keyBinding:       DefaultScrollKeyBinding,
		top:              document.AppendState(element, document.NewProperty(0)),
		left:             document.AppendState(element, document.NewProperty(0)),
	}
}

func (c *Creator) NewScroll() *Scroll {
	return Create(c, NewScroll)
}

func (s *Scroll) Focusable() bool {
	return false
}

func (s *Scroll) Render(c *document.RenderWriter) {
	lines := s.Child().Render()
	left := s.left.Value()
	top := s.top.Value()
	maxWidth := 0
	for _, line := range lines {
		l := line.Len()
		if maxWidth < l {
			maxWidth = l
		}
	}
	if maxWidth > 0 && left >= maxWidth {
		s.left.Set(maxWidth - 1)
	}
	for i, line := range lines {
		if i < top {
			continue
		}
		if left > line.Len() {
			c.WriteLineEmpty(1)
		} else {
			c.WriteLineFormattedText(line.SubStr(left, line.Len()))
		}
	}
}

func (s *Scroll) Layout(c document.LayoutContext) {
	c.Add(s.Child(), document.NewInfiniteRenderContext())
}

type ScrollKeyBinding struct {
	Up    *KeyBinding
	Left  *KeyBinding
	Down  *KeyBinding
	Right *KeyBinding
}

var DefaultScrollKeyBinding = ScrollKeyBinding{
	Up:    NewKeyBinding(ansi.KeyFromRune('w')),
	Left:  NewKeyBinding(ansi.KeyFromRune('a')),
	Down:  NewKeyBinding(ansi.KeyFromRune('s')),
	Right: NewKeyBinding(ansi.KeyFromRune('d')),
}

func (s *Scroll) OnEvent(e *document.Event) {
	ke, ok := e.Event().(*document.KeyboardEvent)
	if !ok {
		return
	}
	if s.keyBinding.Up.Matches(ke.Key) {
		s.top.Map(func(oldVal int) int {
			if oldVal > 0 {
				return oldVal - 1
			}
			return oldVal
		})
		e.StopPropagation()
	}
	if s.keyBinding.Left.Matches(ke.Key) {
		s.left.Map(func(oldVal int) int {
			if oldVal > 0 {
				return oldVal - 1
			}
			return oldVal
		})
		e.StopPropagation()
	}
	if s.keyBinding.Down.Matches(ke.Key) {
		s.top.Map(func(oldVal int) int {
			return oldVal + 1
		})
		e.StopPropagation()
	}
	if s.keyBinding.Right.Matches(ke.Key) {
		s.left.Map(func(oldVal int) int {
			return oldVal + 1
		})
		e.StopPropagation()
	}
}
