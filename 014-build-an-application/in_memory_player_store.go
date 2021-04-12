package main

import "sync"

type InMemoryPlayerStore struct {
	scores map[string]int
	mu     sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: map[string]int{},
	}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.scores[name]++
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, score := range s.scores {
		league = append(league, Player{name, score})
	}
	return league
}
