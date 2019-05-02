package main

import (
	"github.com/tsurubee/eago"
)

type Variables []float64

func (V Variables) Initialization() eago.Genome {
	return Variables(eago.InitFloatVector(2, 10, -10))
}

func (V Variables) Fitness() float64 {
	return V[0] * V[0] + V[1] * V[1]
}

func (V Variables) Mutation() {
	eago.MutateNormalFloat(V, 0.8)
}

func (V Variables) Crossover(X eago.Genome) {
	eago.WeightedAverage(V, X.(Variables))
}

func main() {
	ga := eago.NewGA()
	var v Variables
	ga.Minimize(v)
}
