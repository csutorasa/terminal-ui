package ansi

import (
	"slices"
)

type FormatCode = string

const (
	FormatCodeReset                   FormatCode = "0"
	FormatCodeBold                    FormatCode = "1"
	FormatCodeFaint                   FormatCode = "2"
	FormatCodeItalic                  FormatCode = "3"
	FormatCodeUnderline               FormatCode = "4"
	FormatCodeBlinking                FormatCode = "5"
	FormatCodeInverse                 FormatCode = "7"
	FormatCodeHidden                  FormatCode = "8"
	FormatCodeStrikethrough           FormatCode = "9"
	FormatCodeResetBold               FormatCode = "21"
	FormatCodeResetFaint              FormatCode = "22"
	FormatCodeResetItalic             FormatCode = "23"
	FormatCodeResetUnderline          FormatCode = "24"
	FormatCodeResetBlinking           FormatCode = "25"
	FormatCodeResetInverse            FormatCode = "27"
	FormatCodeResetHidden             FormatCode = "28"
	FormatCodeResetStrikethrough      FormatCode = "29"
	FormatCodeBlackForeground         FormatCode = "30"
	FormatCodeRedForeground           FormatCode = "31"
	FormatCodeGreenForeground         FormatCode = "32"
	FormatCodeYellowForeground        FormatCode = "33"
	FormatCodeBlueForeground          FormatCode = "34"
	FormatCodeMagentaForeground       FormatCode = "35"
	FormatCodeCyanForeground          FormatCode = "36"
	FormatCodeWhiteForeground         FormatCode = "37"
	FormatCodeForegroundColor         FormatCode = "38"
	FormatCodeDefaultForeground       FormatCode = "39"
	FormatCodeBlackBackground         FormatCode = "40"
	FormatCodeRedBackground           FormatCode = "41"
	FormatCodeGreenBackground         FormatCode = "42"
	FormatCodeYellowBackground        FormatCode = "43"
	FormatCodeBlueBackground          FormatCode = "44"
	FormatCodeMagentaBackground       FormatCode = "45"
	FormatCodeCyanBackground          FormatCode = "46"
	FormatCodeWhiteBackground         FormatCode = "47"
	FormatCodeBackgroundColor         FormatCode = "48"
	FormatCodeDefaultBackground       FormatCode = "49"
	FormatCodeBrightBlackForeground   FormatCode = "90"
	FormatCodeBrightRedForeground     FormatCode = "91"
	FormatCodeBrightGreenForeground   FormatCode = "92"
	FormatCodeBrightYellowForeground  FormatCode = "93"
	FormatCodeBrightBlueForeground    FormatCode = "94"
	FormatCodeBrightMagentaForeground FormatCode = "95"
	FormatCodeBrightCyanForeground    FormatCode = "96"
	FormatCodeBrightWhiteForeground   FormatCode = "97"
	FormatCodeBrightBlackBackground   FormatCode = "100"
	FormatCodeBrightRedBackground     FormatCode = "101"
	FormatCodeBrightGreenBackground   FormatCode = "102"
	FormatCodeBrightYellowBackground  FormatCode = "103"
	FormatCodeBrightBlueBackground    FormatCode = "104"
	FormatCodeBrightMagentaBackground FormatCode = "105"
	FormatCodeBrightCyanBackground    FormatCode = "106"
	FormatCodeBrightWhiteBackground   FormatCode = "107"
)

func FormatCodesEquals(a []FormatCode, b []FormatCode) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		if !slices.Contains(b, v) {
			return false
		}
	}
	for _, v := range b {
		if !slices.Contains(a, v) {
			return false
		}
	}
	return true
}
