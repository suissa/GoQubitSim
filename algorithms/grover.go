package algorithms

import (
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

type Grover struct {
	Oracle     func(*core.Qubit) bool
	NumQubits  int
	Iterations int
}

func NewGrover(oracle func(*core.Qubit) bool, n, iter int) *Grover {
	return &Grover{
		Oracle:     oracle,
		NumQubits:  n,
		Iterations: iter,
	}
}

func (g *Grover) Execute() int {
	// Cria registro quântico
	q := core.NewQubit(make([]float64, 1<<g.NumQubits))
	q.Coefficients[0] = 1

	// Aplica Hadamard
	for i := 0; i < g.NumQubits; i++ {
		gates.ApplyHadamard(q)
	}

	// Amplitude amplification
	for i := 0; i < g.Iterations; i++ {
		// Oracle
		if g.Oracle(q) {
			q.Coefficients[0] *= -1
		}

		// Diffusor
		mean := average(q.Coefficients)
		for j := range q.Coefficients {
			q.Coefficients[j] = 2*mean - q.Coefficients[j]
		}
	}

	return q.Measure()
}

func average(coeff []float64) float64 {
	sum := 0.0
	for _, v := range coeff {
		sum += v
	}
	return sum / float64(len(coeff))
}
