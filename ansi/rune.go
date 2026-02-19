package ansi

import (
	"fmt"
	"io"
	"slices"
)

// Immutable ANSI formatted rune
type FormattedRune struct {
	r         rune
	arguments []FormatCode
}

func NewFormattedRune(r rune, arguments ...FormatCode) FormattedRune {
	if r == 0 {
		panic("null character is forbidden")
	}
	return FormattedRune{
		arguments: arguments,
		r:         r,
	}
}

// Creates a [FormattedString] by repeating the [FormattedRune].
func (s FormattedRune) Repeat(count int) FormattedString {
	return FormattedString{
		str:       slices.Repeat([]rune{s.r}, count),
		arguments: s.arguments,
	}
}

// Creates a [FormattedString] from the [FormattedRune].
func (s FormattedRune) ToString() FormattedString {
	return NewFormattedString([]rune{s.r}, s.arguments...)
}

// Creates a [FormattedText] from the [FormattedRune].
func (s FormattedRune) ToText() FormattedText {
	return FormattedText{s.ToString()}
}

// Gets the length.
func (s FormattedRune) Len() int {
	return 1
}

// Returns a string without ANSI formatting.
func (s FormattedRune) Unformat() string {
	return string(s.r)
}

// Implements [io.WriterTo], so result can be printed to the terminal.
func (s FormattedRune) WriteTo(w io.Writer) (int64, error) {
	total := int64(0)
	n, err := NewEscapeSequence(CommandFormat, s.arguments...).WriteTo(w)
	total += n
	if err != nil {
		return total, err
	}
	n1, err := w.Write(([]byte)(string(s.r)))
	total += int64(n1)
	if err != nil {
		return total, err
	}
	n, err = NewEscapeSequence(CommandFormat, string(FormatCodeReset)).WriteTo(w)
	total += n
	return total, err
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (s FormattedRune) String() string {
	if len(s.arguments) == 0 {
		return string(s.r)
	}
	return fmt.Sprintf("%s%s%s", NewEscapeSequence(CommandFormat, s.arguments...), string(s.r), NewEscapeSequence(CommandFormat, string(FormatCodeReset)))
}

// Returns if the two [FormattedRune]s are equal.
func FormattedRuneEquals(a FormattedRune, b FormattedRune) bool {
	if a.r != b.r {
		return false
	}
	return FormatCodesEquals(a.arguments, b.arguments)
}
