package analysis

// Rsi is Relative Strength Index (RSI). It is a momentum indicator that measures the magnitude
// of recent price changes to evaluate overbought and oversold conditions.
func (df *DataFrame) Rsi() ([]float64, []float64) {
	gains := make([]float64, len(df.Series))
	losses := make([]float64, len(df.Series))

	for i := 1; i < len(df.Series); i++ {
		difference := df.Series[i].Close - df.Series[i-1].Close

		if difference > 0 {
			gains[i] = difference
			losses[i] = 0
		} else {
			losses[i] = -difference
			gains[i] = 0
		}
	}

	meanGains := Sma(14, gains)
	meanLosses := Sma(14, losses)

	rsi := make([]float64, len(df.Series))
	rs := make([]float64, len(df.Series))

	for i := 0; i < len(rsi); i++ {
		rs[i] = meanGains[i] / meanLosses[i]
		rsi[i] = 100 - (100 / (1 + rs[i]))
	}

	return rs, rsi
}

// IsSmaShowUpTrend ...
func (df *DataFrame) IsSmaShowUpTrend() bool {
	closes := make([]float64, len(df.Series))

	for i := 0; i < len(df.Series); i++ {
		closes[i] = df.Series[i].Close
	}

	meanGains := Sma(10, closes)

	return closes[len(closes)-1] > meanGains[len(meanGains)-1]
}

// IsSmaShowDownTrend ...
func (df *DataFrame) IsSmaShowDownTrend() bool {
	closes := make([]float64, len(df.Series))

	for i := 0; i < len(df.Series); i++ {
		closes[i] = df.Series[i].Close
	}

	meanGains := Sma(10, closes)

	return closes[len(closes)-1] < meanGains[len(meanGains)-1]
}
