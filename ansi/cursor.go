package ansi

import "fmt"

// Cursor shape options
type CursorShape string

const (
	CursorShapeSteadyBlock       CursorShape = "0"
	CursorShapeSteadyBlockAlso   CursorShape = "1"
	CursorShapeBlinkingBlock     CursorShape = "2"
	CursorShapeSteadyUnderline   CursorShape = "3"
	CursorShapeBlinkingUnderline CursorShape = "4"
	CursorShapeSteadyBar         CursorShape = "5"
	CursorShapeBlinkingBar       CursorShape = "6"
)

func CursorShapeSequence(code EraseScreen) EscapeSequence {
	return NewEscapeSequence(CommandCursorShape, string(code))
}

func MoveUpSequence(line int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveUp, fmt.Sprintf("%d", line))
}

func MoveDownSequence(line int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveDown, fmt.Sprintf("%d", line))
}

func MoveRightSequence(column int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveRight, fmt.Sprintf("%d", column))
}

func MoveLeftSequence(column int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveLeft, fmt.Sprintf("%d", column))
}

func MoveBeginningOfLineDownSequence(line int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveBeginningOfLineDown, fmt.Sprintf("%d", line))
}

func MoveBeginningOfLineUpSequence(line int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveBeginningOfLineUp, fmt.Sprintf("%d", line))
}

func MoveToColumnSequence(colimn int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMoveToColumn, fmt.Sprintf("%d", colimn))
}

func MoveSequence(line int, column int) EscapeSequence {
	return NewEscapeSequence(CommandCursorMove, fmt.Sprintf("%d", line), fmt.Sprintf("%d", column))
}
