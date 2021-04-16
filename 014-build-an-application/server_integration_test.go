package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	file, removeFile := createTempFile(t, "[]")
	defer removeFile()

	store, err := NewFileSystemPlayerStore(file)
	assertNoError(t, err)

	server := NewPlayerServer(store)

	player := "Jerome"
	wantedCount := 1000

	wg := sync.WaitGroup{}
	wg.Add(wantedCount)
	for i := 0; i < wantedCount; i++ {
		go func(w *sync.WaitGroup) {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			w.Done()
		}(&wg)
	}
	wg.Wait()

	t.Run("works with an empty file", func(t *testing.T) {
		file, removeFile := createTempFile(t, "")
		defer removeFile()

		_, err := NewFileSystemPlayerStore(file)

		assertNoError(t, err)
	})

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), strconv.Itoa(wantedCount))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())

		wantedLeague := League{{Name: player, Wins: wantedCount}}
		got := getLeagueFromResponse(t, response.Body)

		assertLeague(t, got, wantedLeague)
	})
}
