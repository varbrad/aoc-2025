package main

import (
	"testing"
)

func TestDoThing(t *testing.T) {
	result := DoThing()
	expected := 3

	if result != expected {
		t.Errorf("DoThing() = %d; want %d", result, expected)
	}
}
