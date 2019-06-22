package eago

import (
	"sort"
	"github.com/golang/sync/errgroup"
)

type Individuals []Individual

type Individual struct {
	Chromosome Genome
	Fitness    float64
	Evaluated  bool
}

func (indis Individuals) Evaluate(parallel bool) {
	if !parallel {
		for i := range indis {
			indis[i].Evaluate()
		}
	} else {
		var g errgroup.Group
		for i := range indis {
			g.Go(func() error {
				indis[i].Evaluate()
				return nil
			})
		}
		g.Wait()
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
