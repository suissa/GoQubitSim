package algorithms

import (
	"math"
	"math/rand"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

type Shor struct {
	N         int
	Precision int
}

func NewShor(n, precision int) *Shor {
	return &Shor{N: n, Precision: precision}
}

func (s *Shor) Factor() (int, int) {
	for {
		a := rand.Intn(s.N-2) + 2
		gcd := computeGCD(a, s.N)
		if gcd > 1 {
			return gcd, s.N / gcd
		}

		// Simulação do período quântico
		r := s.findPeriod(a)

		if r%2 == 0 {
			if x := int(math.Pow(float64(a), float64(r/2))) % s.N; x != s.N-1 {
				return computeGCD(x+1, s.N), computeGCD(x-1, s.N)
			}
		}
	}
}

func (s *Shor) findPeriod(a int) int {
	// Implementação simplificada
	q := core.NewQubit([]float64{1, 0})
	gates.ApplyHadamard(q)

	// Oracle mock
	if int(math.Pow(float64(a), float64(q.Measure())))%s.N == 1 {
		return q.Measure()
	}
	return -1
}

func computeGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
