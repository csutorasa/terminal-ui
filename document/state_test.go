package document

import (
	"testing"
)

func TestMultiState(t *testing.T) {
	multiState := NewMultiState()
	var _ State = multiState

	if multiState.Changed() {
		t.Log("Expected an uchanged property")
		t.Fail()
	}

	state := &state{
		changed: true,
	}

	multiState.Append(state)

	if !multiState.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	multiState.resetChanged()
	if multiState.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}
	if state.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}
}

type state struct {
	changed bool
}

func (s *state) Changed() bool {
	return s.changed
}

func (s *state) resetChanged() {
	s.changed = false
}
