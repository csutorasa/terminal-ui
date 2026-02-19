package ansi

// ANSI escape sequence command
type Command string

const (
	CommandCursorMoveUp                  Command = "A"
	CommandCursorMoveDown                Command = "B"
	CommandCursorMoveRight               Command = "C"
	CommandCursorMoveLeft                Command = "D"
	CommandCursorMoveBeginningOfLineDown Command = "E"
	CommandCursorMoveBeginningOfLineUp   Command = "F"
	CommandCursorMoveToColumn            Command = "G"
	CommandCursorMove                    Command = "H"
	CommandEraseScreen                   Command = "J"
	CommandEraseLine                     Command = "K"
	CommandEnablePrivateMode             Command = "h"
	CommandDisablePrivateMode            Command = "l"
	CommandFormat                        Command = "m"
	CommandCursorShape                   Command = "q"
)
