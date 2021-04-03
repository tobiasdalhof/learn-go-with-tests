package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tobi")

	got := buffer.String()
	want := "Hello, Tobi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
