package eago

import (
	"sort"
)

type Individuals []Individual

type Individual struct {
	Chromosome Genome
	Fitness    float64
}

func (indi Individuals) SortByFitness() {
	var less = func(i, j int) bool { return indi[i].Fitness < indi[j].Fitness }
	sort.Slice(indi, less)
}
