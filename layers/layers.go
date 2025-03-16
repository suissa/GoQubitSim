package layer

import (
	"math"
	"math/rand"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
)

// Interface para camadas quânticas
type QuantumLayerInterface interface {
	Forward(input []*core.Qubit) []*core.Qubit
	Backward(gradients [][]float64) [][]float64
}

// Camada de Atenção Quântica
type QuantumAttention struct {
	QueryWeights [][]float64
	KeyWeights   [][]float64
	ValueWeights [][]float64
}

// Camada Densa Quântica
type QuantumDense struct {
	Weights        [][]float64
	RotationAngles []float64
	NumQubits      int
}

// Camada de Pooling Quântico
type QuantumPooling struct {
	PoolSize int
}

func NewQuantumDense(inputSize, outputSize int) *QuantumDense {
	// Inicialização quântica dos pesos
	weights := make([][]float64, inputSize)
	for i := range weights {
		weights[i] = make([]float64, outputSize)
		for j := range weights[i] {
			weights[i][j] = rand.NormFloat64() * math.Pi / math.Sqrt(float64(inputSize))
		}
	}

	return &QuantumDense{
		Weights:        weights,
		RotationAngles: make([]float64, outputSize),
		NumQubits:      outputSize,
	}
}

func (qa *QuantumAttention) Forward(input []*core.Qubit) []*core.Qubit {
	numQubits := len(input)
	output := make([]*core.Qubit, numQubits)

	// 1. Gerar Queries, Keys e Values quânticos
	queries := make([]*core.Qubit, numQubits)
	keys := make([]*core.Qubit, numQubits)
	values := make([]*core.Qubit, numQubits)

	// Aplicar pesos quânticos
	for i := range input {
		queries[i] = qa.applyLinearTransform(input, qa.QueryWeights[i])
		keys[i] = qa.applyLinearTransform(input, qa.KeyWeights[i])
		values[i] = qa.applyLinearTransform(input, qa.ValueWeights[i])
	}

	// 2. Calcular scores de atenção quântica
	attentionQubits := make([]*core.Qubit, numQubits)
	for i := range attentionQubits {
		attentionQubits[i] = NewQubit([]float64{1, 0})
		for j := range queries {
			// Entrelaçar queries e keys
			qa.entanglePair(queries[j], keys[j], attentionQubits[i])
		}
		gates.ApplyHadamard(attentionQubits[i])
	}

	// 3. Aplicar atenção aos valores
	for i := range output {
		output[i] = NewQubit([]float64{0, 0})
		for j := range values {
			// Porta controlada pela atenção
			qa.applyControlledRotation(values[j], attentionQubits[j], output[i])
		}
		output[i].Normalize()
	}

	return output
}

// Métodos auxiliares
func (qa *QuantumAttention) applyLinearTransform(qubits []*core.Qubit, weights []float64) *core.Qubit {
	result := NewQubit([]float64{0, 0})
	for i, q := range qubits {
		angle := weights[i] * math.Pi
		temp := q.Clone()
		temp.RotateY(angle)
		result.Entangle(temp)
	}
	return result
}

func (qa *QuantumAttention) entanglePair(q1, q2, control *core.Qubit) {
	gates.ControlledSwap(control, q1, q2)
}

// func (qa *QuantumAttention) entanglePair(q1, q2, control *core.Qubit) {
// 	// Porta SWAP controlada
// 	if control.Measure() == 1 {
// 			q1.Coefficients[0], q1.Coefficients[1], q2.Coefficients[0], q2.Coefficients[1] =
// 					q2.Coefficients[0], q2.Coefficients[1], q1.Coefficients[0], q1.Coefficients[1]
// 	}
// }

func (qa *QuantumAttention) applyControlledRotation(source, control, target *core.Qubit) {
	angle := math.Atan2(source.Coefficients[1], source.Coefficients[0])
	if control.Measure() == 1 {
		target.RotateY(angle)
		target.Entangle(source)
	}
}
