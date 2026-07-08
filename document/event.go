package document

// Document event.
type Event struct {
	propagate bool
	event     any
}

// Creates a new [Event].
func NewEvent(e any) *Event {
	return &Event{
		event: e,
	}
}

// Stops the event progagation, so parents will not receive this event anymore.
func (e *Event) StopPropagation() {
	e.propagate = true
}

// Gets the [Event]'s raw data.
func (e *Event) Event() any {
	return e.event
}

func (e *Event) propagating() bool {
	return e.propagate
}
