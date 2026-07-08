package document_test

import (
	"testing"

	"github.com/csutorasa/terminal-ui/document"
)

func TestDispatcher(t *testing.T) {
	dispatcher := document.NewDispatcher()
	calledTimes := 0
	dispatcher.Dispatch(func() {
		calledTimes++
	})
	if calledTimes != 0 {
		t.Fatal("No calls are expected yet")
	}
	dispatcher.Dispatch(func() {
		calledTimes++
	})
	if calledTimes != 0 {
		t.Fatal("No calls are expected yet")
	}
	dispatcher.Run()
	if calledTimes != 2 {
		t.Fatal("2 calls are expected")
	}
}
