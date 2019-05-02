package main

import (
	"fmt"
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

}

func (V Variables) Crossover() eago.Genome {
	var v Variables
	return v
}

func (V Variables) Clone() eago.Genome {
	var v Variables
	return v
}

func main() {
	fmt.Println("Hello!")
	ga := eago.NewGA()

	var v Variables
	ga.Minimize(v)
}
