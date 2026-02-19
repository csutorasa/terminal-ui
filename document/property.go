package document

import (
	"slices"
	"sync"
)

// Func to check if two values are the same.
type EqualsFunc[T any] = func(T, T) bool

// Func to handle changes in value.
type OnChangeFunc[T any] = func(old T, new T)

// Func to handle changes in slice values.
type OnItemChangeFunc[T any] = func(T)

// State for storing a property.
type Property[T any] struct {
	change   bool
	value    T
	equals   EqualsFunc[T]
	onChange OnChangeFunc[T]
	mutex    sync.RWMutex
}

// Creates a new comparable property.
func NewProperty[T comparable](initial T) *Property[T] {
	return NewPropertyFunc(initial, func(a T, b T) bool {
		return a == b
	})
}

// Creates a new slice property.
func NewSliceProperty[T comparable, S ~[]T](initial S) *Property[S] {
	return NewPropertyFunc(initial, slices.Equal)
}

// Creates a new property with custom equals function.
func NewPropertyFunc[T any](initial T, equals EqualsFunc[T]) *Property[T] {
	return &Property[T]{
		change: true,
		value:  initial,
		equals: equals,
	}
}

// Creates a new property with custom on change handler.
func NewPropertyOnChange[T comparable](initial T, onChange OnChangeFunc[T]) *Property[T] {
	return NewPropertyFuncOnChange(initial, func(a T, b T) bool {
		return a == b
	}, onChange)
}

// Creates a new property with custom equals function and on change handler.
func NewPropertyFuncOnChange[T any](initial T, equals EqualsFunc[T], onChange OnChangeFunc[T]) *Property[T] {
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
func (p *Property[T]) Set(v T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.equals(p.value, v) {
		return
	}
	if p.onChange != nil {
		p.onChange(p.value, v)
	}
	p.value = v
	p.change = true
}

// Modifies the the current value of the property. Passes the current value.
func (p *Property[T]) Map(f func(oldVal T) T) {
	p.Set(f(p.value))
}

func (p *Property[T]) changed() bool {
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
	onAdd    OnItemChangeFunc[T]
	onRemove OnItemChangeFunc[T]
	mutex    sync.RWMutex
}

// Gets the current value of the property.
func (p *UniqueSliceProperty[T]) Value() []T {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.value
}

// Adds the values to the slice.
func (p *UniqueSliceProperty[T]) Add(values ...T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if slices.ContainsFunc(p.value, containsFunc) {
			return
		}
		if p.onAdd != nil {
			p.onAdd(v)
		}
		p.value = append(p.value, v)
		p.change = true
	}
}

// Removes the values from the slice.
func (p *UniqueSliceProperty[T]) Remove(values ...T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range values {
		containsFunc := func(a T) bool {
			return p.equals(a, v)
		}
		if !slices.ContainsFunc(p.value, containsFunc) {
			return
		}
		if p.onRemove != nil {
			p.onRemove(v)
		}
		p.value = slices.DeleteFunc(p.value, containsFunc)
		p.change = true
	}
}

// Removes all values from the slice.
func (p *UniqueSliceProperty[T]) RemoveAll() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.value) == 0 {
		return
	}
	for _, v := range p.value {
		if p.onRemove != nil {
			p.onRemove(v)
		}
	}
	p.value = []T{}
	p.change = true
}

func (p *UniqueSliceProperty[T]) changed() bool {
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
func NewUniqueSlicePropertyOnChange[T comparable](add OnItemChangeFunc[T], remove OnItemChangeFunc[T]) *UniqueSliceProperty[T] {
	return NewUniqueSlicePropertyFuncOnChange(func(a T, b T) bool {
		return a == b
	}, add, remove)
}

// Creates a new unique slice property.
func NewUniqueSlicePropertyFuncOnChange[T any](equals EqualsFunc[T], add OnItemChangeFunc[T], remove OnItemChangeFunc[T]) *UniqueSliceProperty[T] {
	return &UniqueSliceProperty[T]{
		change:   true,
		value:    []T{},
		equals:   equals,
		onAdd:    add,
		onRemove: remove,
	}
}
