package eago

import (
	"fmt"
	"math/rand"
	"time"
)

type PSO struct {
	NParticle     int
	NStep         int
	NDim          int
	Min           float64
	Max           float64
	Coefficients  []float64
	objectiveFunc func([]float64) float64
	GlobalBest    float64
	GlobalBestPos []float64
	swarm         swarm
	PrintCallBack func()
}

type particle struct {
	position        []float64
	velocity        []float64
	fitness         float64
	personalBest    float64
	personalBestPos []float64
}

type swarm []*particle

func NewDefaultPSO() *PSO {
	return &PSO{
		Coefficients: []float64{0.8, 0.8, 0.8},
	}
}

func(pso *PSO) initSwarm() {
	var min float64
	pos := make([]float64, pso.NDim)
	swarm := make(swarm, pso.NParticle)
	for i := range swarm {
		swarm[i] = pso.initParticle()
		if i == 0 {
			min = swarm[i].fitness
			pos = swarm[i].position
		} else {
			if swarm[i].fitness < min {
				min = swarm[i].fitness
				pos = swarm[i].position
			}
		}
	}
	pso.swarm = swarm
	pso.GlobalBestPos = append([]float64{}, pos...)
	pso.GlobalBest = min
}

func(pso *PSO) initParticle() *particle {
	rand.Seed(time.Now().UnixNano())
	pos := InitFloatVector(pso.NDim, pso.Max, pso.Min)
	fitness := pso.objectiveFunc(pos)
	return &particle{
		position:        pos,
		velocity:        []float64{rand.Float64(), rand.Float64()},
		fitness:         fitness,
		personalBest:    fitness,
		personalBestPos: pos,
	}
}

func(pso *PSO) move() {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Float64()
	r2 := rand.Float64()
	c  := pso.Coefficients

	for _, p := range pso.swarm {
		for i := range p.position {
			p.velocity[i] = c[0]*p.velocity[i] + c[1]*(pso.GlobalBestPos[i]-p.position[i])*r1 + c[2]*(p.personalBestPos[i]-p.position[i])*r2
			p.position[i] += p.velocity[i]
		}
		p.fitness = pso.objectiveFunc(p.position)

		// Update personal best
		if p.fitness < p.personalBest {
			p.personalBest    = p.fitness
			p.personalBestPos = append([]float64{}, p.position...)
		}
		// Update global best
		if p.fitness < pso.GlobalBest {
			pso.GlobalBest    = p.fitness
			pso.GlobalBestPos = append([]float64{}, p.position...)
		}
	}
}

func (pso *PSO) Minimize(f func([]float64) float64, nDim int) error {
	pso.objectiveFunc = f
	pso.NDim = nDim
	pso.initSwarm()

	for i := 0; i < pso.NStep; i++ {
		if pso.PrintCallBack != nil {
			pso.PrintCallBack()
		} else {
			fmt.Printf("Step %3d: Fitness=%.3f Position=%.3f\n", i, pso.GlobalBest, pso.GlobalBestPos)
		}
		pso.move()
	}

	return nil
}
