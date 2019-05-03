package eago

import (
	"time"
	"math/rand"
	"math"
)

func BLXalpha(x1 []float64, x2 []float64, alpha float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	child := make([]float64, len(x1))
	for i := range x1 {
		dx := math.Abs(x1[i] - x2[i])
		min := math.Min(x1[i], x2[i]) - alpha*dx
		max := math.Max(x1[i], x2[i]) + alpha*dx
		child[i] = min + rand.Float64()*(max - min)
	}
	return child
}
