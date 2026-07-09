package style

type Directional[T any] struct {
	left   T
	top    T
	right  T
	bottom T
}

func NewDirectional[T any](t T) Directional[T] {
	return Directional[T]{
		left:   t,
		top:    t,
		right:  t,
		bottom: t,
	}
}

func NewDirectionalValues[T any](left T, top T, right T, bottom T) Directional[T] {
	return Directional[T]{
		left:   left,
		top:    top,
		right:  right,
		bottom: bottom,
	}
}

func (d Directional[T]) Get() (T, T, T, T) {
	return d.left, d.top, d.right, d.bottom
}

func (d Directional[T]) Left(t T) Directional[T] {
	return NewDirectionalValues(t, d.top, d.right, d.bottom)
}

func (d Directional[T]) Top(t T) Directional[T] {
	return NewDirectionalValues(d.left, t, d.right, d.bottom)
}

func (d *Directional[T]) Right(t T) Directional[T] {
	return NewDirectionalValues(d.left, d.top, t, d.bottom)
}

func (d Directional[T]) Bottom(t T) Directional[T] {
	return NewDirectionalValues(d.left, d.top, d.right, t)
}

func (d Directional[T]) Vertical(t T) Directional[T] {
	return NewDirectionalValues(d.left, t, d.right, t)
}

func (d Directional[T]) Horizontal(t T) Directional[T] {
	return NewDirectionalValues(t, d.top, t, d.bottom)
}

func (d Directional[T]) Values(left T, top T, right T, bottom T) Directional[T] {
	return NewDirectionalValues(left, top, right, bottom)
}

func DirectionalEquals[T comparable](a Directional[T], b Directional[T]) bool {
	return a == b
}

func DirectionalEqualsFunc[T any](equals func(T, T) bool) func(Directional[T], Directional[T]) bool {
	return func(a Directional[T], b Directional[T]) bool {
		return equals(a.left, b.left) && equals(a.top, b.top) && equals(a.right, b.right) && equals(a.bottom, b.bottom)
	}
}
