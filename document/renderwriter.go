package document

import "github.com/csutorasa/terminal-ui/ansi"

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
	return rw.WriteLineFormattedString(ansi.NewFormattedString(line, new(ansi.Format)))
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
