package eago

import (
	"fmt"
)

type PSO struct {
	NParticle     int
	NStep         int
	NDim          int
	Min           float64
	Max           float64
	objectiveFunc func([]float64) float64
	GlobalBest    float64
	swarm         []Particle
	PrintCallBack func() string
}

type Particle struct {
	Position     []float64
	Velocity     []float64
	Fitness      float64
	PersonalBest BestPosition
}

type BestPosition struct {
	Position []float64
	Fitness  float64
}

func(pso *PSO) initSwarm() {
	var min float64
	swarm := make([]Particle, pso.NParticle)
	for i := range swarm {
		swarm[i] = pso.initParticle()
		if i == 0 {
			min = swarm[i].Fitness
		} else {
			if swarm[i].Fitness < min {
				min = swarm[i].Fitness
			}
		}
	}
	pso.swarm = swarm
	pso.GlobalBest = min
}

func(pso *PSO) initParticle() Particle {
	pos := InitFloatVector(pso.NDim, pso.Max, pso.Min)
	fitness := pso.objectiveFunc(pos)
	return Particle{
		Position: pos,
		Fitness: fitness,
		PersonalBest: BestPosition{
			Position: pos,
			Fitness: fitness,
		},
	}
}

func(pso *PSO) updateGlobalBest(swarm []Particle) {
}

func (pso *PSO) Minimize(f func([]float64) float64, nDim int) error {
	pso.objectiveFunc = f
	pso.NDim = nDim
	pso.initSwarm()

	fmt.Printf("%v\n", pso.swarm)
	fmt.Printf("%v\n", pso.GlobalBest)

	for i := 0; i < pso.NStep; i++ {
		fmt.Printf("Step %v\n", i+1)
	}

	return nil
}

