package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/crypto"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Demonstração do Sistema Quântico GoQubitSim ===")

	// 1. Criação de Qubits Básicos
	fmt.Println("\n1. Criando Qubits Básicos:")
	q0 := core.NewQubit([]float64{1, 0}) // |0⟩
	q1 := core.NewQubit([]float64{0, 1}) // |1⟩
	fmt.Printf("Qubit 0 Estado Inicial: %v\n", q0.Coefficients)
	fmt.Printf("Qubit 1 Estado Inicial: %v\n", q1.Coefficients)

	// 2. Operações Quânticas
	fmt.Println("\n2. Aplicando Portas Quânticas:")
	q0.RotateX(math.Pi / 2)
	q1.RotateY(math.Pi / 4)
	fmt.Printf("Qubit 0 após Rotação X: %v\n", q0.Coefficients)
	fmt.Printf("Qubit 1 após Rotação Y: %v\n", q1.Coefficients)

	// 3. Entrelaçamento Quântico
	fmt.Println("\n3. Entrelaçando Qubits:")
	q0.Entangle(q1)
	fmt.Println("Qubits Entrelaçados:")
	fmt.Printf("Estado Combinado: %v\n", q0.Coefficients)

	// 4. Geração de Chave Quântica
	fmt.Println("\n4. Gerando Chave Quântica:")
	keyConfig := crypto.GenerateQuantumKey(128)
	fmt.Printf("Chave Gerada: %v...\n", keyConfig.SiftedKey[:16])
	fmt.Printf("Tamanho da Chave: %d bits\n", len(keyConfig.SiftedKey))

	// 5. Criptografia de Mensagem
	fmt.Println("\n5. Criptografando Mensagem:")
	message := []int{0, 1, 0, 0, 1, 1, 0, 1} // "01001101"
	encrypted := crypto.EncryptMessage(message)
	fmt.Printf("Paridade da Mensagem: %d\n", encrypted.Parity)
	fmt.Printf("Qubits Transmitidos: %d\n", len(encrypted.Qubits))

	// 6. Simulação de Transmissão com Interferência
	fmt.Println("\n6. Simulando Transmissão:")
	if rand.Float32() < 0.3 { // 30% chance de interferência
		fmt.Println(">>> Interferência Detectada! <<<")
		encrypted.Qubits[2].RotateX(math.Pi / 4) // Altera um qubit
	}

	// 7. Decriptação da Mensagem
	fmt.Println("\n7. Decifrando Mensagem:")
	decrypted, err := crypto.DecryptMessage(encrypted.Qubits, encrypted.Parity)

	if err != nil {
		fmt.Println("Erro na Decriptação:", err)
	} else {
		fmt.Printf("Mensagem Decifrada: %s\n", decrypted)
		fmt.Printf("Mensagem Original:  %v\n", message)
	}

	// 8. Verificação Quântica de Segurança
	fmt.Println("\n8. Verificação de Segurança:")
	errorRate := keyConfig.CheckEavesdropping(50)
	fmt.Printf("Taxa de Erro Detectada: %.2f%%\n", errorRate*100)
	if errorRate < 0.1 {
		fmt.Println("Status: Sistema Seguro!")
	} else {
		fmt.Println("Status: Possível Ataque Detectado!")
	}
}
