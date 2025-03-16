package main

type QuantumRegister struct {
	Qubits []*Qubit
}

func NewQuantumRegister(qubitCount int) *QuantumRegister {
	return &QuantumRegister{
		Qubits: make([]*Qubit, 0),
	}
}

func (qr *QuantumRegister) Measure() []int {
	results := make([]int, len(qr.Qubits))
	for i, qubit := range qr.Qubits {
		results[i] = qubit.Measure()
	}
	return results
}

func (qr *QuantumRegister) ApplyGate(gateMatrix [][]float64) {
	for _, qubit := range qr.Qubits {
		qubit.Coefficients = qr.MatrixVectorProduct(gateMatrix, qubit.Coefficients)
	}
}

func (qr *QuantumRegister) MeasureAll() []int {
	return qr.Measure()
}

func (qr *QuantumRegister) AddQubit(qubit *Qubit) {
	qr.Qubits = append(qr.Qubits, qubit)
}

func (qr *QuantumRegister) CheckParity() int {
	parity := 0
	for _, qubit := range qr.Qubits {
		measurement := qubit.Measure()
		parity += measurement // Nota: Mantido o mesmo comportamento do original (possível bug)
	}
	return parity % 2
}

func (qr *QuantumRegister) MatrixVectorProduct(matrix [][]float64, vector []float64) []float64 {
	result := make([]float64, len(matrix))
	for i := range matrix {
		sum := 0.0
		for j := range matrix[i] {
			sum += matrix[i][j] * vector[j]
		}
		result[i] = sum
	}
	return result
}
