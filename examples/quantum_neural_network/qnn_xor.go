package main

import (
	"fmt"
	"math"

	"github.com/suissa/goqubitsim/core"
)

func main() {
	// Cria QuantumNetwork com arquitetura 2-2-1
	qnn := core.NewQuantumNetwork([]int{2, 2, 1}, 0.1)

	// Dados XOR
	dataset := []struct {
		input  []*core.Qubit
		target []float64
	}{
		{
			input:  []*core.Qubit{core.NewQubit([]float64{1, 0}), core.NewQubit([]float64{1, 0})},
			target: []float64{0},
		},
		{
			input:  []*core.Qubit{core.NewQubit([]float64{1, 0}), core.NewQubit([]float64{0, 1})},
			target: []float64{1},
		},
		{
			input:  []*core.Qubit{core.NewQubit([]float64{0, 1}), core.NewQubit([]float64{1, 0})},
			target: []float64{1},
		},
		{
			input:  []*core.Qubit{core.NewQubit([]float64{0, 1}), core.NewQubit([]float64{0, 1})},
			target: []float64{0},
		},
	}

	// Treinamento
	for epoch := 0; epoch < 100; epoch++ {
		totalLoss := 0.0

		for _, data := range dataset {
			// Forward pass
			output := qnn.Forward(data.input)

			// Calcula perda
			loss := math.Pow(output[0].Coefficients[0]-data.target[0], 2)
			totalLoss += loss

			// Backward pass
			qnn.Backward(data.input, output, data.target)
		}

		fmt.Printf("Epoch %d Loss: %.4f\n", epoch+1, totalLoss/4)
	}

	// Teste
	fmt.Println("\nResultados:")
	for _, data := range dataset {
		output := qnn.Forward(data.input)
		fmt.Printf("Entrada: [%.0f, %.0f] Saída: %.2f Esperado: %.0f\n",
			data.input[0].Coefficients[0],
			data.input[1].Coefficients[0],
			output[0].Coefficients[0],
			data.target[0],
		)
	}
}
