package main

import (
	"fmt"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/learning"
)

func main() {
	// Cria QBM com 2 visíveis e 2 ocultos
	qbm := core.NewQuantumBoltzmannMachine(2, 2)
	qbm.Temperature = 0.5

	// Dados XOR
	data := [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	// Treinamento
	learning.Train(qbm, data, 1000, 0.01)

	// Teste
	fmt.Println("Testando XOR:")
	for _, d := range data {
		visible, _ := qbm.Sample(100)
		result := visible[0] ^ visible[1]
		fmt.Printf("Entrada: %v Saída: %d\n", d, result)
	}

	// Visualização dos pesos
	fmt.Println("\nPesos aprendidos:")
	for i, row := range qbm.Weights {
		fmt.Printf("Visible %d: %v\n", i, row)
	}
}
