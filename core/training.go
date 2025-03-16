package core

import "sync"

type QuantumBatchTrainer struct {
	BatchSize    int
	CurrentBatch [][]*Qubit
	Gradients    [][][][]float64
}

// Implementação do Adam Quântico
type QuantumAdamOptimizer struct {
	Beta1, Beta2 float64
	M, V         [][][]float64
}

// AccumulateGradients calcula o gradiente quântico de uma lista de inputs e targets,
// processando cada batch em paralelo e acumulando os gradientes em qbt.Gradients.
// O cálculo de gradiente é feito com um forward pass e um backward pass,
// utilizando as layer.Backward() e layer.Forward() para propagar os erros e
// calcular as derivadas parciais.
func (qbt *QuantumBatchTrainer) AccumulateGradients(inputs [][]*Qubit, targets [][]float64) {
	var wg sync.WaitGroup
	gradChan := make(chan [][][]float64, len(inputs))

	// Processamento paralelo de batches
	for _, input := range inputs {
		wg.Add(1)
		go func(in []*Qubit) {
			defer wg.Done()

			// Forward pass
			output := qbt.Model.Forward(in)

			// Cálculo quântico do gradiente
			gradients := make([][][]float64, len(qbt.Model.Layers))
			currentGrad := make([]float64, len(output))

			for i := range currentGrad {
				currentGrad[i] = output[i].Coefficients[0] - targets[i][0]
			}

			// Backward pass
			for l := len(qbt.Model.Layers) - 1; l >= 0; l-- {
				layer := qbt.Model.Layers[l]
				gradients[l] = layer.Backward(currentGrad)
				currentGrad = flattenGradients(gradients[l])
			}

			gradChan <- gradients
		}(input)
	}

	go func() {
		wg.Wait()
		close(gradChan)
	}()

	// Acumulação de gradientes
	for grad := range gradChan {
		for l := range grad {
			for i := range grad[l] {
				for j := range grad[l][i] {
					qbt.Gradients[l][i][j] += grad[l][i][j]
				}
			}
		}
	}
}

// UpdateWeights atualiza os pesos da rede quântica baseado nos gradientes quânticos
// acumulados em qbt.Gradients.
//
// Para cada camada densa, os pesos são atualizados com a fórmula:
// w_i,j -= learning_rate * gradients[l][i][j] / batch_size
//
// Após atualizar os pesos, os gradientes são resetados para zero.
func (qbt *QuantumBatchTrainer) UpdateWeights() {
	batchSize := float64(qbt.BatchSize)

	for l := range qbt.Model.Layers {
		if dense, ok := qbt.Model.Layers[l].(*QuantumDense); ok {
			for i := range dense.Weights {
				for j := range dense.Weights[i] {
					dense.Weights[i][j] -= qbt.LearningRate * qbt.Gradients[l][i][j] / batchSize
				}
			}
		}
	}

	// Reset gradients
	qbt.Gradients = make([][][][]float64, len(qbt.Model.Layers))
}
