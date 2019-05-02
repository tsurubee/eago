package eago

import (
	"math/rand"
)

func MutateNormalFloat(genome []float64, rate float64) {
	for i := range genome {
		if rand.Float64() < rate {
			genome[i] += rand.NormFloat64() * genome[i]
		}
	}
}
