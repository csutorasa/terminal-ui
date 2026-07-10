package document

import (
	"iter"
	"math"

	"github.com/csutorasa/terminal-ui/ansi"
)

// Formatted output, each element of the slice is a line of output
type RenderOutput []ansi.FormattedText

// Gets the [RenderOutput] lines.
func (ro RenderOutput) Lines() iter.Seq[ansi.FormattedText] {
	return func(yield func(ansi.FormattedText) bool) {
		for _, line := range ro {
			if !yield(line) {
				return
			}
		}
	}
}

// Fills the output to the expected size with the padding.
func (ro RenderOutput) Fill(rc RenderContext, padding ansi.FormattedRune) RenderOutput {
	w, h := rc.Size()
	if h == math.MaxInt || w == math.MaxInt {
		panic(ErrInfiniteRenderContext)
	}
	rw := NewRenderWriterFromOutput(rc, ro)
	for i, line := range rw.lines {
		l := line.Len()
		if l < w {
			rw.lines[i] = line.PadRight(w, padding)
		}
	}
	s := padding.Repeat(w)
	for {
		if !rw.WriteLineFormattedString(s) {
			return rw.lines
		}
	}
}

func render(rc RenderContext, c Component) RenderOutput {
	if rc.Empty() {
		return RenderOutput{}
	}
	rw := NewRenderWriter(rc)
	c.Render(rw)
	return rw.lines
}
