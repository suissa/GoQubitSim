package main

import (
	"fmt"
	"math"
	
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/gpu"
)

func main() {
	// Configurar rede com diferentes camadas
	layers := []core.QuantumLayer{
		core.NewQuantumDense(4, 8),
		&core.QuantumAttention{...},
		&core.QuantumPooling{PoolSize: 2},
		core.NewQuantumDense(4, 1),
	}

	qnn := core.NewQuantumNetwork(layers,
		core.WithRegularizer(&core.EntanglementRegularizer{Lambda: 0.01}),
		core.WithBatchSize(32),
		core.WithGPUAccelerator(&gpu.QuantumGPUAccelerator{}),
	)

	// Treinamento com dataset quântico
	dataset := loadQuantumDataset()
	
	// Loop de treinamento
	for epoch := 0; epoch < 100; epoch++ {
		qnn.Train(dataset)
		
		// Validação
		loss := qnn.Evaluate(validationData)
		fmt.Printf("Epoch %d Loss: %.4f Entanglement: %.2f\n", 
			epoch, loss, qnn.GetEntanglement())
	}
}