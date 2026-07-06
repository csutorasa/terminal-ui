package document

import (
	"slices"
	"sync"
)

// Func to check if two values are the same.
type EqualsFunc[T any] = func(T, T) bool

// Func to handle changes in value.
type OnPropertyChangeFunc[T any] = func(old T, new T)

// Func to handle changes in slice values.
type OnPropertyItemChangeFunc[T any] = func(T)

// State for storing a property.
type Property[T any] struct {
	change   bool
	value    T
	equals   EqualsFunc[T]
	onChange OnPropertyChangeFunc[T]
	mutex    sync.RWMutex
}

// Creates a new comparable property.
func NewProperty[T comparable](initial T) *Property[T] {
	return NewPropertyFuncOnChange(initial, func(a T, b T) bool {
		return a == b
	}, nil)
}

// Creates a new property with custom equals function.
func NewPropertyFunc[T any](initial T, equals EqualsFunc[T]) *Property[T] {
	return NewPropertyFuncOnChange(initial, equals, nil)
}

// Creates a new property with custom on change handler.
func NewPropertyOnChange[T comparable](initial T, onChange OnPropertyChangeFunc[T]) *Property[T] {
	return NewPropertyFuncOnChange(initial, func(a T, b T) bool {
		return a == b
	}, onChange)
}

// Creates a new property with custom equals function and on change handler.
func NewPropertyFuncOnChange[T any](initial T, equals EqualsFunc[T], onChange OnPropertyChangeFunc[T]) *Property[T] {
	return &Property[T]{
		change:   true,
		value:    initial,
		equals:   equals,
		onChange: onChange,
	}
}

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

// Gets the current value of the property.
func (p *Property[T]) Value() T {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.value
}

// Sets the current value of the property.
func (p *Property[T]) Set(v T) *Property[T] {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.equals(p.value, v) {
		return p
	}
	if p.onChange != nil {
		p.onChange(p.value, v)
	}
	p.value = v
	p.change = true
	return p
}

// Modifies the the current value of the property. Passes the current value.
func (p *Property[T]) Map(f func(oldVal T) T) *Property[T] {
	p.Set(f(p.value))
	return p
}

func (p *Property[T]) Changed() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.change
}

func (p *Property[T]) resetChanged() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.change = false
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
func (p *UniqueSliceProperty[T]) Value() []T {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.value
}

// Sets the current value of the property.
func (p *UniqueSliceProperty[T]) Set(values []T) *UniqueSliceProperty[T] {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range p.value {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if slices.ContainsFunc(values, containsFunc) {
			continue
		}
		if p.onRemove != nil {
			p.onRemove(v)
		}
		p.change = true
	}
	for _, v := range values {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if slices.ContainsFunc(p.value, containsFunc) {
			continue
		}
		if p.onAdd != nil {
			p.onAdd(v)
		}
		p.change = true
	}
	if p.change {
		p.value = values
	}
	return p
}

// Adds the values to the slice.
func (p *UniqueSliceProperty[T]) Add(values ...T) *UniqueSliceProperty[T] {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if slices.ContainsFunc(p.value, containsFunc) {
			continue
		}
		if p.onAdd != nil {
			p.onAdd(v)
		}
		p.value = append(p.value, v)
		p.change = true
	}
	return p
}

// Removes the values from the slice.
func (p *UniqueSliceProperty[T]) Remove(values ...T) *UniqueSliceProperty[T] {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if !slices.ContainsFunc(p.value, containsFunc) {
			continue
		}
		if p.onRemove != nil {
			p.onRemove(v)
		}
		p.value = slices.DeleteFunc(p.value, containsFunc)
		p.change = true
	}
	return p
}

// Removes all values from the slice.
func (p *UniqueSliceProperty[T]) RemoveAll() *UniqueSliceProperty[T] {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.value) == 0 {
		return p
	}
	for _, v := range p.value {
		if p.onRemove != nil {
			p.onRemove(v)
		}
	}
	p.value = []T{}
	p.change = true
	return p
}

func (p *UniqueSliceProperty[T]) Changed() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.change
}

func (p *UniqueSliceProperty[T]) resetChanged() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.change = false
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
