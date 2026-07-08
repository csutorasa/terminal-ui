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
		panic("cannot fill infinite render context")
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

type RenderWriter struct {
	RenderContext
	lines RenderOutput
}

func NewRenderWriter(rc RenderContext) *RenderWriter {
	return NewRenderWriterFromOutput(rc, RenderOutput{})
}

func NewRenderWriterFromOutput(rc RenderContext, ro RenderOutput) *RenderWriter {
	return &RenderWriter{
		RenderContext: rc,
		lines:         ro,
	}
}

// Writes runes to the output.
// Returns if the context can accept more lines.
func (rw *RenderWriter) WriteLineRunes(line []rune) bool {
	return rw.WriteLineFormattedString(ansi.NewFormattedString(line))
}

// Writes [ansi.FormattedString] to the output.
// Returns if the context can accept more lines.
func (rw *RenderWriter) WriteLineFormattedString(line ansi.FormattedString) bool {
	return rw.WriteLineFormattedText(ansi.FormattedText{line})
}

// Writes [ansi.FormattedText] to the output.
// Returns if the context can accept more lines.
func (rw *RenderWriter) WriteLineFormattedText(line ansi.FormattedText) bool {
	w, h := rw.Size()
	if len(rw.lines) >= h {
		return false
	}
	if line.Len() > w {
		line = line.SubStr(0, w)
	}
	rw.lines = append(rw.lines, line)
	return len(rw.lines) < h
}

// Writes count number of empty lines.
// Returns if the context can accept more lines.
func (rw *RenderWriter) WriteLineEmpty(count int) bool {
	for range count {
		if !rw.WriteLineRunes([]rune{}) {
			return false
		}
	}
	return true
}

func render(rc RenderContext, c Component) RenderOutput {
	rw := NewRenderWriter(rc)
	c.Render(rw)
	return rw.lines
}
