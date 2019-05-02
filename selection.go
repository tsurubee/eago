package eago

import (
	"errors"
	"time"
	"math/rand"
)

type Selector interface {
	Select(Individuals) (Individuals, error)
}

type Tournament struct {
	NContestants int
}

func (t Tournament) Select(indi Individuals) (Individuals, error) {
	if len(indi) < t.NContestants {
		return nil, errors.New("invalid NSelection: Too large NSelection")
	}
	selected := make(Individuals, len(indi))
	rand.Seed(time.Now().UnixNano())
	for i := range selected {
		winner := indi[rand.Intn(len(indi))]
		for j := 0; j < t.NContestants; j++ {
			next := indi[rand.Intn(len(indi))]
			if winner.Fitness > next.Fitness {
				winner = next
			}
		}
		selected[i] = winner
	}

	return selected, nil
}
