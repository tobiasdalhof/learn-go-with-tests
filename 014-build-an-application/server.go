package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{}
	p.store = store

	router := http.NewServeMux()
	router.HandleFunc("/players/", p.playersHandler)
	router.HandleFunc("/leagues/", p.leaguesHandler)
	p.Handler = router

	return p
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		s.processWin(w, player)
	case http.MethodGet:
		s.showScore(w, player)
	}
}

func (s *PlayerServer) leaguesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, player string) {
	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
