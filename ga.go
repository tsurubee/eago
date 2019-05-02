package eago

import (
	"fmt"
	"log"
	"time"
	"math/rand"
)

type GA struct {
	GAConfig
	Population
	Selector     Selector
	BestSolution Individual
}

type GAConfig struct {
	PopulationSize uint
	NGenerations   uint
	CrossoverRate  float64
	MutationRate   float64
}

type Population struct {
	Individuals Individuals
	Generations uint
}

func NewGA() *GA {
	return &GA{
		GAConfig: GAConfig{
			PopulationSize: 5,
			NGenerations:   5,
			CrossoverRate:  0.8,
			MutationRate:   0.01,
		},
		Selector: Tournament{
			NSelection: 2,
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
	ga.BestSolution = ga.Population.Individuals[0]
}

func (ga *GA) evolve() error {
	ga.Generations++
	rand.Seed(time.Now().UnixNano())
	offSprings := make(Individuals, ga.PopulationSize)
	selected, err:= ga.Selector.Select(ga.Population.Individuals)
	if err != nil {
		log.Fatal(err)
	}
	for i := range offSprings {
		fmt.Println(i)
		if i == len(selected)-1 {
			offSprings[i] = selected[i]
		} else {
			if rand.Float64() < ga.CrossoverRate {
				selected[i].Chromosome.Crossover(selected[i+1].Chromosome)
				offSprings[i] = selected[i]
				offSprings[i].Fitness = offSprings[i].Chromosome.Fitness()
			}
		}
		if rand.Float64() < ga.MutationRate {
			offSprings[i].Chromosome.Mutation()
			offSprings[i].Fitness = offSprings[i].Chromosome.Fitness()
		}
	}

	offSprings.SortByFitness()
	ga.updateBest(offSprings[0])
	ga.Population.Individuals = offSprings
	return nil
}

func (ga *GA) updateBest(indi Individual) {
	if ga.BestSolution.Fitness > indi.Fitness {
		ga.BestSolution = indi
	}
}

func (ga *GA) Minimize(g Genome) error {
	ga.initPopulation(g)
	fmt.Println(ga.Population)

	for i := uint(0); i < ga.NGenerations; i++ {
		if err := ga.evolve(); err != nil {
			return err
		}
		fmt.Printf("Best: %v(%v)", ga.BestSolution.Chromosome, ga.BestSolution.Fitness)
	}

	return nil
}
