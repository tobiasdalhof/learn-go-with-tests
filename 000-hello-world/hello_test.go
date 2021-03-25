package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("test")
	want := "Hello, test"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
