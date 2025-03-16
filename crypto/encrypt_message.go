package crypto

import (
	"math"

	"github.com/suissa/goqubitsim/core"
)

type EncryptedMessage struct {
	Qubits []*core.Qubit
	Parity int
}

func EncryptMessage(message []int) *EncryptedMessage {
	qr := core.NewQuantumRegister(0)

	// Prepara cada qubit com o bit da mensagem
	for _, bit := range message {
		coefficient0 := float64(bit)
		coefficient1 := math.Sqrt(1 - float64(bit))
		qubit := core.NewQubit([]float64{coefficient0, coefficient1})
		qr.AddQubit(qubit)
	}

	// Calcula a paridade antes da medição
	parity := qr.CheckParity()

	return &EncryptedMessage{
		Qubits: qr.Qubits,
		Parity: parity,
	}
}
