// examples/hello_qubit.go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/gates"
)

func main() {
	// Configuração inicial
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Hello Qubit! Demonstração Básica ===")

	// 1. Criar um qubit no estado |0⟩
	q := core.NewQubit([]float64{1, 0})
	fmt.Printf("\nEstado Inicial:\n%v\n", q.Coefficients)

	// 2. Aplicar porta Hadamard
	gates.ApplyHadamard(q)
	fmt.Printf("\nApós Hadamard:\n[%.2f, %.2f]\n", q.Coefficients[0], q.Coefficients[1])

	// 3. Realizar múltiplas medições
	fmt.Println("\nResultados de 10 medições:")
	for i := 0; i < 10; i++ {
		// Criar nova instância para cada medição
		tempQ := core.NewQubit([]float64{q.Coefficients[0], q.Coefficients[1]})
		result := tempQ.Measure()
		fmt.Printf("Medição %d: %d\n", i+1, result)
	}

	// 4. Demonstração de emaranhamento
	fmt.Println("\nCriando par emaranhado:")
	q1 := core.NewQubit([]float64{1, 0})
	q2 := core.NewQubit([]float64{0, 1})

	q1.Entangle(q2)
	fmt.Printf("Estado Combinado:\n%v\n", q1.Coefficients)
}
