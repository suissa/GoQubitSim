package core

import (
	"math"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/core/gates"
)

type QuantumBoltzmannMachine struct {
	VisibleQubits []*Qubit
	HiddenQubits  []*Qubit
	Weights       [][]float64 // Pesos entre visible e hidden
	VisibleBias   []float64
	HiddenBias    []float64
	Temperature   float64
}

func NewQuantumBoltzmannMachine(visible, hidden int) *QuantumBoltzmannMachine {
	qbm := &QuantumBoltzmannMachine{
		Temperature: 1.0,
	}

	// Inicializa qubits
	qbm.VisibleQubits = make([]*Qubit, visible)
	qbm.HiddenQubits = make([]*Qubit, hidden)

	for i := range qbm.VisibleQubits {
		qbm.VisibleQubits[i] = NewQubit([]float64{1, 0}) // |0⟩
	}

	for i := range qbm.HiddenQubits {
		qbm.HiddenQubits[i] = NewQubit([]float64{1, 0}) // |0⟩
	}

	// Inicializa pesos e bias
	rand.Seed(time.Now().UnixNano())
	qbm.Weights = make([][]float64, visible)
	for i := range qbm.Weights {
		qbm.Weights[i] = make([]float64, hidden)
		for j := range qbm.Weights[i] {
			qbm.Weights[i][j] = rand.NormFloat64() * 0.1
		}
	}

	qbm.VisibleBias = make([]float64, visible)
	qbm.HiddenBias = make([]float64, hidden)

	return qbm
}

func (qbm *QuantumBoltzmannMachine) Energy(visibleState, hiddenState []int) float64 {
	energy := 0.0

	// Termos de interação
	for i := range qbm.VisibleQubits {
		for j := range qbm.HiddenQubits {
			energy -= qbm.Weights[i][j] * float64(visibleState[i]) * float64(hiddenState[j])
		}
	}

	// Bias visíveis
	for i, b := range qbm.VisibleBias {
		energy -= b * float64(visibleState[i])
	}

	// Bias ocultos
	for j, b := range qbm.HiddenBias {
		energy -= b * float64(hiddenState[j])
	}

	return energy
}

func (qbm *QuantumBoltzmannMachine) Sample(k int) ([]int, []int) {
	// Pré-aquecimento quântico
	for _, q := range qbm.HiddenQubits {
		gates.ApplyHadamard(q)
		q.RotateX(math.Pi / 4 * qbm.Temperature)
	}

	// Amostragem por medição
	visible := make([]int, len(qbm.VisibleQubits))
	hidden := make([]int, len(qbm.HiddenQubits))

	for i := 0; i < k; i++ {
		for j, q := range qbm.VisibleQubits {
			if rand.Float64() < math.Pow(q.Coefficients[0], 2) {
				visible[j] = 0
			} else {
				visible[j] = 1
			}
		}

		for j, q := range qbm.HiddenQubits {
			if rand.Float64() < math.Pow(q.Coefficients[0], 2) {
				hidden[j] = 0
			} else {
				hidden[j] = 1
			}
		}
	}

	return visible, hidden
}
