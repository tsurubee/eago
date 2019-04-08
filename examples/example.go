package main

import (
	"github.com/tsurubee/eago"
	"log"
)

func objectiveFunc(x []float64) float64 {
	return x[0] * x[0] + x[1] * x[1]
}

func main() {
	pso := eago.NewDefaultPSO()
	pso.NParticle =  2
	pso.NStep = 3
	pso.Min = -10
	pso.Max = 10

	if err := pso.Minimize(objectiveFunc, 2); err != nil {
		log.Fatal(err)
	}
}
