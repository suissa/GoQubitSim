package main

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Qubit struct {
	Coefficients []float64
}

func NewQubit(coefficients []float64) *Qubit {
	return &Qubit{Coefficients: coefficients}
}

func (q *Qubit) Measure() int {
	probabilities := make([]float64, len(q.Coefficients))
	for i, c := range q.Coefficients {
		probabilities[i] = math.Pow(math.Abs(c), 2)
	}

	randomNum := rand.Float64()
	cumulativeProb := 0.0

	for i, prob := range probabilities {
		cumulativeProb += prob
		if randomNum < cumulativeProb {
			return i
		}
	}

	return len(q.Coefficients) - 1
}

func (q *Qubit) RotateX(angle float64) {
	cos := math.Cos(angle / 2)
	sin := math.Sin(angle / 2)
	matrix := [][]float64{
		{cos, -sin},
		{-sin, cos},
	}
	q.Coefficients = q.applyMatrix(matrix, q.Coefficients)
}

func (q *Qubit) RotateY(angle float64) {
	cos := math.Cos(angle / 2)
	sin := math.Sin(angle / 2)
	matrix := [][]float64{
		{cos, -sin},
		{sin, cos},
	}
	q.Coefficients = q.applyMatrix(matrix, q.Coefficients)
}

func (q *Qubit) RotateZ(angle float64) {
	expMinus := math.Exp(-angle / 2)
	expPlus := math.Exp(angle / 2)
	matrix := [][]float64{
		{expMinus, 0},
		{0, expPlus},
	}
	q.Coefficients = q.applyMatrix(matrix, q.Coefficients)
}

func (q *Qubit) applyMatrix(matrix [][]float64, coefficients []float64) []float64 {
	result := make([]float64, len(matrix))
	for i := range matrix {
		sum := 0.0
		for j := range matrix[i] {
			sum += matrix[i][j] * coefficients[j]
		}
		result[i] = sum
	}
	return result
}

func (q *Qubit) Entangle(other *Qubit) {
	newCoefficients := make([]float64, 0, len(q.Coefficients)*len(other.Coefficients))
	for _, c1 := range q.Coefficients {
		for _, c2 := range other.Coefficients {
			newCoefficients = append(newCoefficients, c1*c2)
		}
	}
	q.Coefficients = newCoefficients
}