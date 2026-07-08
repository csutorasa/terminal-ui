package document

import (
	"testing"
)

func TestProperty(t *testing.T) {
	property := NewProperty(1)
	var _ State = property

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value() != 1 {
		t.Logf("Expected value %d got %d", 1, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Set(2)

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value() != 2 {
		t.Logf("Expected value %d got %d", 2, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Map(func(oldVal int) int {
		if oldVal != 2 {
			t.Logf("Expected value %d", 2)
			t.Fail()
		}
		return 3
	})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value() != 3 {
		t.Logf("Expected value %d got %d", 3, property.Value())
		t.Fail()
	}
}

func TestPropertyFunc(t *testing.T) {
	property := NewPropertyFunc(test{i: 1}, testEquals)
	var _ State = property

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value().i != 1 {
		t.Logf("Expected value %d got %d", 1, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Set(test{i: 2})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value().i != 2 {
		t.Logf("Expected value %d got %d", 2, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Map(func(oldVal test) test {
		if oldVal.i != 2 {
			t.Logf("Expected value %d got %d", 2, property.Value())
			t.Fail()
		}
		return test{i: 3}
	})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if property.Value().i != 3 {
		t.Logf("Expected value %d got %d", 3, property.Value())
		t.Fail()
	}
}

func TestPropertyOnChange(t *testing.T) {
	callCount := 0
	property := NewPropertyOnChange(1, func(old int, new int) {
		callCount++
	})
	var _ State = property

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 0 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value() != 1 {
		t.Logf("Expected value %d got %d", 1, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Set(2)

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 1 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value() != 2 {
		t.Logf("Expected value %d got %d", 2, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Map(func(oldVal int) int {
		if oldVal != 2 {
			t.Logf("Expected value %d got %d", 2, property.Value())
			t.Fail()
		}
		return 3
	})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 2 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value() != 3 {
		t.Logf("Expected value %d got %d", 3, property.Value())
		t.Fail()
	}
}

func TestPropertyFuncOnChange(t *testing.T) {
	callCount := 0
	property := NewPropertyFuncOnChange(test{i: 1}, testEquals, func(old test, new test) {
		callCount++
	})
	var _ State = property

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 0 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value().i != 1 {
		t.Logf("Expected value %d got %d", 1, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Set(test{i: 2})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 1 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value().i != 2 {
		t.Logf("Expected value %d got %d", 2, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Map(func(oldVal test) test {
		if oldVal.i != 2 {
			t.Logf("Expected value %d got %d", 2, property.Value())
			t.Fail()
		}
		return test{i: 3}
	})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if callCount != 2 {
		t.Log("Expected a func call")
		t.Fail()
	}
	if property.Value().i != 3 {
		t.Logf("Expected value %d got %d", 3, property.Value())
		t.Fail()
	}
}

func testEquals(a test, b test) bool {
	return a.i == b.i
}

type test struct {
	i int
}
