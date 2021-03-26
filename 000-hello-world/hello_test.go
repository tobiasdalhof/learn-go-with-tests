package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("test", "")
		want := "Hello, test"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("test", "German")
		want := "Hallo, test"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("test", "French")
		want := "Bonjour, test"
		assertCorrectMessage(t, got, want)
	})
}
