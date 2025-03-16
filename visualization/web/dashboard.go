package web

import (
	"fmt"
	"math"
	"net/http"
	"text/template"

	"github.com/suissa/goqubitsim/core"
)

type Dashboard struct {
	StateVisualizer  func([]*core.Qudit) string
	CircuitRenderer  func([]core.QuantumGate) string
	TelemetryMonitor func() map[string]float64
}

func (d *Dashboard) Serve(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
		tmpl.Execute(w, map[string]interface{}{
			"State":     d.StateVisualizer(getCurrentState()),
			"Circuit":   d.CircuitRenderer(getCurrentCircuit()),
			"Telemetry": d.TelemetryMonitor(),
		})
	})
	http.ListenAndServe(":"+port, nil)
}

// Exemplo de visualização de estado qudit
func QuditStateVisualizer(qudits []*core.Qudit) string {
	visualization := ""
	for i, qd := range qudits {
		visualization += fmt.Sprintf("Qudit %d:\n", i)
		for state, coeff := range qd.Coefficients {
			prob := math.Pow(real(coeff), 2) + math.Pow(imag(coeff), 2)
			visualization += fmt.Sprintf("|%d⟩: %.2f%%\n", state, prob*100)
		}
	}
	return visualization
}
