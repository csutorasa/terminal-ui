package ansi

import (
	"io"
	"slices"
	"unicode/utf8"
)

var bannedRunes = map[rune]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
	8: true,
	9: true,
	// 10 is \n
	11: true,
	12: true,
	// 13 is \r
	14: true,
	15: true,
	16: true,
	17: true,
	18: true,
	19: true,
	20: true,
	21: true,
	22: true,
	23: true,
	24: true,
	25: true,
	26: true,
	27: true,
	28: true,
	29: true,
	30: true,
	31: true,
}

func ValidateRune(r rune) {
	if !utf8.ValidRune(r) {
		panic(&ErrInvalidRune{r: r})
	}
	if bannedRunes[r] {
		panic(&ErrInvalidRune{r: r})
	}
}

// Immutable ANSI formatted rune
type FormattedRune struct {
	r      rune
	format *Format
}

func NewSimpleRune(r rune) FormattedRune {
	return NewFormattedRune(r, new(Format))
}

func NewFormattedRune(r rune, format *Format) FormattedRune {
	if format == nil {
		panic("format is nil")
	}
	ValidateRune(r)
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
