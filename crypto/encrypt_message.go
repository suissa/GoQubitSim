package crypto

import (
	"math"
)

type EncryptedMessage struct {
	Qubits []*Qubit
	Parity int
}

func EncryptMessage(message []int) *EncryptedMessage {
	qr := NewQuantumRegister(0)

	// Prepara cada qubit com o bit da mensagem
	for _, bit := range message {
		coefficient0 := float64(bit)
		coefficient1 := math.Sqrt(1 - float64(bit))
		qubit := NewQubit([]float64{coefficient0, coefficient1})
		qr.AddQubit(qubit)
	}

	// Calcula a paridade antes da medição
	parity := qr.CheckParity()

	return &EncryptedMessage{
		Qubits: qr.Qubits,
		Parity: parity,
	}
}
