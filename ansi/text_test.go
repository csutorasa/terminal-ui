package ansi_test

import (
	"slices"
	"testing"

	"github.com/csutorasa/terminal-ui/ansi"
)

var text = ansi.FormattedText{
	ansi.NewFormattedString([]rune("abc"), ansi.FormatCodeBold, ansi.FormatCodeRedForeground),
	ansi.NewFormattedString([]rune("xyz"), ansi.FormatCodeItalic, ansi.FormatCodeGreenForeground),
}

func TestFormattedTextLen(t *testing.T) {
	assertLen(t, 6, text)
}

func TestFormattedTextString(t *testing.T) {
	expected := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm',
	})

	actual := text.String()

	assertString(t, expected, actual)
}

func TestFormattedTextSubStr(t *testing.T) {
	expected := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 0x1B, '[', '0', 'm',
	})

	substr := text.SubStr(2, 5)

	assertLen(t, 3, substr)
	assertString(t, expected, substr.String())
}

func TestFormattedTextPadRight(t *testing.T) {
	expected := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', '3', 'm', 'f', 'f', 'f', 0x1B, '[', '0', 'm',
	})

	padRight := text.PadRight(9, ansi.NewFormattedRune('f', ansi.FormatCodeYellowForeground))

	assertLen(t, 9, padRight)
	assertString(t, expected, padRight.String())
}

func TestFormattedTextConcat(t *testing.T) {
	expected := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', '3', 'm', 'f', 'g', 'h', 0x1B, '[', '0', 'm',
	})

	concat := text.Concat(ansi.NewFormattedText([]rune("fgh"), ansi.FormatCodeYellowForeground))

	assertLen(t, 9, concat)
	assertString(t, expected, concat.String())
}

func TestFormattedTextSplitSeq(t *testing.T) {
	expected0 := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 0x1B, '[', '0', 'm'})
	expected1 := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm',
	})

	texts := slices.Collect(text.SplitSeq("b"))

	assertInt(t, 2, len(texts))

	assertInt(t, 1, len(texts[0]))
	assertString(t, expected0, texts[0].String())
	assertLen(t, 1, texts[0])

	assertInt(t, 2, len(texts[1]))
	assertString(t, expected1, texts[1].String())
	assertLen(t, 4, texts[1])
}

func TestFormattedTextSplitSeqMiddle(t *testing.T) {
	expected0 := string([]rune{0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 0x1B, '[', '0', 'm'})
	expected1 := string([]rune{0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 'z', 0x1B, '[', '0', 'm'})

	texts := slices.Collect(text.SplitSeq("c"))

	assertInt(t, 2, len(texts))

	assertInt(t, 1, len(texts[0]))
	assertString(t, expected0, texts[0].String())
	assertLen(t, 2, texts[0])

	assertInt(t, 2, len(texts[1]))
	assertString(t, expected1, texts[1].String())
	assertLen(t, 3, texts[1])
}

func TestFormattedTextSplitSeqEnd(t *testing.T) {
	expected0 := string([]rune{
		0x1B, '[', '1', ';', '3', '1', 'm', 'a', 'b', 'c', 0x1B, '[', '0', 'm',
		0x1B, '[', '3', ';', '3', '2', 'm', 'x', 'y', 0x1B, '[', '0', 'm',
	})
	expected1 := ""

	texts := slices.Collect(text.SplitSeq("z"))

	assertInt(t, 2, len(texts))

	assertInt(t, 2, len(texts[0]))
	assertString(t, expected0, texts[0].String())
	assertLen(t, 5, texts[0])

	assertInt(t, 1, len(texts[1]))
	assertString(t, expected1, texts[1].String())
	assertLen(t, 0, texts[1])
}
