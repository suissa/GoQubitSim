package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type QuantumKeyConfig struct {
	KeyLength     int
	Bases         []rune // 'X' ou 'Z' para diferentes bases quânticas
	PhotonStates  []int  // Estados dos fótons (0 ou 1)
	SiftedKey     []int
	MatchingBases []bool
}

func GenerateQuantumKey(length int) *QuantumKeyConfig {
	config := &QuantumKeyConfig{
		KeyLength:    length,
		Bases:        make([]rune, length*2), // Gera o dobro para compensar a peneiração
		PhotonStates: make([]int, length*2),
	}

	// Geração aleatória de estados e bases
	config.generateQuantumStates()

	// Simulação do processo de peneiração quântica
	config.siftKey()

	return config
}

func (q *QuantumKeyConfig) generateQuantumStates() {
	for i := 0; i < len(q.Bases); i++ {
		// Escolha aleatória da base (X ou Z)
		q.Bases[i] = 'Z'
		if rand.Float64() < 0.5 {
			q.Bases[i] = 'X'
		}

		// Geração do estado quântico
		q.PhotonStates[i] = rand.Intn(2)
	}
}

func (q *QuantumKeyConfig) siftKey() {
	// Simula a troca de bases entre Alice e Bob
	bobBases := make([]rune, len(q.Bases))
	for i := range bobBases {
		bobBases[i] = 'Z'
		if rand.Float64() < 0.5 {
			bobBases[i] = 'X'
		}
	}

	// Seleciona apenas os bits com bases coincidentes
	for i := 0; i < len(q.Bases) && len(q.SiftedKey) < q.KeyLength; i++ {
		if q.Bases[i] == bobBases[i] {
			q.SiftedKey = append(q.SiftedKey, q.PhotonStates[i])
			q.MatchingBases = append(q.MatchingBases, true)
		} else {
			q.MatchingBases = append(q.MatchingBases, false)
		}
	}
}

// Implementação de verificação de integridade quântica
func (q *QuantumKeyConfig) CheckEavesdropping(sampleSize int) float64 {
	errorRate := 0.0
	checked := 0

	for i := 0; i < len(q.SiftedKey) && checked < sampleSize; i++ {
		if i%5 == 0 { // Verifica 20% dos bits aleatoriamente
			// Simula detecção de espionagem (deveria comparar com Bob)
			if rand.Float64() < 0.1 { // Assume 10% de taxa de erro por espionagem
				errorRate += 1.0
			}
			checked++
		}
	}

	if checked > 0 {
		return errorRate / float64(checked)
	}
	return 0.0
}
