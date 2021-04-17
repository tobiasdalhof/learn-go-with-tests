package poker

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league sorted", func(t *testing.T) {
		file, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		AssertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Fil", 33},
			{"Jerome", 10},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		file, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Fil")
		want := 33

		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		file, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		AssertNoError(t, err)

		store.RecordWin("Fil")

		got := store.GetPlayerScore("Fil")
		want := 34
		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		file, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		AssertNoError(t, err)

		store.RecordWin("Tobi")

		got := store.GetPlayerScore("Tobi")
		want := 1
		AssertScoreEquals(t, got, want)
	})
}
