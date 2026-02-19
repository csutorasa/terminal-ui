package ansi

import (
	"io"
	"strings"
)

// ANSI escape sequence
type EscapeSequence string

func NewEscapeSequence(command Command, arguments ...string) EscapeSequence {
	if command == "" {
		panic("command cannot be empty")
	}
	var sb strings.Builder
	sb.WriteString("\033[")
	for i, a := range arguments {
		sb.WriteString(a)
		if i != len(arguments)-1 {
			sb.WriteString(";")
		}
	}
	sb.WriteString(string(command))
	return EscapeSequence(sb.String())
}

// Implements [io.WriterTo], so result can be printed to the terminal.
func (s EscapeSequence) WriteTo(w io.Writer) (int64, error) {
	if len(s) == 0 {
		return 0, nil
	}
	n, err := w.Write(([]byte)(s))
	return int64(n), err
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (s EscapeSequence) String() string {
	return string(s)
}
