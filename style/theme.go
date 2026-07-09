package style

import (
	"github.com/csutorasa/terminal-ui/ansi"
)

var DefaultTheme = &Theme{
	Default: new(ansi.Format).Foreground(ansi.DefaultForeground()).Background(ansi.DefaultBackground()),
	Input:   new(ansi.Format).Foreground(ansi.DefaultForeground()).BackgroundColor(ansi.FormatColorBlue),
	Cursor:  new(ansi.Format).ForegroundColor(ansi.FormatColorBlack).BackgroundColor(ansi.FormatColorWhite),
	Border:  ansi.NewFormattedRune(' ', new(ansi.Format).BackgroundColor(ansi.FormatColorWhite)),
}

var TestTheme = &Theme{
	Default: new(ansi.Format).ForegroundColor(ansi.FormatColorRed).BrightBackgroundColor(ansi.FormatColorMagenta),
	Input:   new(ansi.Format).BrightForegroundColor(ansi.FormatColorGreen).BackgroundColor(ansi.FormatColorWhite),
	Cursor:  new(ansi.Format).BackgroundColor(ansi.FormatColorGreen),
	Border:  ansi.NewFormattedRune(' ', new(ansi.Format).BackgroundColor(ansi.FormatColorYellow)),
}

// Theme to apply to all components
type Theme struct {
	Default *ansi.Format
	Input   *ansi.Format
	Cursor  *ansi.Format
	Border  ansi.FormattedRune
}

// Create [ansi.FormattedString] with default styling.
func (t *Theme) CreateString(str []rune) ansi.FormattedString {
	return ansi.NewFormattedString(str, t.Default)
}

// Create [ansi.FormattedText] with default styling.
func (t *Theme) CreateText(str []rune) ansi.FormattedText {
	return ansi.NewFormattedText(str, t.Default)
}
