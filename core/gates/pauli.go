package gates

import (
	"github.com/suissa/goqubitsim/core"
)

type PauliGate struct {
	X [][]float64
	Y [][]float64
	Z [][]float64
}

func NewPauliGates() *PauliGate {
	return &PauliGate{
		X: [][]float64{{0, 1}, {1, 0}},
		Y: [][]float64{{0, -1}, {1, 0}},
		Z: [][]float64{{1, 0}, {0, -1}},
	}
}

func (pg *PauliGate) ApplyX(q *core.Qubit) {
	q.Coefficients = []float64{
		pg.X[0][0]*q.Coefficients[0] + pg.X[0][1]*q.Coefficients[1],
		pg.X[1][0]*q.Coefficients[0] + pg.X[1][1]*q.Coefficients[1],
	}
}

func (pg *PauliGate) ApplyY(q *core.Qubit) {
	q.Coefficients = []float64{
		pg.Y[0][0]*q.Coefficients[0] + pg.Y[0][1]*q.Coefficients[1],
		pg.Y[1][0]*q.Coefficients[0] + pg.Y[1][1]*q.Coefficients[1],
	}
}

func (pg *PauliGate) ApplyZ(q *core.Qubit) {
	q.Coefficients = []float64{
		pg.Z[0][0]*q.Coefficients[0] + pg.Z[0][1]*q.Coefficients[1],
		pg.Z[1][0]*q.Coefficients[0] + pg.Z[1][1]*q.Coefficients[1],
	}
}
