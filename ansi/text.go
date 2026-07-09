package ansi

import (
	"io"
	"iter"
	"slices"
	"strings"
)

// ANSI formatted text
type FormattedText []FormattedString

func NewSimpleText(str []rune) FormattedText {
	return NewFormattedText(str, new(Format))
}

func NewFormattedText(str []rune, format *Format) FormattedText {
	return FormattedText{NewFormattedString(str, format)}
}

// Returns the substring of the [FormattedText].
func (t FormattedText) SubStr(from int, to int) FormattedText {
	if from < 0 {
		panic("from is out of bounds")
	}
	if to < 0 {
		panic("to is out of bounds")
	}
	if from > to {
		panic("from is greater than to")
	}
	n := FormattedText{}
	total := 0
	for _, s := range t {
		l := s.Len()
		if total > from { // already writing
			if total+l < to { // can fully write
				n = append(n, s)
			} else { // finish writing
				endIndex := to - total
				n = append(n, s.SubStr(0, endIndex))
				return n
			}
		} else { //not writing yet
			if total+l >= from { // can write already
				startIndex := from - total
				if total+l < to { // can fully write
					n = append(n, s.SubStr(startIndex, s.Len()))
				} else { // finish writing
					endIndex := to - total
					n = append(n, s.SubStr(startIndex, endIndex))
					return n
				}
			}
		}
		total += l
	}
	if total < from {
		panic("from is out of bounds")
	}
	panic("to is out of bounds")
}

// Adds the padding to the end of the text.
func (t FormattedText) PadRight(to int, padding FormattedRune) FormattedText {
	if to < 0 {
		panic("from is out of bounds")
	}
	if padding.Len() != 1 {
		panic("padding must be 1 character")
	}
	if to <= t.Len() {
		return t
	}
	return t.Concat(padding.Repeat(to - t.Len()).ToText())
}

// Splits the [FormattedText] into [FormattedText]s by the given separator.
// Splitting happens only inside a [FormattedString] and not across them.
func (t FormattedText) SplitSeq(sep string) iter.Seq[FormattedText] {
	return func(yield func(FormattedText) bool) {
		n := FormattedText{}
		for _, s := range t {
			parts := slices.Collect(s.SplitSeq(sep))
			n = append(n, parts[0])
			for _, line := range parts[1:] {
				if !yield(n) {
					return
				}
				n = FormattedText{line}
			}
		}
		if len(n) > 0 {
			yield(n)
		}
	}
}

// Concatenates [FormattedText]s.
func (t FormattedText) Concat(o FormattedText) FormattedText {
	return append(t, o...)
}

// Gets the length.
func (t FormattedText) Len() int {
	l := 0
	for _, s := range t {
		l += s.Len()
	}
	return l
}

// Returns a string without ANSI formatting.
func (t FormattedText) Unformat() string {
	var sb strings.Builder
	for _, s := range t {
		sb.WriteString(s.Unformat())
	}
	return sb.String()
}

// Implements [io.WriterTo], so result can be printed to the terminal.
func (t FormattedText) WriteTo(w io.Writer) (int64, error) {
	total := int64(0)
	for _, s := range t {
		n, err := s.WriteTo(w)
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, nil
}

// Implements [fmt.Stringer], so result can be printed to the terminal.
func (t FormattedText) String() string {
	var sb strings.Builder
	for _, s := range t {
		sb.WriteString(s.String())
	}
	return sb.String()
}

// Returns if the two [FormattedText]s are equal.
func FormattedTextEquals(a FormattedText, b FormattedText) bool {
	return slices.EqualFunc(a, b, FormattedStringEquals)
}
