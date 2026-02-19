package ansi_test

import (
	"testing"
)

func assertString(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Logf("expected: %v", []byte(expected))
		t.Logf("     got: %v", []byte(actual))
		t.Fail()
	}
}

func assertLen(t *testing.T, expected int, actual interface{ Len() int }) {
	l := actual.Len()
	if l != expected {
		t.Logf("expected: %d got: %d", expected, l)
		t.Fail()
	}
}

func assertInt(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Fatalf("expected: %d got: %d", expected, actual)
	}
}
