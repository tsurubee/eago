package main

import (
	"github.com/tsurubee/eago"
	"log"
	"testing"
)

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
	var v Variables
	ga := eago.NewGA(eago.GAConfig{
		PopulationSize: 30,
		NGenerations:   20,
		CrossoverRate:  0.8,
		MutationRate:   0.01,
		ParallelEval:   parallel,
	})
	if err := ga.Minimize(v); err != nil {
		log.Fatal(err)
	}
}
