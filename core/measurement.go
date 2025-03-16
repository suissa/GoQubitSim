package core

import (
	"fmt"
	"math"
)

func main() {
	// Cria dois qubits
	qubit1 := NewQubit([]float64{1, 0}) // Estado |0>
	qubit2 := NewQubit([]float64{0, 1}) // Estado |1>

	// Realiza uma rotação em torno do eixo X no primeiro qubit
	qubit1.RotateX(math.Pi / 2) // Rotação de 90 graus

	// Realiza uma rotação em torno do eixo Y no segundo qubit
	qubit2.RotateY(math.Pi / 4) // Rotação de 45 graus

	// Entrelaça os dois qubits
	qubit1.Entangle(qubit2)

	// Realiza medições
	measurement1 := qubit1.Measure()
	measurement2 := qubit2.Measure()

	// Exibe os resultados
	fmt.Println("Medição do qubit 1:", measurement1) // Resultado 0 ou 1
	fmt.Println("Medição do qubit 2:", measurement2) // Resultado 0 ou 1
}
