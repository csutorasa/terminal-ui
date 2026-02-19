package application

import (
	"fmt"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
)

type Application struct {
	document     *document.Document
	renderer     Renderer
	eventDecoder EventDecoder
	quitting     bool
}

// Creates a new [Application].
func New(d *document.Document) *Application {
	if d == nil {
		d = document.Default
	}
	return &Application{
		document: d,
	}
}

// Gets the [document.Document] of the application.
func (a *Application) Document() *document.Document {
	return a.document
}

// Triggers a graceful quit.
func (a *Application) Quit() {
	a.quitting = true
}

// Triggers an event.
func (a *Application) Trigger(event Event) error {
	if event.quitting() {
		a.quitting = true
		return nil
	}
	key, ok := event.Data().(ansi.Input)
	if ok {
		a.document.Trigger(&document.KeyboardEvent{
			Key: ansi.Key(key),
		})
	}
	err := a.Render()
	if err != nil {
		return err
	}
	return nil
}

// Runs the application. This blocks until an error happens or the application quits.
func (a *Application) Run(eventDecoder EventDecoder, renderer Renderer) error {
	if (a.eventDecoder != nil) || a.renderer != nil {
		return fmt.Errorf("application is already running")
	}
	a.quitting = false
	a.renderer = renderer
	a.eventDecoder = eventDecoder
	defer func() {
		a.renderer = nil
		a.eventDecoder = nil
	}()
	err := a.renderer.Init()
	if err != nil {
		return err
	}
	defer a.renderer.Close()
	events := make(chan Event)
	err = a.eventDecoder.Start(events)
	if err != nil {
		return err
	}
	err = a.Render()
	if err != nil {
		return err
	}
	for event := range events {
		err := a.Trigger(event)
		if err != nil {
			return err
		}
		if a.quitting {
			break
		}
	}
	close(events)
	return nil
}

// Renders the document of the application. It can be called only if the application runs.
func (a *Application) Render() error {
	if a.renderer == nil {
		return fmt.Errorf("application is not running")
	}
	return a.renderer.Render(func(rc document.RenderContext) (document.RenderOutput, bool) {
		return a.document.Render(rc)
	})
}
