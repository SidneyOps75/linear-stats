package formula

func CalculateMean(data []float64) float64 {
	var mean float64
	length := len(data)
	var sum float64
	for _, n := range data {
		sum += n
		mean = sum / float64(length)
	}
	return mean
}
