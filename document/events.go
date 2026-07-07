package document

import "github.com/csutorasa/terminal-ui/ansi"

// Triggered when an element receives focus.
type OnFocusEvent struct {
}

// Triggered when an element loses focus.
type OnFocusLostEvent struct {
}

// Triggered when a keyboard input is received.
type KeyboardEvent struct {
	Key ansi.Key
}

// Triggered when a dispatch is requested.
type DispatchEvent struct {
}
