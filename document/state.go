package document

import (
	"sync"
)

// Interface for storing state and tracking changes.
type State interface {
	// Gets if the stored state has changed.
	Changed() bool
	resetChanged()
}

// Combines multiple [State]s into one.
type MultiState struct {
	states []State
	mutex  sync.RWMutex
}

// Creates a new [MultiState].
func NewMultiState() *MultiState {
	return &MultiState{
		states: []State{},
	}
}

// Adds a new [State].
func (ms *MultiState) Append(s State) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.states = append(ms.states, s)
}

// Implements [State].
func (ms *MultiState) Changed() bool {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	for _, c := range ms.states {
		if c.Changed() {
			return true
		}
	}
	return false
}

// Implements [State].
func (ms *MultiState) resetChanged() {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	for _, c := range ms.states {
		c.resetChanged()
	}
}

// Appends a new [State] to the [MultiState] and returns the state.
func appendState[S State](ms *MultiState, s S) S {
	ms.Append(s)
	return s
}
