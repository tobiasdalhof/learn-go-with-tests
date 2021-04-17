package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	store := &StubPlayerStore{
		Scores: map[string]int{
			"Jerome": 20,
			"Fil":    10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("returns Jerome's score", func(t *testing.T) {
		request := NewGetScoreRequest("Jerome")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Fil's score", func(t *testing.T) {
		request := NewGetScoreRequest("Fil")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := NewGetScoreRequest("Tobi")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {

	store := &StubPlayerStore{
		Scores: map[string]int{},
	}
	server := NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {

		player := "Fil"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWinCalls(t, store.WinCalls, 1)
		AssertPlayerWon(t, store.WinCalls[0], player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := League{
			{"Jerome", 100},
			{"Fil", 50},
		}

		store := &StubPlayerStore{League: wantedLeague}
		server := NewPlayerServer(store)

		request := NewGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, "application/json")
	})
}
