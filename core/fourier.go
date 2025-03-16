package core

import (
	"math"

	"github.com/suissa/goqubitsim/gates"
)

type Fourier struct {
	NumQubits int
}

func NewFourier(n int) *Fourier {
	return &Fourier{NumQubits: n}
}

func (qft *Fourier) Apply(register *QuantumRegister) {
	// Aplica a QFT em cada qubit
	for i := 0; i < qft.NumQubits; i++ {
		// Aplica Hadamard no qubit atual
		gates.ApplyHadamard(register.Qubits[i])

		// Aplica portas de fase controladas
		for j := 1; j < qft.NumQubits-i; j++ {
			control := i
			target := i + j
			angle := math.Pi / math.Pow(2, float64(j))
			applyControlledPhase(register.Qubits[control], register.Qubits[target], angle)
		}
	}

	// Inverte a ordem dos qubits
	reverseQubits(register.Qubits)
}

func applyControlledPhase(control, target *Qubit, angle float64) {
	// Porta de fase controlada
	if control.Measure() == 1 {
		target.RotateZ(angle)
	}
}

func reverseQubits(qubits []*Qubit) {
	n := len(qubits)
	for i := 0; i < n/2; i++ {
		qubits[i], qubits[n-1-i] = qubits[n-1-i], qubits[i]
	}
}

// QFT Inversa
func (qft *Fourier) ApplyInverse(register *QuantumRegister) {
	reverseQubits(register.Qubits)
	qft.Apply(register)
	reverseQubits(register.Qubits)
}
