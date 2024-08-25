package formula

import (
	"log"
	"math"
)

func LinearRegression(x, y []float64) (float64, float64) {
	if len(x) != len(y) || len(x) == 0 {
		log.Fatal("x and y must have the same non-zero length")
	}
	meanX := CalculateMean(x)
	meanY := CalculateMean(y)

	var sumNumerator float64
	var sumDenominator float64
	for i := 0; i < len(x); i++ {
		sumNumerator += (x[i] - meanX) * (y[i] - meanY)
		sumDenominator += math.Pow((x[i] - meanX), 2)
	}
	m := sumNumerator / sumDenominator
	b := meanY - m*meanX
	return m, b
}
