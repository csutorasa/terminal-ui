package ansi

import (
	"fmt"
	"io"
	"iter"
	"slices"
	"strings"
)

// Immutable ANSI formatted string
type FormattedString struct {
	str       []rune
	arguments []FormatCode
}

func NewFormattedString(str []rune, arguments ...FormatCode) FormattedString {
	return FormattedString{
		arguments: arguments,
		str:       str,
	}
}

// Splits the [FormattedString] into [FormattedString]s by the given separator.
func (s FormattedString) SplitSeq(sep string) iter.Seq[FormattedString] {
	return func(yield func(FormattedString) bool) {
		for line := range strings.SplitSeq(string(s.str), sep) {
			str := FormattedString{
				str:       []rune(line),
				arguments: s.arguments,
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
		str:       s.str[from:to],
		arguments: s.arguments,
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
	if len(s.str) == 0 {
		return 0, nil
	}
	if len(s.arguments) == 0 {
		n, err := w.Write(([]byte)(string(s.str)))
		return int64(n), err
	}
	total := int64(0)
	n, err := NewEscapeSequence(CommandFormat, s.arguments...).WriteTo(w)
	total += n
	if err != nil {
		return total, err
	}
	n1, err := w.Write(([]byte)(string(s.str)))
	total += int64(n1)
	if err != nil {
		return total, err
	}
	n, err = NewEscapeSequence(CommandFormat, string(FormatCodeReset)).WriteTo(w)
	total += n
	return total, err
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (s FormattedString) String() string {
	if len(s.str) == 0 {
		return ""
	}
	if len(s.arguments) == 0 {
		return string(s.str)
	}
	return fmt.Sprintf("%s%s%s", NewEscapeSequence(CommandFormat, s.arguments...), string(s.str), NewEscapeSequence(CommandFormat, string(FormatCodeReset)))
}

// Returns if the two [FormattedString]s are equal.
func FormattedStringEquals(a FormattedString, b FormattedString) bool {
	if !slices.Equal(a.str, b.str) {
		return false
	}
	return FormatCodesEquals(a.arguments, b.arguments)
}
