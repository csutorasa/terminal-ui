package ansi

import (
	"io"
	"slices"
)

// Immutable ANSI formatted rune
type FormattedRune struct {
	r      rune
	format *Format
}

func NewSimpleRune(r rune) FormattedRune {
	return NewFormattedRune(r, new(Format))
}

func NewFormattedRune(r rune, format *Format) FormattedRune {
	if r == 0 {
		panic("null character is forbidden")
	}
	if format == nil {
		panic("null format is forbidden")
	}
	return FormattedRune{
		r:      r,
		format: format,
	}
}

// Creates a [FormattedString] by repeating the [FormattedRune].
func (s FormattedRune) Repeat(count int) FormattedString {
	return FormattedString{
		str:    slices.Repeat([]rune{s.r}, count),
		format: s.format,
	}
}

// Creates a [FormattedString] from the [FormattedRune].
func (s FormattedRune) ToString() FormattedString {
	return NewFormattedString([]rune{s.r}, s.format)
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
	return s.format.WriteTo(w, string(s.r))
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (s FormattedRune) String() string {
	return s.format.FormatString(string(s.r))
}

// Returns if the two [FormattedRune]s are equal.
func FormattedRuneEquals(a FormattedRune, b FormattedRune) bool {
	if a.r != b.r {
		return false
	}
	return FormatEquals(a.format, b.format)
}
