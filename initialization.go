package eago

import (
	"math/rand"
	"time"
)

func InitFloatVector(n int, upperLimit, lowerLimit float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	vector := make([]float64, n)
	for i := range vector {
		vector[i] = lowerLimit + rand.Float64()*(upperLimit-lowerLimit)
	}
	return vector
}
