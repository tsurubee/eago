package eago

import (
	"time"
	"math/rand"
)

func WeightedAverage(i1 []float64, i2 []float64) {
	rand.Seed(time.Now().UnixNano())
	for i := range i1 {
		ratio := rand.Float64()
		i1[i] = ratio*i1[i] + (1-ratio)*i2[i]
		i2[i] = (1-ratio)*i1[i] + ratio*i2[i]
	}
}
