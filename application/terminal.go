package application

import (
	"errors"
	"fmt"
	"os"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"golang.org/x/term"
)

var ErrIsNotTerminal = fmt.Errorf("is not terminal")

// A [Renderer] for the terminal.
type TerminalRenderer struct {
	out      *os.File
	terminal *Terminal
}

func NewTerminalRenderer(out *os.File) Renderer {
	return &TerminalRenderer{
		out:      out,
		terminal: NewTerminal(out),
	}
}

// Implements [Renderer].
func (tr *TerminalRenderer) Init() error {
	err := tr.terminal.MakeRaw()
	if err != nil {
		return err
	}
	_, err1 := ansi.DisablePrivateMode(ansi.PrivateModeCursorVisible).WriteTo(tr.out)
	_, err2 := ansi.EnablePrivateMode(ansi.PrivateModeAlternativeBuffer).WriteTo(tr.out)
	_, err3 := ansi.EnablePrivateMode(ansi.PrivateModeAlternativeScroll).WriteTo(tr.out)
	return errors.Join(err1, err2, err3)
}

// Implements [Renderer].
func (tr *TerminalRenderer) Close() error {
	err := tr.terminal.Close()
	_, err1 := ansi.NewEscapeSequence(ansi.CommandEraseScreen, string(ansi.EraseScreenEntire)).WriteTo(tr.out)
	_, err2 := ansi.NewEscapeSequence(ansi.CommandCursorMove).WriteTo(tr.out)
	_, err3 := ansi.EnablePrivateMode(ansi.PrivateModeCursorVisible).WriteTo(tr.out)
	_, err4 := ansi.DisablePrivateMode(ansi.PrivateModeAlternativeScroll).WriteTo(tr.out)
	_, err5 := ansi.DisablePrivateMode(ansi.PrivateModeAlternativeBuffer).WriteTo(tr.out)
	return errors.Join(err, err1, err2, err3, err4, err5)
}

// Implements [Renderer].
func (tr *TerminalRenderer) Render(renderDocument func(document.RenderContext) (document.RenderOutput, bool)) error {
	_, err := ansi.NewEscapeSequence(ansi.CommandEraseScreen, string(ansi.EraseScreenEntire)).WriteTo(tr.out)
	if err != nil {
		return err
	}
	_, err = ansi.NewEscapeSequence(ansi.CommandCursorMove).WriteTo(tr.out)
	if err != nil {
		return err
	}
	width, height, err := term.GetSize(int(tr.out.Fd()))
	if err != nil {
		return err
	}
	c, _ := renderDocument(document.NewRenderContext(width, height))
	for line := range c.Lines() {
		if line.Len() == width {
			_, err := fmt.Fprint(tr.out, line)
			if err != nil {
				return err
			}
		} else {
			_, err := fmt.Fprintf(tr.out, "%s\r\n", line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// An [EventDecoder] for the terminal.
type TerminalDecoder struct {
	in       *os.File
	terminal *Terminal
	quitKey  ansi.Input
}

func NewTerminalDecoder(in *os.File) EventDecoder {
	return &TerminalDecoder{
		in:       in,
		terminal: NewTerminal(in),
		quitKey:  ansi.KeyCtrlC,
	}
}

// Implements [EventDecoder].
func (td *TerminalDecoder) Start(events chan<- Event) error {
	err := td.terminal.MakeRaw()
	if err != nil {
		return err
	}
	decoder := ansi.NewInputDecoder(td.in)
	go func() {
		for {
			i, err := decoder.Decode()
			if err != nil {
				return
			}
			event := NewEvent(i)
			if i == td.quitKey {
				event.Quit()
			}
			events <- event
		}
	}()
	return nil
}

// Implements [EventDecoder].
func (td *TerminalDecoder) Close() error {
	return td.terminal.Close()
}

// A helper for managing the terminal.
type Terminal struct {
	f        *os.File
	fd       int
	oldState *term.State
}

func NewTerminal(f *os.File) *Terminal {
	return &Terminal{
		f:  f,
		fd: int(f.Fd()),
	}
}

// Enables raw mode of the terminal.
func (t *Terminal) MakeRaw() error {
	if !term.IsTerminal(t.fd) {
		return ErrIsNotTerminal
	}
	var err error
	t.oldState, err = term.MakeRaw(t.fd)
	return err
}

// Implements [io.Closer], to restore the terminal to previous state (cooked mode).
func (t *Terminal) Close() error {
	return term.Restore(t.fd, t.oldState)
}
