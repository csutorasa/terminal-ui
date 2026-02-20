package document

import (
	"sync"
)

// Interface for storing state and tracking changes.
type State interface {
	Changed() bool
	resetChanged()
}

// Combines multiple states into one.
type MultiState struct {
	states []State
	mutex  sync.RWMutex
}

// Creates a new multi state.
func NewMultiState() *MultiState {
	return &MultiState{
		states: []State{},
	}
}

// Adds a new state.
func (ms *MultiState) Append(s State) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.states = append(ms.states, s)
}

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

func (ms *MultiState) resetChanged() {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	for _, c := range ms.states {
		c.resetChanged()
	}
}

// Appends a new state to the multi state and returns it.
func appendState[S State](ms *MultiState, s S) S {
	ms.Append(s)
	return s
}
