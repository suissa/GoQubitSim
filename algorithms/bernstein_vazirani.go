package algorithms

import (
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

type BernsteinVazirani struct {
	Oracle func(*core.Qubit) bool
	N      int
}

func NewBernsteinVazirani(oracle func(*core.Qubit) bool, n int) *BernsteinVazirani {
	return &BernsteinVazirani{Oracle: oracle, N: n}
}

func (bv *BernsteinVazirani) FindHiddenString() string {
	hidden := make([]byte, bv.N)

	for i := 0; i < bv.N; i++ {
		q := core.NewQubit([]float64{1, 0})
		gates.ApplyHadamard(q)

		if bv.Oracle(q) {
			hidden[i] = '1'
		} else {
			hidden[i] = '0'
		}
	}

	return string(hidden)
}
