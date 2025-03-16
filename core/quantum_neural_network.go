package core

import (
	"math"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/core/gates"
)

type QuantumLayer struct {
	Weights        [][]float64
	RotationAngles []float64
}

type QuantumNetwork struct {
	Layers       []QuantumLayer
	LearningRate float64
}

func NewQuantumNetworks(layerSizes []int, learningRate float64) *QuantumNetwork {
	rand.Seed(time.Now().UnixNano())

	qnn := &QuantumNetwork{
		LearningRate: learningRate,
	}

	for i := 0; i < len(layerSizes)-1; i++ {
		layer := QuantumLayer{
			Weights:        make([][]float64, layerSizes[i]),
			RotationAngles: make([]float64, layerSizes[i+1]),
		}

		for j := range layer.Weights {
			layer.Weights[j] = make([]float64, layerSizes[i+1])
			for k := range layer.Weights[j] {
				layer.Weights[j][k] = rand.NormFloat64() * 0.1
			}
		}

		qnn.Layers = append(qnn.Layers, layer)
	}

	return qnn
}

func (qnn *QuantumNetwork) Forward(input []*Qubit) []*Qubit {
	currentState := input

	for _, layer := range qnn.Layers {
		// Aplica transformações quânticas
		newState := make([]*Qubit, len(layer.RotationAngles))

		for i := range newState {
			newState[i] = NewQubit([]float64{0, 0})

			// Combinação linear dos qubits anteriores
			for j, q := range currentState {
				angle := layer.Weights[j][i] * math.Pi
				q.RotateY(angle)
				newState[i].Entangle(q)
			}

			// Aplica rotação final
			newState[i].RotateZ(layer.RotationAngles[i])
			gates.ApplyHadamard(newState[i])
		}

		currentState = newState
	}

	return currentState
}

func (qnn *QuantumNetwork) Backward(input, output []*Qubit, target []float64) {
	// Calcula gradientes quânticos
	gradients := make([][][]float64, len(qnn.Layers))

	// Propagação reversa
	for l := len(qnn.Layers) - 1; l >= 0; l-- {
		layer := qnn.Layers[l]
		gradients[l] = make([][]float64, len(layer.Weights))

		for i := range layer.Weights {
			gradients[l][i] = make([]float64, len(layer.Weights[i]))

			for j := range layer.Weights[i] {
				// Cálculo do gradiente usando parâmetros quânticos
				delta := (output[j].Coefficients[0] - target[j]) * qnn.LearningRate
				gradients[l][i][j] = delta * input[i].Coefficients[0]
			}
		}
	}

	// Atualiza pesos
	for l := range qnn.Layers {
		for i := range qnn.Layers[l].Weights {
			for j := range qnn.Layers[l].Weights[i] {
				qnn.Layers[l].Weights[i][j] -= gradients[l][i][j]
			}
		}
	}
}
