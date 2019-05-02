package eago

import (
	"time"
	"math/rand"
)

func AddNormalFloat(genome []float64, rate float64) {
	rand.Seed(time.Now().UnixNano())
	for i := range genome {
		if rand.Float64() < rate {
			genome[i] += rand.NormFloat64() * genome[i]
		}
	}
}
