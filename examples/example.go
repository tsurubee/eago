package main

import (
	"github.com/tsurubee/eago"
	"log"
)

type Variables []float64

func objectiveFunc(x []float64) float64 {
	return x[0] * x[0] * x[1] * x[1]
}

func main() {
	pso := eago.PSO{
		NParticle: 3,
		NStep: 10,
		Min: -5,
		Max: 5,
	}

	if err := pso.Minimize(objectiveFunc, 2); err != nil {
		log.Fatal(err)
	}
}
