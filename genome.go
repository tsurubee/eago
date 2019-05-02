package eago

type Genome interface {
	Initialization() Genome
	Fitness() float64
	Mutation()
	Crossover(Genome)
}
