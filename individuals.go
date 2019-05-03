package eago

import (
	"sort"
)

type Individuals []Individual

type Individual struct {
	Chromosome Genome
	Fitness    float64
}

func (indis Individuals) SortByFitness() {
	var less = func(i, j int) bool { return indis[i].Fitness < indis[j].Fitness }
	sort.Slice(indis, less)
}

func (indis Individuals) Clone() Individuals {
	clone := make(Individuals, len(indis))
	for i := range indis {
		clone[i] = indis[i].Clone()
	}
	return clone
}

func (indi Individual) Clone() Individual {
	clone := Individual{
		Chromosome: indi.Chromosome,
		Fitness:    indi.Fitness,
	}
	return clone
}
