package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/algorithms"
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/core/gates"
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

	// 2. Operações com Portas Quânticas
	fmt.Println("\n2. Aplicando Portas Quânticas:")

	// Hadamard no Qubit 0
	gates.ApplyHadamard(q0)
	fmt.Printf("Qubit 0 após Hadamard: %.2f\n", q0.Coefficients)

	// Pauli-X no Qubit 1
	pauli := gates.NewPauliGates()
	pauli.ApplyX(q1)
	fmt.Printf("Qubit 1 após Pauli-X: %v\n", q1.Coefficients)

	// Rotação Z de 90 graus
	q0.RotateZ(math.Pi / 2)
	fmt.Printf("Qubit 0 após Rotação Z: %.2f\n", q0.Coefficients)

	// 3. Entrelaçamento e CNOT
	fmt.Println("\n3. Entrelaçamento e Operação CNOT:")

	// Cria novos qubits para demonstração do CNOT
	control := core.NewQubit([]float64{1, 0}) // |0⟩
	target := core.NewQubit([]float64{1, 0})  // |0⟩

	// Aplica Hadamard no qubit de controle
	gates.ApplyHadamard(control)
	fmt.Printf("\nEstado pré-CNOT:")
	fmt.Printf("\nControle: %.2f", control.Coefficients)
	fmt.Printf("\nAlvo: %v\n", target.Coefficients)

	// Aplica porta CNOT
	gates.ApplyCNOT(control, target)
	fmt.Printf("\nEstado pós-CNOT:")
	fmt.Printf("\nControle: %.2f", control.Coefficients)
	fmt.Printf("\nAlvo: %v\n", target.Coefficients)

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
		// Aplica Pauli-Z em um qubit aleatório
		pauli.ApplyZ(encrypted.Qubits[3])
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

	// 9. Demonstração de Algoritmos Quânticos
	fmt.Println("\n9. Demonstração de Algoritmos Quânticos:")

	// Deutsch-Jozsa
	djOracle := func(input, output *core.Qubit) {
		// Função constante f(x) = 0
		gates.ApplyHadamard(output)
	}
	dj := algorithms.NewDeutschJozsa(djOracle)
	fmt.Printf("\nDeutsch-Jozsa: Função é %s\n", map[bool]string{true: "constante", false: "balanceada"}[dj.Execute()])

	// Grover
	groverOracle := func(q *core.Qubit) bool {
		return q.Measure() == 3 // Busca pelo número 3
	}
	grov := algorithms.NewGrover(groverOracle, 2, 2)
	fmt.Printf("\nGrover: Resultado encontrado: %d\n", grov.Execute())

	// Bernstein-Vazirani
	bvOracle := func(q *core.Qubit) bool {
		// String escondida '101'
		q.RotateX(math.Pi)
		return q.Measure() == 1
	}
	bv := algorithms.NewBernsteinVazirani(bvOracle, 3)
	fmt.Printf("\nBernstein-Vazirani: String encontrada: %s\n", bv.FindHiddenString())

	// Shor
	shor := algorithms.NewShor(15, 4)
	f1, f2 := shor.Factor()
	fmt.Printf("\nShor: Fatores de 15: %d e %d\n", f1, f2)
}
