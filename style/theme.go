package style

import (
	"github.com/csutorasa/terminal-ui/ansi"
)

var DefaultTheme = &Theme{
	DefaultForeground: ansi.FormatCodeDefaultForeground,
	DefaultBackground: ansi.FormatCodeDefaultBackground,
	InputForeground:   ansi.FormatCodeDefaultForeground,
	InputBackground:   ansi.FormatCodeBlueBackground,
	Cursor:            []ansi.FormatCode{ansi.FormatCodeWhiteBackground, ansi.FormatCodeBlackForeground},
	Border:            ansi.NewFormattedRune(' ', ansi.FormatCodeWhiteBackground),
}

var TestTheme = &Theme{
	DefaultForeground: ansi.FormatCodeRedForeground,
	DefaultBackground: ansi.FormatCodeBrightMagentaBackground,
	InputForeground:   ansi.FormatCodeBrightGreenForeground,
	InputBackground:   ansi.FormatCodeWhiteBackground,
	Cursor:            []ansi.FormatCode{ansi.FormatCodeGreenBackground},
	Border:            ansi.NewFormattedRune(' ', ansi.FormatCodeYellowBackground),
}

// Theme to apply to all components
type Theme struct {
	DefaultForeground ansi.FormatCode
	DefaultBackground ansi.FormatCode
	InputForeground   ansi.FormatCode
	InputBackground   ansi.FormatCode
	Cursor            []ansi.FormatCode
	Border            ansi.FormattedRune
}

// Create [ansi.FormattedString] with default styling.
func (t *Theme) CreateString(str []rune) ansi.FormattedString {
	return ansi.NewFormattedString(str, t.DefaultForeground, t.DefaultBackground)
}

// Create [ansi.FormattedText] with default styling.
func (t *Theme) CreateText(str []rune) ansi.FormattedText {
	return ansi.NewFormattedText(str, t.DefaultForeground, t.DefaultBackground)
}
