package document

// DOM Element properties.
type ElementProperties struct {
	children      *UniqueSliceProperty[*Element]
	focus         *Property[bool]
	renderContext *Property[RenderContext]
	state         *MultiState
}

func newElementProperties(children *UniqueSliceProperty[*Element]) *ElementProperties {
	state := NewMultiState()
	return &ElementProperties{
		children:      appendState(state, children),
		focus:         appendState(state, NewProperty(false)),
		renderContext: appendState(state, NewProperty(RenderContext{})),
		state:         state,
	}
}

// Gets the property to read and write the children of the element.
func (ep *ElementProperties) Children() *UniqueSliceProperty[*Element] {
	return ep.children
}

// Gets if the element has focus.
func (ep *ElementProperties) Focused() bool {
	return ep.focus.Value()
}

// Implements [State].
func (ep *ElementProperties) Changed() bool {
	return ep.state.Changed()
}

// Implements [State].
func (ep *ElementProperties) resetChanged() {
	ep.state.resetChanged()
}

// Appends a new [State] to the state of the element and returns it.
func AppendState[S State](element *Element, s S) S {
	element.properties.state.Append(s)
	return s
}
