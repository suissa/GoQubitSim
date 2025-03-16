package main

import (
	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/qml"
	"github.com/suissa/goqubitsim/visualization/web"
)

func main() {
	// 1. Criar QNN base pré-treinada
	baseQNN := core.NewQuantumNetwork([]core.QuantumLayer{
		core.NewQuantumDense(4, 8),
		core.NewQuantumDense(8, 4),
	})

	// 2. Configurar transfer learning
	transferModel := qml.NewTransferLearner(baseQNN, []int{0})
	transferModel.AddAdapterLayer(core.NewQuantumDense(4, 2))

	// 3. Adicionar correção de erros
	qec := core.NewShorCode()
	transferModel.BaseModel.ErrorCorrection = qec

	// 4. Configurar visualização
	dashboard := web.Dashboard{
		StateVisualizer: web.QuditStateVisualizer,
		CircuitRenderer: core.DrawQuantumCircuit,
	}
	go dashboard.Serve("8080")

	// 5. Treinamento com qudits
	quditDataset := loadQuditDataset()
	transferModel.FineTune(quditDataset, 100)

	// 6. Execução com entrada qudit
	input := []*core.Qudit{
		core.NewQudit(3),
		core.NewQudit(3),
	}
	output := transferModel.Forward(input)
}
