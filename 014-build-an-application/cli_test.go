package poker_test

import (
	"strings"
	"testing"

	poker "github.com/tobidalhof/learn-go-with-tests/014-build-an-application"
)

func TestCLI(t *testing.T) {

	t.Run("record Jerome win from user input", func(t *testing.T) {
		in := strings.NewReader("Jerome wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWon(t, playerStore.WinCalls[0], "Jerome")
	})

	t.Run("record Fil win from user input", func(t *testing.T) {
		in := strings.NewReader("Fil wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWon(t, playerStore.WinCalls[0], "Fil")
	})
}
