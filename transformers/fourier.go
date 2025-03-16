package transformers

import (
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

type Fourier struct {
	NumQubits int
}

func NewFourier(n int) *Fourier {
	return &Fourier{NumQubits: n}
}

func (qft *Fourier) Apply(register *core.QuantumRegister) {
	for i := 0; i < qft.NumQubits; i++ {
		gates.ApplyHadamard(register.Qubits[i]) // Corrigido!

		for j := 1; j < qft.NumQubits-i; j++ {
			control := i
			target := i + j
			applyControlledPhase(register.Qubits[control], register.Qubits[target], j)
		}
	}
	reverseQubits(register.Qubits)
}

func applyControlledPhase(control, target *core.Qubit, angle float64) {
	// Porta de fase controlada
	if control.Measure() == 1 {
		target.RotateZ(angle)
	}
}

func reverseQubits(qubits []*core.Qubit) {
	n := len(qubits)
	for i := 0; i < n/2; i++ {
		qubits[i], qubits[n-1-i] = qubits[n-1-i], qubits[i]
	}
}

// QFT Inversa
func (qft *Fourier) ApplyInverse(register *core.QuantumRegister) {
	reverseQubits(register.Qubits)
	qft.Apply(register)
	reverseQubits(register.Qubits)
}
