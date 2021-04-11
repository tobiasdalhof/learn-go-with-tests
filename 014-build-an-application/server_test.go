package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
	store := &StubPlayerStore{}
	server := NewPlayerServer(store)

	t.Run("it returns 200 on /leagues", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/leagues/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
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
