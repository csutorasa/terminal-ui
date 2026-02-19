package components

import "github.com/csutorasa/terminal-ui/ansi"

type KeyBinding struct {
	bindings map[ansi.Key]bool
}

func NewKeyBinding(keys ...ansi.Key) *KeyBinding {
	b := &KeyBinding{
		bindings: map[ansi.Key]bool{},
	}
	for _, k := range keys {
		b.Add(k)
	}
	return b
}

func (b *KeyBinding) Add(k ansi.Key) *KeyBinding {
	b.bindings[k] = true
	return b
}

func (b *KeyBinding) Remove(k ansi.Key) *KeyBinding {
	delete(b.bindings, k)
	return b
}

func (b *KeyBinding) Matches(k ansi.Key) bool {
	_, ok := b.bindings[k]
	return ok
}
