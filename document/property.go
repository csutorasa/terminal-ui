package document

import (
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
