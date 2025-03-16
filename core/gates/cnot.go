package gates

import (
	"github.com/suissa/goqubitsim/core"
)

func ApplyCNOT(control *core.Qubit, target *core.Qubit) {
	// Entrelaçar os qubits primeiro
	control.Entangle(target)

	// Aplicar porta CNOT
	if control.Measure() == 1 {
		target.Coefficients[0], target.Coefficients[1] =
			target.Coefficients[1], target.Coefficients[0]
	}
}
