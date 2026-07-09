package ansi

import (
	"io"
	"iter"
	"slices"
	"strings"
)

// Immutable ANSI formatted string
type FormattedString struct {
	str    []rune
	format *Format
}

func NewSimpleString(str []rune) FormattedString {
	return NewFormattedString(str, new(Format))
}

func NewFormattedString(str []rune, format *Format) FormattedString {
	return FormattedString{
		str:    str,
		format: format,
	}
}

// Splits the [FormattedString] into [FormattedString]s by the given separator.
func (s FormattedString) SplitSeq(sep string) iter.Seq[FormattedString] {
	return func(yield func(FormattedString) bool) {
		for line := range strings.SplitSeq(string(s.str), sep) {
			str := FormattedString{
				str:    []rune(line),
				format: s.format,
			}
			if !yield(str) {
				return
			}
		}
	}
}

// Returns the runes from the [FormattedString].
func (s FormattedString) RunesSeq() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s.str {
			if !yield(r) {
				return
			}
		}
	}
}

// Returns the substring of the [FormattedString].
func (s FormattedString) SubStr(from int, to int) FormattedString {
	return FormattedString{
		str:    s.str[from:to],
		format: s.format,
	}
}

// Creates a [FormattedText] from the [FormattedString].
func (s FormattedString) ToText() FormattedText {
	return FormattedText{s}
}

// Gets the length.
func (s FormattedString) Len() int {
	return len(s.str)
}

// Returns a string without ANSI formatting.
func (s FormattedString) Unformat() string {
	return string(s.str)
}

// Implements [io.WriterTo], so result can be printed to the terminal.
func (s FormattedString) WriteTo(w io.Writer) (int64, error) {
	return s.format.WriteTo(w, string(s.str))
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (s FormattedString) String() string {
	return s.format.FormatString(string(s.str))
}

// Returns if the two [FormattedString]s are equal.
func FormattedStringEquals(a FormattedString, b FormattedString) bool {
	if !slices.Equal(a.str, b.str) {
		return false
	}
	return FormatEquals(a.format, b.format)
}
