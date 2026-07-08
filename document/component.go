package document

// Base interface for a component.
// This interface is used by the document package to communicate to the custom components.
type Component interface {
	// Handles the event.
	OnEvent(e *Event)
	// Renders the component.
	Render(rw *RenderWriter)
	// Sets the layout of the child elements.
	Layout(lc LayoutContext)
	// Gets if the component can recieve focus.
	Focusable() bool
}
