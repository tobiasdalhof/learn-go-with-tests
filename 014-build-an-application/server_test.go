package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	store := &StubPlayerStore{
		scores: map[string]int{
			"Jerome": 20,
			"Fil":    10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("returns Jerome's score", func(t *testing.T) {
		request := newGetScoreRequest("Jerome")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Fil's score", func(t *testing.T) {
		request := newGetScoreRequest("Fil")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Tobi")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {

	store := &StubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {

		player := "Fil"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		assertPlayerWinCalls(t, store.winCalls, 1)
		assertPlayerWon(t, store.winCalls[0], player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Jerome", 100},
			{"Fil", 50},
		}

		store := &StubPlayerStore{league: wantedLeague}
		server := NewPlayerServer(store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, "application/json")
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	var league []Player
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong: got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong status code: got %d, want %d", got, want)
	}
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error asserting league: got %v, want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Header().Get("content-type")
	if got != want {
		t.Errorf("wrong content-type header: got %q, want %q", got, want)
	}
}

func assertPlayerWon(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wrong player: got %q, want %q", got, want)
	}
}

func assertPlayerWinCalls(t testing.TB, got []string, want int) {
	t.Helper()
	if len(got) != want {
		t.Errorf("wrong player win calls count: got %d, want %d", len(got), want)
	}
}
