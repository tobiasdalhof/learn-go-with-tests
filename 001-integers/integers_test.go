package integers

import "testing"

func TestAdd(t *testing.T) {
	got := Add(10, 10)
	want := 20
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
