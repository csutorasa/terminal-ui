package ansi_test

import (
	"testing"

	"github.com/csutorasa/terminal-ui/ansi"
)

func TestEscapeSequenceNoArgument(t *testing.T) {
	expected := string([]rune{0x1B, '[', 'J'})
	actual := ansi.NewEscapeSequence(ansi.CommandEraseScreen).String()
	assertString(t, expected, actual)
}

func TestEscapeSequenceSingleArgument(t *testing.T) {
	expected := string([]rune{0x1B, '[', '3', 'A'})
	actual := ansi.NewEscapeSequence(ansi.CommandCursorMoveUp, "3").String()
	assertString(t, expected, actual)
}

func TestEscapeSequenceMultipleArguments(t *testing.T) {
	expected := string([]rune{0x1B, '[', '5', ';', '1', '0', 'H'})
	actual := ansi.NewEscapeSequence(ansi.CommandCursorMove, "5", "10").String()
	assertString(t, expected, actual)
}
