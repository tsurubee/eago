package main

import (
	"github.com/tsurubee/eago"
)

type Variables []float64

func (V Variables) Initialization() eago.Genome {
	return Variables(eago.InitFloatVector(2, 20, -20))
}

func (V Variables) Fitness() float64 {
	return V[0] * V[0] + V[1] * V[1]
}

func (V Variables) Mutation() {
	eago.AddNormalFloat(V, 0.5)
}

func (V Variables) Crossover(X eago.Genome) eago.Genome {
	return Variables(eago.BLXalpha(V, X.(Variables), 0.3))
}

func main() {
	ga := eago.NewGA()
	var v Variables
	ga.Minimize(v)
}
