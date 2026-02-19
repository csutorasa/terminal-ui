package document

import (
	"iter"
	"math"

	"github.com/csutorasa/terminal-ui/ansi"
)

// Formatted output
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
	if rc.h == math.MaxInt || rc.w == math.MaxInt {
		panic("cannot fill infinite render context")
	}
	rw := NewRenderWriterFromOutput(rc, ro)
	for i, line := range rw.lines {
		l := line.Len()
		if l < rw.w {
			rw.lines[i] = line.PadRight(rw.w, padding)
		}
	}
	s := padding.Repeat(rw.w)
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
func (w *RenderWriter) WriteLineRunes(line []rune) bool {
	return w.WriteLineFormattedString(ansi.NewFormattedString(line))
}

// Writes [ansi.FormattedString] to the output.
// Returns if the context can accept more lines.
func (w *RenderWriter) WriteLineFormattedString(line ansi.FormattedString) bool {
	return w.WriteLineFormattedText(ansi.FormattedText{line})
}

// Writes [ansi.FormattedText] to the output.
// Returns if the context can accept more lines.
func (w *RenderWriter) WriteLineFormattedText(line ansi.FormattedText) bool {
	if len(w.lines) >= w.h {
		return false
	}
	if line.Len() > w.w {
		line = line.SubStr(0, w.w)
	}
	w.lines = append(w.lines, line)
	return len(w.lines) < w.h
}

// Writes count number of empty lines.
// Returns if the context can accept more lines.
func (w *RenderWriter) WriteLineEmpty(count int) bool {
	for range count {
		if !w.WriteLineRunes([]rune{}) {
			return false
		}
	}
	return true
}

func render(rc RenderContext, component Component) RenderOutput {
	rw := NewRenderWriter(rc)
	component.Render(rw)
	return rw.lines
}
