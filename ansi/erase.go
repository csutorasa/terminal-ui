package ansi

type EraseScreen string

const (
	EraseScreenFromCursorToEnd       EraseScreen = "0"
	EraseScreenFromBeginningToCursor EraseScreen = "1"
	EraseScreenEntire                EraseScreen = "2"
	EraseScreenSavedLines            EraseScreen = "3"
)

func EraseScreenSequence(code EraseScreen) EscapeSequence {
	return NewEscapeSequence(CommandEraseScreen, string(code))
}

type EraseLine string

const (
	EraseLineFromCursorToEnd       EraseLine = "0"
	EraseLineFromBeginningToCursor EraseLine = "1"
	EraseLineEntire                EraseLine = "2"
)

func EraseLineSequence(code EraseLine) EscapeSequence {
	return NewEscapeSequence(CommandEraseLine, string(code))
}
