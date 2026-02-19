package ansi

// Terminal private modes
type PrivateMode string

const (
	PrivateModeBockCursor        PrivateMode = "12"
	PrivateModeCursorVisible     PrivateMode = "25"
	PrivateModeAlternativeBuffer PrivateMode = "1049"
	PrivateModeAlternativeScroll PrivateMode = "1007"
	PrivateModeBracketedPaste    PrivateMode = "2004"
)

func EnablePrivateMode(m PrivateMode) EscapeSequence {
	return NewEscapeSequence(CommandEnablePrivateMode, "?"+string(m))
}

func DisablePrivateMode(m PrivateMode) EscapeSequence {
	return NewEscapeSequence(CommandDisablePrivateMode, "?"+string(m))
}
