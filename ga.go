package eago

import "fmt"

type GA struct {
	GAConfig    *GAConfig
	Population  Population
	Generations uint
	Best        Individual
}

type GAConfig struct {
	PopulationSize uint
	NGenerations   uint
	MutationRate   float64
}

type Individual struct {
	Chromosome Genome
	Fitness    float64
}

type Population []Individual

func NewGA() *GA {
	return &GA{
		GAConfig: &GAConfig{
			PopulationSize: 10,
			NGenerations:   10,
			MutationRate:   0.05,
		},
	}
}

func (ga *GA) initPopulation(g Genome) Population {
	pop := make(Population, ga.GAConfig.PopulationSize)
	for i := range pop {
		pop[i].Chromosome = g.Initialization()
	}
	return pop
}

func (ga *GA) Run(g Genome) {
	pop := ga.initPopulation(g)
	fmt.Println(pop)
}