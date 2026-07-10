package document

import (
	"slices"
	"testing"
)

func TestUniqueSliceProperty(t *testing.T) {
	property := NewUniqueSliceProperty[int]([]int{1})
	var _ State = property

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if !slices.Equal(property.Value(), []int{1}) {
		t.Logf("Expected value %v got %v", []int{1}, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Add(2)

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if !slices.Equal(property.Value(), []int{1, 2}) {
		t.Logf("Expected value %v got %v", []int{1, 2}, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Remove(1)

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if !slices.Equal(property.Value(), []int{2}) {
		t.Logf("Expected value %v got %v", []int{2}, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Set([]int{3, 1})

	if !property.Changed() {
		t.Log("Expected a changed property")
		t.Fail()
	}
	if !slices.Equal(property.Value(), []int{3, 1}) {
		t.Logf("Expected value %v got %v", []int{3, 1}, property.Value())
		t.Fail()
	}
	property.resetChanged()
	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}

	property.Add(3)

	if property.Changed() {
		t.Log("Expected an unchanged property")
		t.Fail()
	}
	if !slices.Equal(property.Value(), []int{3, 1}) {
		t.Logf("Expected value %v got %v", []int{3, 1}, property.Value())
		t.Fail()
	}
}
