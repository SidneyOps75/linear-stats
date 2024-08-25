package formula

import (
	"log"
	"math"
)

func PearsonCorrelation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		log.Fatal("x and y must have a non-zero length")
	}
	meanX := CalculateMean(x)
	meanY := CalculateMean(y)

	var sumNumerator float64
	var sumXDeviation float64
	var sumYDeviation float64
	for i := 0; i < len(x); i++ {
		xDiff := x[i] - meanX
		yDiff := y[i] - meanY
		sumNumerator += xDiff * yDiff
		sumXDeviation += math.Pow(xDiff, 2)
		sumYDeviation += math.Pow(yDiff, 2)
	}

	denominator := math.Sqrt(sumXDeviation * sumYDeviation)
	if denominator == 0 {
		return 0
	}
	return sumNumerator / denominator

}
