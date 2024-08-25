package formula

func CalculateMean(data []float64) float64 {
	var mean float64
	lengthOfData := len(data)
	var sum float64
	for _, value := range data {
		sum += value
	}
	mean = sum / float64(lengthOfData)
	return mean
}
