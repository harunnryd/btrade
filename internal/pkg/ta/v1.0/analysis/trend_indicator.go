package analysis

// Sma is Simple Moving Average (SMA).
func Sma(period int, values []float64) []float64 {
	result := make([]float64, len(values))
	sum := float64(0)

	for i, value := range values {
		count := i + 1
		sum += value

		if i >= period {
			sum -= values[i-period]
			count = period
		}

		result[i] = sum / float64(count)
	}

	return result
}
