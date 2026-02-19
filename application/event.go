package application

// Base Event interface for all events.
type Event interface {
	// Quits the application.
	Quit()
	// Event specific data.
	Data() any

	quitting() bool
}

type defaultEvent struct {
	event any
	quit  bool
}

// Creates a new event.
func NewEvent(e any) Event {
	return &defaultEvent{
		event: e,
	}
}

// Gracefully quits the application.
func (e *defaultEvent) Quit() {
	e.quit = true
}

func (e *defaultEvent) Data() any {
	return e.event
}

func (e *defaultEvent) quitting() bool {
	return e.quit
}
