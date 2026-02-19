package ansi_test

import (
	"slices"
	"testing"

	"github.com/csutorasa/terminal-ui/ansi"
)

var s = ansi.NewFormattedString([]rune("abcxyz"), ansi.FormatCodeBold, ansi.FormatCodeRedForeground)

func TestFormattedStringLen(t *testing.T) {
	assertLen(t, 6, s)
}

func TestFormattedStringString(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 'x', 'y', 'z', 0x1B, '[', '0', 'm'})

	actual := s.String()

	assertString(t, expected, actual)
}

func TestFormattedStringToText(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 'x', 'y', 'z', 0x1B, '[', '0', 'm'})

	toText := s.ToText()

	assertString(t, expected, toText.String())
	assertLen(t, 6, toText)
}

func TestFormattedStringSubStr(t *testing.T) {
	expected := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'c', 'x', 'y', 0x1B, '[', '0', 'm'})

	substr := s.SubStr(2, 5)

	assertLen(t, 3, substr)
	assertString(t, expected, substr.String())
}

func TestFormattedStringRunesSeq(t *testing.T) {
	expected := string([]rune{'a', 'b', 'c', 'x', 'y', 'z'})

	runes := string(slices.Collect(s.RunesSeq()))

	assertString(t, expected, runes)
}

func TestFormattedStringSplitSeq(t *testing.T) {
	expected0 := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 0x1B, '[', '0', 'm'})
	expected1 := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm'})

	strs := slices.Collect(s.SplitSeq("c"))

	assertInt(t, 2, len(strs))

	assertString(t, expected0, strs[0].String())
	assertLen(t, 2, strs[0])

	assertString(t, expected1, strs[1].String())
	assertLen(t, 3, strs[1])
}

func TestFormattedStringSplitSeqEnd(t *testing.T) {
	expected0 := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 'x', 'y', 0x1B, '[', '0', 'm'})
	expected1 := ""

	strs := slices.Collect(s.SplitSeq("z"))

	assertInt(t, 2, len(strs))

	assertString(t, expected0, strs[0].String())
	assertLen(t, 5, strs[0])

	assertString(t, expected1, strs[1].String())
	assertLen(t, 0, strs[1])
}
