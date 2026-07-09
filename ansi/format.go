package ansi

import (
	"fmt"
	"image/color"
	"io"
)

const (
	FormatCodeReset string = "0"
)

var resetSequence = NewEscapeSequence(CommandFormat, FormatCodeReset)

func FormatResetSequence() EscapeSequence {
	return resetSequence
}

const (
	FormatFlagSetOffset     uint8      = 0
	FormatFlagResetOffset   uint8      = 20
	FormatFlagBold          FormatFlag = 1
	FormatFlagFaint         FormatFlag = 2
	FormatFlagItalic        FormatFlag = 3
	FormatFlagUnderline     FormatFlag = 4
	FormatFlagBlinking      FormatFlag = 5
	FormatFlagInverse       FormatFlag = 7
	FormatFlagHidden        FormatFlag = 8
	FormatFlagStrikethrough FormatFlag = 9
)

type FormatFlag uint8
type FormatFlagCode string

func (f FormatFlag) Set() FormatFlagCode {
	return flagToStr(f + FormatFlag(FormatFlagSetOffset))
}

func (f FormatFlag) Reset() FormatFlagCode {
	return flagToStr(f + FormatFlag(FormatFlagResetOffset))
}

func (f FormatFlag) Value(on bool) FormatFlagCode {
	if on {
		return f.Set()
	}
	return f.Reset()
}

func flagToStr(flag FormatFlag) FormatFlagCode {
	return FormatFlagCode(fmt.Sprintf("%d", flag))
}

const (
	FormatColorForegroundOffset       uint8             = 30
	FormatColorBackgroundOffset       uint8             = 40
	FormatColorBrightForegroundOffset uint8             = 90
	FormatColorBrightBackgroundOffset uint8             = 100
	FormatColorBlack                  FormatColor       = 0
	FormatColorRed                    FormatColor       = 1
	FormatColorGreen                  FormatColor       = 2
	FormatColorYellow                 FormatColor       = 3
	FormatColorBlue                   FormatColor       = 4
	FormatColorMagenta                FormatColor       = 5
	FormatColorCyan                   FormatColor       = 6
	FormatColorWhite                  FormatColor       = 7
	FormatColorTrueColor              uint8             = 8
	FormatColorDefault                uint8             = 9
	FormatRGBTrueColor                FormatColorFormat = 2
)

type FormatColor uint8
type FormatColorFormat uint8
type FormatTrueColorCode string

func FormatColorFromColor(c color.Color) FormatTrueColorCode {
	r, g, b, _ := c.RGBA()
	return FormatColorFromRGB(uint8(uint16(r)/256), uint8(uint16(g)/256), uint8(uint16(b)/256))
}

func FormatColorFromRGB(r uint8, g uint8, b uint8) FormatTrueColorCode {
	return FormatTrueColorCode(fmt.Sprintf("%d;%d;%d;%d", FormatRGBTrueColor, r, g, b))
}

type FormatColorForegroundCode string

var defaultForeground = foregroundColorToStr(FormatColor(FormatColorDefault + FormatColorForegroundOffset))

func DefaultForeground() FormatColorForegroundCode {
	return defaultForeground
}

func (f FormatColor) Foreground() FormatColorForegroundCode {
	return foregroundColorToStr(f + FormatColor(FormatColorForegroundOffset))
}

func (f FormatColor) BrightForeground() FormatColorForegroundCode {
	return foregroundColorToStr(f + FormatColor(FormatColorBrightForegroundOffset))
}

func (tcc FormatTrueColorCode) Foreground() FormatColorForegroundCode {
	return FormatColorForegroundCode(fmt.Sprintf("%d;%s", FormatColorForegroundOffset+uint8(FormatColorTrueColor), tcc))
}

func foregroundColorToStr(color FormatColor) FormatColorForegroundCode {
	return FormatColorForegroundCode(fmt.Sprintf("%d", color))
}

type FormatColorBackgroundCode string

var defaultBackground = backgroundColorToStr(FormatColor(FormatColorDefault + FormatColorBackgroundOffset))

func DefaultBackground() FormatColorBackgroundCode {
	return defaultBackground
}

func (f FormatColor) Background() FormatColorBackgroundCode {
	return backgroundColorToStr(f + FormatColor(FormatColorBackgroundOffset))
}

func (f FormatColor) BrightBackground() FormatColorBackgroundCode {
	return backgroundColorToStr(f + FormatColor(FormatColorBrightBackgroundOffset))
}

func (tcc FormatTrueColorCode) Background() FormatColorBackgroundCode {
	return FormatColorBackgroundCode(fmt.Sprintf("%d;%s", FormatColorBackgroundOffset+uint8(FormatColorTrueColor), tcc))
}

func backgroundColorToStr(color FormatColor) FormatColorBackgroundCode {
	return FormatColorBackgroundCode(fmt.Sprintf("%d", color))
}

type bitMask = uint8

const (
	boldMask          bitMask = 1 << 0
	faintMask         bitMask = 1 << 1
	italicMask        bitMask = 1 << 2
	underlineMask     bitMask = 1 << 3
	blinkingMask      bitMask = 1 << 4
	inverseMask       bitMask = 1 << 5
	hiddenMask        bitMask = 1 << 6
	strikethroughMask bitMask = 1 << 7
)

