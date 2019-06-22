package main

import (
	"github.com/tsurubee/eago"
	"log"
	"testing"
	"time"
)

type Vector []float64

func (V Vector) Initialization() eago.Genome {
	return Vector(eago.InitFloatVector(2, 32, -32))
}

func (V Vector) Fitness() float64 {
	time.Sleep(1 * time.Millisecond)
	var s float64
	for _, x := range V {
		s += x * x
	}
	return s
}

func (V Vector) Mutation() {
	eago.AddNormalFloat(V, 0.5)
}

func (V Vector) Crossover(X eago.Genome) eago.Genome {
	return Vector(eago.BLXalpha(V, X.(Vector), 0.3))
}

func BenchmarkParallelGA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runGA(true)
	}
}

func BenchmarkNonParallelGA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runGA(false)
	}
}

func runGA(parallel bool) {
	var v Vector
	ga := eago.NewGA(eago.GAConfig{
		PopulationSize: 30,
		NGenerations:   20,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
		ParallelEval:   parallel,
	})
	ga.PrintCallBack = func() {} // Do not print messages while running
	if err := ga.Minimize(v); err != nil {
		log.Fatal(err)
	}
}
