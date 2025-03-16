package learning

import (
	"math/rand"

	"github.com/suissa/goqubitsim/core"
)

func (qbm *core.QuantumBoltzmannMachine) Train(data [][]int, steps int, lr float64) {
	for step := 0; step < steps; step++ {
		// Fase positiva (dados reais)
		posVisible := data[rand.Intn(len(data))]
		_, posHidden := qbm.Sample(1)

		// Fase negativa (amostras do modelo)
		negVisible, negHidden := qbm.Sample(100)

		// Atualização de pesos
		for i := range qbm.Weights {
			for j := range qbm.Weights[i] {
				positive := float64(posVisible[i]) * float64(posHidden[j])
				negative := float64(negVisible[i]) * float64(negHidden[j])
				qbm.Weights[i][j] += lr * (positive - negative)
			}
		}

		// Atualização de bias
		for i := range qbm.VisibleBias {
			qbm.VisibleBias[i] += lr * (float64(posVisible[i]) - float64(negVisible[i]))
		}

		for j := range qbm.HiddenBias {
			qbm.HiddenBias[j] += lr * (float64(posHidden[j]) - float64(negHidden[j]))
		}
	}
}
