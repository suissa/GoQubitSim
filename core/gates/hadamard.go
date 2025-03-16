package gates

import (
	"math"

	"github.com/suissa/goqubitsim/core"
)

func ApplyHadamard(q *core.Qubit) {
	h := 1 / math.Sqrt(2)
	newCoefficients := make([]float64, 2)

	newCoefficients[0] = h * (q.Coefficients[0] + q.Coefficients[1])
	newCoefficients[1] = h * (q.Coefficients[0] - q.Coefficients[1])

	q.Coefficients = newCoefficients
}
