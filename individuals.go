package eago

import (
	"sort"
	"sync"
)

type Individuals []Individual

type Individual struct {
	Chromosome Genome
	Fitness    float64
	Evaluated  bool
}

func (indis Individuals) Evaluate(parallel bool) {
	var wg sync.WaitGroup
	if !parallel {
		for i := range indis {
			indis[i].Evaluate()
		}
	} else {
		for i := range indis {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				indis[i].Evaluate()
			}(i)
		}
		wg.Wait()
	}
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

func (indi *Individual) Evaluate() {
	if !indi.Evaluated {
		indi.Fitness   = indi.Chromosome.Fitness()
		indi.Evaluated = true
	}
}

func (indi Individual) Clone() Individual {
	clone := Individual{
		Chromosome: indi.Chromosome,
		Fitness:    indi.Fitness,
	}
	return clone
}
