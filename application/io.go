package application

import (
	"fmt"
	"io"

	"github.com/csutorasa/terminal-ui/document"
)

// Application output renderer.
type Renderer interface {
	io.Closer
	// Initializes the renderer.
	Init() error
	// Renders the output.
	Render(func(document.RenderContext) (document.RenderOutput, bool)) error
}

// Application input event decoder.
type EventDecoder interface {
	io.Closer
	// Starts reading events and send them to the channel.
	Start(chan<- Event) error
}

// Default [io.Writer] renderer.
type WriterRenderer struct {
	w      io.Writer
	width  int
	height int
}

// Create a renderer with predefined size.
func NewWriterRenderer(w io.Writer, width int, height int) Renderer {
	return &WriterRenderer{
		w:      w,
		width:  width,
		height: height,
	}
}

// Implements [Renderer].
func (wr *WriterRenderer) Init() error {
	return nil
}

// Implements [Renderer].
func (wr *WriterRenderer) Close() error {
	return nil
}

// Implements [Renderer].
func (wr *WriterRenderer) Render(renderDocument func(document.RenderContext) (document.RenderOutput, bool)) error {
	c, changed := renderDocument(document.NewRenderContext(wr.width, wr.height))
	if !changed {
		return nil
	}
	for line := range c.Lines() {
		_, err := fmt.Fprintf(wr.w, "%s\r\n", line.Unformat())
		if err != nil {
			return err
		}
	}
	return nil
}
