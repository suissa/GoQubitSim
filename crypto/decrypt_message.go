package crypto

import (
	"fmt"
	"strconv"

	"github.com/suissa/goqubitsim/core"
)

func DecryptMessage(qubits []*core.Qubit, receivedParity int) (string, error) {
	qr := core.NewQuantumRegister(0)
	qr.Qubits = qubits

	// Verificação de paridade
	calculatedParity := qr.CheckParity()

	if receivedParity != calculatedParity {
		return "", fmt.Errorf("interferência detectada durante a transmissão")
	}

	// Medição dos qubits
	measurements := qr.MeasureAll()

	// Decodificação da mensagem
	var decryptedMessage string
	for _, bit := range measurements {
		decryptedMessage += strconv.Itoa(bit)
	}

	return decryptedMessage, nil
}
