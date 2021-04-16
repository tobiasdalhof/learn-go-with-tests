package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league sorted", func(t *testing.T) {
		file, cleanDatabase := createTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Fil", 33},
			{"Jerome", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		file, cleanDatabase := createTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		got := store.GetPlayerScore("Fil")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		file, cleanDatabase := createTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		store.RecordWin("Fil")

		got := store.GetPlayerScore("Fil")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		file, cleanDatabase := createTempFile(t, `[
			{"Name": "Jerome", "Wins": 10},
			{"Name": "Fil", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(file)
		assertNoError(t, err)

		store.RecordWin("Tobi")

		got := store.GetPlayerScore("Tobi")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong player score: got %d, want %d", got, want)
	}
}
