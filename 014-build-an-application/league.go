package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(reader io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("error decoding league: %v", err)
	}
	return league, nil
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
