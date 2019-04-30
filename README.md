# eago

Evolutionary Algorithm implemented in Go

## Usage
### Installation
```
$ go get github.com/tsurubee/eago
```

### Example
```go
package main

import (
	"github.com/tsurubee/eago"
	"log"
)

func objectiveFunc(x []float64) float64 {
	return x[0] * x[0] + x[1] * x[1]
}

func main() {
	pso := eago.NewDefaultPSO()
	pso.NParticle =  5
	pso.NStep = 20
	pso.Min = -20
	pso.Max = 10

	if err := pso.Minimize(objectiveFunc, 2); err != nil {
		log.Fatal(err)
	}
}
```
If the above sample code is named as `main.go` and executed as follows, the result of parameter search is displayed on the standard output.    
```bash
$ go run main.go
Step   0: Fitness=8.124 Position=[-2.168 1.851]
Step   1: Fitness=8.124 Position=[-2.168 1.851]
Step   2: Fitness=7.394 Position=[-1.618 2.186]
Step   3: Fitness=6.958 Position=[1.068 2.412]
Step   4: Fitness=6.958 Position=[1.068 2.412]
Step   5: Fitness=4.809 Position=[-2.106 0.611]
Step   6: Fitness=2.215 Position=[-1.052 1.053]
Step   7: Fitness=1.535 Position=[0.580 1.095]
Step   8: Fitness=0.163 Position=[0.023 0.403]
Step   9: Fitness=0.163 Position=[0.023 0.403]
Step  10: Fitness=0.055 Position=[-0.175 0.156]
Step  11: Fitness=0.055 Position=[-0.175 0.156]
Step  12: Fitness=0.055 Position=[-0.175 0.156]
Step  13: Fitness=0.055 Position=[-0.175 0.156]
Step  14: Fitness=0.055 Position=[-0.175 0.156]
Step  15: Fitness=0.055 Position=[-0.175 0.156]
Step  16: Fitness=0.055 Position=[-0.180 0.151]
Step  17: Fitness=0.044 Position=[0.019 0.210]
Step  18: Fitness=0.013 Position=[-0.111 0.023]
Step  19: Fitness=0.008 Position=[-0.083 -0.031]
```

## Implemented algorithms

- Genetic algorithm

- Particle swarm optimization

## License

[MIT](https://github.com/tsurubee/eago/blob/master/LICENSE)

## Author

[tsurubee](https://github.com/tsurubee)