type Format struct {
	flagsSet   bitMask
	flags      bitMask
	foreground FormatColorForegroundCode
	background FormatColorBackgroundCode
}

func (f *Format) Bold() *Format {
	f.setFlag(boldMask, true)
	return f
}

func (f *Format) ResetBold() *Format {
	f.setFlag(boldMask, false)
	return f
}

func (f *Format) Faint() *Format {
	f.setFlag(faintMask, true)
	return f
}

func (f *Format) ResetFaint() *Format {
	f.setFlag(faintMask, false)
	return f
}

func (f *Format) Italic() *Format {
	f.setFlag(italicMask, true)
	return f
}

func (f *Format) ResetItalic() *Format {
	f.setFlag(italicMask, false)
	return f
}

func (f *Format) Underline() *Format {
	f.setFlag(underlineMask, true)
	return f
}

func (f *Format) ResetUnderline() *Format {
	f.setFlag(underlineMask, false)
	return f
}

func (f *Format) Blinking() *Format {
	f.setFlag(blinkingMask, true)
	return f
}

func (f *Format) ResetBlinking() *Format {
	f.setFlag(blinkingMask, false)
	return f
}

func (f *Format) Inverse() *Format {
	f.setFlag(inverseMask, true)
	return f
}

func (f *Format) ResetInverse() *Format {
	f.setFlag(inverseMask, false)
	return f
}

func (f *Format) Hidden() *Format {
	f.setFlag(hiddenMask, true)
	return f
}

func (f *Format) ResetHidden() *Format {
	f.setFlag(hiddenMask, false)
	return f
}

func (f *Format) Strikethrough() *Format {
	f.setFlag(strikethroughMask, true)
	return f
}

func (f *Format) ResetStrikethrough() *Format {
	f.setFlag(strikethroughMask, false)
	return f
}

func (f *Format) Foreground(cfc FormatColorForegroundCode) *Format {
	f.foreground = cfc
	return f
}

func (f *Format) ForegroundColor(c FormatColor) *Format {
	f.foreground = c.Foreground()
	return f
}

func (f *Format) BrightForegroundColor(c FormatColor) *Format {
	f.foreground = c.BrightForeground()
	return f
}

func (f *Format) Background(cbc FormatColorBackgroundCode) *Format {
	f.background = cbc
	return f
}

func (f *Format) BackgroundColor(c FormatColor) *Format {
	f.background = c.Background()
	return f
}

func (f *Format) BrightBackgroundColor(c FormatColor) *Format {
	f.background = c.BrightBackground()
	return f
}

func (f *Format) EscapeSequence() EscapeSequence {
	arguments := []string{}
	f.append(boldMask, FormatFlagBold, &arguments)
	f.append(faintMask, FormatFlagFaint, &arguments)
	f.append(italicMask, FormatFlagItalic, &arguments)
	f.append(underlineMask, FormatFlagUnderline, &arguments)
	f.append(blinkingMask, FormatFlagBlinking, &arguments)
	f.append(inverseMask, FormatFlagInverse, &arguments)
	f.append(hiddenMask, FormatFlagHidden, &arguments)
	f.append(strikethroughMask, FormatFlagStrikethrough, &arguments)
	if f.foreground != "" {
		arguments = append(arguments, string(f.foreground))
	}
	if f.background != "" {
		arguments = append(arguments, string(f.background))
	}
	return NewEscapeSequence(CommandFormat, arguments...)
}

func (f *Format) FormatString(s string) string {
	if len(s) == 0 {
		return ""
	}
	if !f.hasFormat() {
		return string(s)
	}
	return fmt.Sprintf("%s%s%s", f.EscapeSequence(), string(s), FormatResetSequence())
}

func (f *Format) WriteTo(w io.Writer, s string) (int64, error) {
	if len(s) == 0 {
		return 0, nil
	}
	if !f.hasFormat() {
		n, err := w.Write(([]byte)(s))
		return int64(n), err
	}
	total := int64(0)
	n, err := f.EscapeSequence().WriteTo(w)
	total += n
	if err != nil {
		return total, err
	}
	n1, err := w.Write(([]byte)(s))
	total += int64(n1)
	if err != nil {
		return total, err
	}
	n, err = FormatResetSequence().WriteTo(w)
	total += n
	return total, err
}

func (f *Format) hasFormat() bool {
	return f.flagsSet > 0 || f.foreground != "" || f.background != ""
}

func (f *Format) setFlag(mask uint8, on bool) {
	f.flagsSet |= mask
	if on {
		f.flags |= mask
	} else {
		f.flags &= ^mask
	}
}

func (f *Format) getFlag(mask uint8) (bool, bool) {
	return f.flags&mask > 0, f.flagsSet&mask > 0
}

func (f *Format) append(mask uint8, flag FormatFlag, arguments *[]string) {
	value, set := f.getFlag(mask)
	if !set {
		return
	}
	*arguments = append(*arguments, string(flag.Value(value)))
}

func FormatEquals(a *Format, b *Format) bool {
	if a == b {
		return true
	}
	if a.flagsSet != b.flagsSet {
		return false
	}
	if a.flags&a.flagsSet != b.flags&b.flagsSet {
		return false
	}
	return a.foreground == b.foreground || a.background == b.background
}
