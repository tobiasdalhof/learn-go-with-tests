package poker

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	file, removeFile := CreateTempFile(t, "[]")
	defer removeFile()

	store, err := NewFileSystemPlayerStore(file)
	AssertNoError(t, err)

	server := NewPlayerServer(store)

	player := "Jerome"
	wantedCount := 1000

	wg := sync.WaitGroup{}
	wg.Add(wantedCount)
	for i := 0; i < wantedCount; i++ {
		go func(w *sync.WaitGroup) {
			server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
			w.Done()
		}(&wg)
	}
	wg.Wait()

	t.Run("works with an empty file", func(t *testing.T) {
		file, removeFile := CreateTempFile(t, "")
		defer removeFile()

		_, err := NewFileSystemPlayerStore(file)

		AssertNoError(t, err)
	})

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), strconv.Itoa(wantedCount))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetLeagueRequest())

		wantedLeague := League{{Name: player, Wins: wantedCount}}
		got := GetLeagueFromResponse(t, response.Body)

		AssertLeague(t, got, wantedLeague)
	})
}
