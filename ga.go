package eago

import "fmt"

type GA struct {
	GAConfig
	//Selector
	Population  Population
	Best        Individual
}

type GAConfig struct {
	PopulationSize uint
	NGenerations   uint
	MutationRate   float64
}

type Population struct {
	Individuals Individuals
	Generations uint
}

func NewGA() *GA {
	return &GA{
		GAConfig: GAConfig{
			PopulationSize: 10,
			NGenerations:   10,
			MutationRate:   0.05,
		},
	}
}

func (ga *GA) initPopulation(g Genome) {
	indi := make(Individuals, ga.PopulationSize)
	for i := range indi {
		indi[i].Chromosome = g.Initialization()
		indi[i].Fitness    = indi[i].Chromosome.Fitness()
	}
	ga.Population.Generations = 0
	ga.Population.Individuals = indi
	ga.Population.Individuals.SortByFitness()
}

func (ga *GA) Minimize(g Genome) {
	ga.initPopulation(g)
	fmt.Println(ga.Population)

	for i := uint(0); i < ga.NGenerations; i++ {
		//evolve
	}
}