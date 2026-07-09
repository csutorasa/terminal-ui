package ansi_test

import (
	"testing"

	"github.com/csutorasa/terminal-ui/ansi"
)

var r = ansi.NewFormattedRune('x', new(ansi.Format).Bold().ForegroundColor(ansi.FormatColorRed))

func TestFormattedRuneLen(t *testing.T) {
	assertLen(t, 1, r)
}

func TestFormattedRuneString(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'x', 0x1B, '[', '0', 'm'})

	actual := r.String()

	assertString(t, expected, actual)
}

func TestFormattedRuneToString(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'x', 0x1B, '[', '0', 'm'})

	toString := r.ToString()

	assertString(t, expected, toString.String())
	assertLen(t, 1, toString)
}

func TestFormattedRuneToText(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'x', 0x1B, '[', '0', 'm'})

	toText := r.ToText()

	assertString(t, expected, toText.String())
	assertLen(t, 1, toText)
}

func TestFormattedRuneRepeat(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'x', 'x', 'x', 0x1B, '[', '0', 'm'})

	repeat := r.Repeat(3)

	assertString(t, expected, repeat.String())
	assertLen(t, 3, repeat)
}
