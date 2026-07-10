package ansi

import "fmt"

// Error that happens if the rune is invalid.
type ErrInvalidRune struct {
	r rune
}

func (i *ErrInvalidRune) Error() string {
	return fmt.Sprintf("invalid rune %d '%c'", i.r, i.r)
}
