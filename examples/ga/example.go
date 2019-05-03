package main

import (
	"github.com/tsurubee/eago"
	"log"
	"math"
)

type Variables []float64

func (V Variables) Initialization() eago.Genome {
	return Variables(eago.InitFloatVector(2, 32, -32))
}

func (V Variables) Fitness() float64 {
	return Ackley(V)
}

func (V Variables) Mutation() {
	eago.AddNormalFloat(V, 0.5)
}

func (V Variables) Crossover(X eago.Genome) eago.Genome {
	return Variables(eago.BLXalpha(V, X.(Variables), 0.3))
}

func Ackley(X []float64) float64 {
	a, b, c, d := 20.0, 0.2, 2*math.Pi, float64(len(X))
	var s1, s2 float64
	for _, x := range X {
		s1 += x * x
		s2 += math.Cos(c * x)
	}
	return -a*math.Exp(-b*math.Sqrt(s1/d)) - math.Exp(s2/d) + a + math.Exp(1)
}

func main() {
	var v Variables
	ga := eago.NewGA(eago.GAConfig{
		PopulationSize: 30,
		NGenerations:   20,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
	})
	if err := ga.Minimize(v); err != nil {
		log.Fatal(err)
	}
}
