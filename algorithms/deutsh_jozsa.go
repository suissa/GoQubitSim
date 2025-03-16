package algorithms

import (
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

type DeutschJozsa struct {
	Oracle func(*core.Qubit, *core.Qubit)
}

func NewDeutschJozsa(oracle func(*core.Qubit, *core.Qubit)) *DeutschJozsa {
	return &DeutschJozsa{Oracle: oracle}
}

func (dj *DeutschJozsa) Execute() bool {
	// Inicializa qubits
	input := core.NewQubit([]float64{1, 0})
	output := core.NewQubit([]float64{0, 1})

	// Aplica Hadamard
	gates.ApplyHadamard(input)
	gates.ApplyHadamard(output)

	// Executa oracle
	dj.Oracle(input, output)

	// Aplica Hadamard novamente
	gates.ApplyHadamard(input)

	// Medição
	return input.Measure() == 0
}
