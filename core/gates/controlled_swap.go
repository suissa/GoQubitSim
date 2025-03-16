// gates/controlled_swap.go
package gates

import "github.com/suissa/goqubitsim/core"

func ControlledSwap(control *core.Qubit, q1, q2 *core.Qubit) {
	// Implementação usando decomposição em portas CNOT
	// SWAP controlado = 3 CNOTs controlados
	CNOT(control, q1, q2)
	CNOT(control, q2, q1)
	CNOT(control, q1, q2)
}

// Implemente CNOT controlado conforme sua biblioteca
func CNOT(control *core.Qubit, target1, target2 *core.Qubit) {
	// Lógica da porta CNOT estendida para SWAP
}
