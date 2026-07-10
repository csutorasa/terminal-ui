package document

import (
	"slices"
	"sync"
)

// Creates a new comparable property.
func NewSliceProperty[T comparable, S ~[]T](initial S) *Property[[]T] {
	return NewPropertyFuncOnChange[[]T](initial, func(a []T, b []T) bool {
		return slices.Equal(a, b)
	}, nil)
}

// Creates a new property with custom equals function.
func NewSlicePropertyFunc[T any, S ~[]T](initial S, equals EqualsFunc[T]) *Property[[]T] {
	return NewPropertyFuncOnChange[[]T](initial, func(a []T, b []T) bool {
		return slices.EqualFunc(a, b, equals)
	}, nil)
}

// Creates a new property with custom on change handler.
func NewSlicePropertyOnChange[T comparable, S ~[]T](initial S, onChange OnPropertyChangeFunc[[]T]) *Property[[]T] {
	return NewPropertyFuncOnChange[[]T](initial, func(a []T, b []T) bool {
		return slices.Equal(a, b)
	}, onChange)
}

// Creates a new property with custom equals function and on change handler.
func NewSlicePropertyFuncOnChange[T any, S ~[]T](initial S, equals EqualsFunc[T], onChange OnPropertyChangeFunc[[]T]) *Property[[]T] {
	return NewPropertyFuncOnChange[[]T](initial, func(a []T, b []T) bool {
		return slices.EqualFunc(a, b, equals)
	}, onChange)
}

// State for storing a unique slice property.
type UniqueSliceProperty[T any] struct {
	change   bool
	value    []T
	equals   EqualsFunc[T]
	onAdd    OnPropertyItemChangeFunc[T]
	onRemove OnPropertyItemChangeFunc[T]
	mutex    sync.RWMutex
}

// Gets the current value of the property.
func (usp *UniqueSliceProperty[T]) Value() []T {
	usp.mutex.RLock()
	defer usp.mutex.RUnlock()
	return usp.value
}

// Sets the current value of the property.
func (usp *UniqueSliceProperty[T]) Set(values []T) *UniqueSliceProperty[T] {
	usp.mutex.Lock()
	defer usp.mutex.Unlock()
	for _, v := range usp.value {
		containsFunc := func(a T) bool {
			return usp.equals(a, v)
		}
		if slices.ContainsFunc(values, containsFunc) {
			continue
		}
		if usp.onRemove != nil {
			usp.onRemove(v)
		}
		usp.change = true
	}
	for _, v := range values {
		containsFunc := func(a T) bool {
			return usp.equals(a, v)
		}
		if slices.ContainsFunc(usp.value, containsFunc) {
			continue
		}
		if usp.onAdd != nil {
			usp.onAdd(v)
		}
		usp.change = true
	}
	if usp.change {
		usp.value = values
	}
	return usp
}

// Adds the values to the slice.
func (usp *UniqueSliceProperty[T]) Add(values ...T) *UniqueSliceProperty[T] {
	usp.mutex.Lock()
	defer usp.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return usp.equals(a, v)
		}
		if slices.ContainsFunc(usp.value, containsFunc) {
			continue
		}
		if usp.onAdd != nil {
			usp.onAdd(v)
		}
		usp.value = append(usp.value, v)
		usp.change = true
	}
	return usp
}

// Removes the values from the slice.
func (usp *UniqueSliceProperty[T]) Remove(values ...T) *UniqueSliceProperty[T] {
	usp.mutex.Lock()
	defer usp.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return usp.equals(a, v)
		}
		if !slices.ContainsFunc(usp.value, containsFunc) {
			continue
		}
		if usp.onRemove != nil {
			usp.onRemove(v)
		}
		usp.value = slices.DeleteFunc(usp.value, containsFunc)
		usp.change = true
	}
	return usp
}

// Removes all values from the slice.
func (usp *UniqueSliceProperty[T]) RemoveAll() *UniqueSliceProperty[T] {
	usp.mutex.Lock()
	defer usp.mutex.Unlock()
	if len(usp.value) == 0 {
		return usp
	}
	for _, v := range usp.value {
		if usp.onRemove != nil {
			usp.onRemove(v)
		}
	}
	usp.value = []T{}
	usp.change = true
	return usp
}

func (usp *UniqueSliceProperty[T]) Changed() bool {
	usp.mutex.RLock()
	defer usp.mutex.RUnlock()
	return usp.change
}

func (usp *UniqueSliceProperty[T]) resetChanged() {
	usp.mutex.Lock()
	defer usp.mutex.Unlock()
	usp.change = false
}

// Creates a new unique slice property.
func NewUniqueSliceProperty[T comparable, S ~[]T](initial S) *UniqueSliceProperty[T] {
	return NewUniqueSlicePropertyFuncOnChange(initial, func(a T, b T) bool {
		return a == b
	}, nil, nil)
}

// Creates a new unique slice property.
func NewUniqueSlicePropertyFunc[T any, S ~[]T](initial S, equals EqualsFunc[T]) *UniqueSliceProperty[T] {
	return NewUniqueSlicePropertyFuncOnChange(initial, equals, nil, nil)
}

// Creates a new unique slice property.
func NewUniqueSlicePropertyOnChange[T comparable, S ~[]T](initial S, add OnPropertyItemChangeFunc[T], remove OnPropertyItemChangeFunc[T]) *UniqueSliceProperty[T] {
	return NewUniqueSlicePropertyFuncOnChange(initial, func(a T, b T) bool {
		return a == b
	}, add, remove)
}

// Creates a new unique slice property.
func NewUniqueSlicePropertyFuncOnChange[T any, S ~[]T](initial S, equals EqualsFunc[T], add OnPropertyItemChangeFunc[T], remove OnPropertyItemChangeFunc[T]) *UniqueSliceProperty[T] {
	return &UniqueSliceProperty[T]{
		change:   true,
		value:    initial,
		equals:   equals,
		onAdd:    add,
		onRemove: remove,
	}
}
