package analysis

import (
	"math"
	"sort"
	"time"
)

// Action ...
type Action int

const (
	SELL Action = -1
	HOLD Action = 0
	BUY  Action = 1
)

// Asset ...
type Asset struct {
	Time  time.Time
	High  float64
	Low   float64
	Open  float64
	Close float64
}

type Series []Asset

func (df *DataFrame) Sort() {
	sort.Slice(df.Series, func(i, j int) bool {
		return df.Series[i].Time.Before(df.Series[j].Time)
	})

	if df.LastIndex > 0 {
		df.Last = df.Series[df.LastIndex]
		df.Prev = df.Series[df.LastIndex-1]
	}
}

// DataFrame ...
type DataFrame struct {
	Series    Series
	LastIndex int
	Last      Asset
	Prev      Asset
	DebugMode bool
}

// NewDataFrame ...
func NewDataFrame() *DataFrame {
	return &DataFrame{}
}

// Add ...
func (df *DataFrame) Add(asset Asset) {
	df.Series = append(df.Series, asset)
	df.LastIndex = len(df.Series) - 1
}

// Next ...
func (s Series) Next() *DataFrame {
	df := new(DataFrame)
	df.Series = s
	df.LastIndex = len(s) - 1
	df.Last = df.Series[df.LastIndex]
	df.Prev = df.Series[df.LastIndex-1]

	return df
}

// StatusSummary ...
func (df *DataFrame) StatusSummary() (actions []string, summaries []string) {
	if df.IsSpinningTop() {
		summaries = append(summaries, "SpinningTop")
		actions = append(actions, "Reversal")
	}

	if df.IsHammer() {
		summaries = append(summaries, "Hammer")
		actions = append(actions, "ReversalUpwards")
	}

	if df.IsHangingMan() {
		summaries = append(summaries, "HangingMan")
		actions = append(actions, "ReversalDownwards")
	}

	if df.IsInvertedHammer() {
		summaries = append(summaries, "InvertedHammer")
	}

	if df.IsShootingStar() {
		summaries = append(summaries, "ShootingStar")
		actions = append(actions, "ReversalDownwards")
	}

	if df.IsEveningStar() {
		summaries = append(summaries, "EveningStar")
		actions = append(actions, "Top")
	}

	if df.IsMorningStar() {
		summaries = append(summaries, "MorningStar")
		actions = append(actions, "Bottom")
	}

	if df.IsInsideDown() {
		summaries = append(summaries, "InsideDown")
	}

	if df.IsInsideUp() {
		summaries = append(summaries, "InsideUp")
	}

	if df.IsThreeBlackCrows() {
		summaries = append(summaries, "ThreeBlackCrows")
	}

	if df.IsThreeWhiteSoldiers() {
		summaries = append(summaries, "ThreeWhiteSoldiers")
	}

	if df.IsBullishBeltHold() {
		summaries = append(summaries, "BullishBeltHold")
	}

	if df.IsBearishBeltHold() {
		summaries = append(summaries, "BearishBeltHold")
	}

	if df.IsDownTrend() {
		summaries = append(summaries, "TheDownTrend")
	}

	if df.IsUpTrend() {
		summaries = append(summaries, "TheUpTrend")
	}

	return
}

// IsFuzzyEqual ...
func (df *DataFrame) IsFuzzyEqual(a float64, b float64) bool {
	var e float64

	if a >= 100000.0 {
		e = 1000.0
	} else if a >= 10000.0 {
		e = 100.0
	} else if a >= 1000.0 {
		e = 10.0
	} else if a >= 100.0 {
		e = 1.0
	} else if a >= 10.0 {
		e = 0.1
	} else if a >= 1.0 {
		e = 0.01
	} else if a >= 0.1 {
		e = 0.001
	} else if a >= 0.01 {
		e = 0.0001
	} else if a >= 0.001 {
		e = 0.00001
	} else if a >= 0.0001 {
		e = 0.000001
	} else {
		e = 0.0000001
	}

	return math.Abs(a-b) <= e
}

// IsBlack ...
func (df *DataFrame) IsBlack() bool {
	return df.Last.Open > df.Last.Close
}

// IsWhite ...
func (df *DataFrame) IsWhite() bool {
	return df.Last.Open < df.Last.Close
}

// IsDownTrend ...
func (df *DataFrame) IsDownTrend() bool {
	var blackCount int

	if df.LastIndex-3 < 0 {
		return false
	}

	if df.Series[:df.LastIndex-1].Next().IsBlack() {
		blackCount++
	}
	if df.Series[:df.LastIndex-2].Next().IsBlack() {
		blackCount++
	}
	if df.Series[:df.LastIndex-3].Next().IsBlack() {
		blackCount++
	}

	if blackCount > 1 {
		return true
	}

	return false
}

// IsUpTrend ...
func (df *DataFrame) IsUpTrend() bool {
	var whiteCount int

	if df.LastIndex-3 < 0 {
		return false
	}

	if df.Series[:df.LastIndex-1].Next().IsWhite() {
		whiteCount++
	}
	if df.Series[:df.LastIndex-2].Next().IsWhite() {
		whiteCount++
	}
	if df.Series[:df.LastIndex-3].Next().IsWhite() {
		whiteCount++
	}

	if whiteCount > 1 {
		return true
	}

	return false
}

// IsThreeBlackCrows chart pattern is the opposite of the three white soldiers chart pattern.
// Instead of three bullish candles with the three white soldiers, you have three bearish candles instead.
// Also, the three black crows pattern neeDataSet to come after an extended uptrend and consolidation for it to confirm a new downtrend.
func (df *DataFrame) IsThreeBlackCrows() bool {
	var retVal bool
	var cond [13]bool

	if df.LastIndex-2 < 0 {
		return false
	}

	cond[0] = df.Last.Open > df.Last.Close
	cond[1] = df.Last.High > df.Last.Open
	cond[2] = df.Last.Low < df.Last.Close

	retVal = true
	for c := 0; c < 3; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[3] = df.Prev.Open > df.Prev.Close
	cond[4] = df.Prev.High > df.Prev.Open
	cond[5] = df.Prev.Low < df.Prev.Close

	retVal = true
	for c := 3; c < 6; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[6] = df.Series[df.LastIndex-2].Open > df.Series[df.LastIndex-2].Close
	cond[7] = df.Series[df.LastIndex-2].High > df.Series[df.LastIndex-2].Open
	cond[8] = df.Series[df.LastIndex-2].Low < df.Series[df.LastIndex-2].Close

	retVal = true
	for c := 6; c < 9; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[9] = df.Last.Open < df.Prev.Open
	cond[10] = df.Prev.Open < df.Series[df.LastIndex-2].Open
	cond[11] = df.Last.Close < df.Prev.Close
	cond[12] = df.Prev.Close < df.Series[df.LastIndex-2].Close

	retVal = true
	for c := 9; c < 13; c++ {
		retVal = retVal && cond[c]
	}

	return retVal
}

// IsThreeWhiteSoldiers pattern can appear after an extended downtrend and a period of consolidation.
// The first candlestick of the chart pattern that neeDataSet to appear is a bullish candlestick with a long body.
// The next candlestick in the pattern is another bullish candlestick, but this candlestick neeDataSet to have a body of greater size than the first candlestick.
// This second candlestick also neeDataSet to have little to no shadow. The last candlestick is another bullish candlestick that neeDataSet to be equal or greater length of a body than the second candlestick.
// When all three candlesticks appear, this chart pattern can be used to confirm the start of a new uptrend.
func (df *DataFrame) IsThreeWhiteSoldiers() bool {
	if df.LastIndex-2 < 0 {
		return false
	}

	if df.IsWhite() && df.Series[:df.LastIndex-1].Next().IsWhite() && df.Series[:df.LastIndex-2].Next().IsWhite() == false {
		return false
	}

	if df.IsFuzzyEqual(df.Last.Open, df.Last.Low) && df.IsFuzzyEqual(df.Last.Close, df.Last.High) == false {
		return false
	}

	if df.IsFuzzyEqual(df.Prev.Open, df.Prev.Low) && df.IsFuzzyEqual(df.Prev.Close, df.Prev.High) == false {
		return false
	}

	if df.IsFuzzyEqual(df.Series[df.LastIndex-2].Open, df.Series[df.LastIndex-2].Low) && df.IsFuzzyEqual(df.Series[df.LastIndex-2].Close, df.Series[df.LastIndex-2].High) == false {
		return false
	}

	if math.Abs(df.Last.Open-df.Last.Close) > math.Abs(df.Prev.Open-df.Prev.Close) == false {
		return false
	}

	if math.Abs(df.Prev.Open-df.Prev.Close) > math.Abs(df.Series[df.LastIndex-2].Open-df.Series[df.LastIndex-2].Close) == false {
		return false
	}

	if df.Last.Close > df.Prev.Close && df.Prev.Close > df.Series[df.LastIndex-2].Close == false {
		return false
	}

	if df.Last.Open > df.Prev.Open && df.Prev.Open > df.Series[df.LastIndex-2].Open == false {
		return false
	}

	return true
}

// IsBullishBeltHold is a bullish belt-hold line is a tall white candle that has very little or no lower shadow and little or no upper shadow
// In a downtrend, the Low becomes the support. If you are short, then time to take profit
func (df *DataFrame) IsBullishBeltHold() bool {
	if df.IsWhite() == false {
		return false
	}

	if df.IsFuzzyEqual(df.Last.Open, df.Last.Low) {
		if df.IsFuzzyEqual(df.Last.Close, df.Last.High) {
			return true
		} else {
			realBody := math.Abs(df.Last.Open - df.Last.Close)
			upperShadow := math.Abs(df.Last.Close - df.Last.High)

			if 4.0*upperShadow <= realBody {
				return true
			}
		}
	}

	return false
}

// IsBearishBeltHold is a bearish belt-hold line is a long, black real body that opens at the high of the session and closes at or near the low of the session.
// It has small or nonexistent upper or lower shadows.
// In an uptrend, the High becomes the resistance. If you are Long, then time to take profit
func (df *DataFrame) IsBearishBeltHold() bool {
	if df.IsBlack() == false {
		return false
	}

	if df.IsFuzzyEqual(df.Last.Open, df.Last.High) {
		if df.IsFuzzyEqual(df.Last.Close, df.Last.Low) {
			return true
		} else {
			realBody := math.Abs(df.Last.Open - df.Last.Close)
			lowerShadow := math.Abs(df.Last.Close - df.Last.Low)

			if 4.0*lowerShadow <= realBody {
				return true
			}
		}
	}

	return false
}

// IsDoji candlestick where the opening price is almost the exact same as the opening price, with long shadows in one direction or both.
// What this can signal is indecision between buyers and sellers.
// If these occur at the top or bottom of a trend it can signal a reversal as it shows a slowing of momentum.
func (df *DataFrame) IsDoji() bool {
	return df.IsFuzzyEqual(df.Last.Open, df.Last.Close)
}

// IsSpinningTop has two long equal length shadows with a small body and typically signals a reversal when they occur during a trend.
// The reason behind the reversal is that it shows indecision between buyers and sellers, and that neither of them can close much higher or lower than the opening.
func (df *DataFrame) IsSpinningTop() bool {
	body := math.Abs(df.Last.Open - df.Last.Close)
	height := math.Abs(df.Last.High - df.Last.Low)

	if body <= (height / 3.0) {
		return true
	}
	return false
}

// IsMarubozu candlestick forms with a long body and little to no shadow.
// This signals strong movement in one direction, which will likely continue movement in that direction in the near future.
// In the bullish Marubozu case, the opening price is equal to the low and the closing price is equal to the high.
// With a bearish Marubozu the opening price is the high and the closing is the low.
func (df *DataFrame) IsMarubozu() bool {
	if df.IsFuzzyEqual(df.Last.Open, df.Last.Low) && df.IsFuzzyEqual(df.Last.Close, df.Last.High) {
		return true
	}

	return false
}

// IsHammer chart pattern is a Japanese candlestick that has a small body with a short to no shadow on top of the body with a long shadow on the bottom.
// When this candlestick occurs at the bottom of a trend, it can signal for a reversal.
func (df *DataFrame) IsHammer() bool {
	var diffHighLow float64
	var diffOpenClose float64

	if df.LastIndex < 0 {
		return false
	}

	if df.IsFuzzyEqual(df.Last.Close, df.Last.High) || df.IsFuzzyEqual(df.Last.Open, df.Last.High) == false {
		return false
	}

	diffOpenClose = math.Abs(df.Last.Open - df.Last.Close)
	diffHighLow = math.Abs(df.Last.High - df.Last.Low)
	if diffHighLow/2.0 >= diffOpenClose {
		return true
	}

	return false
}

// IsHangingMan candlestick pattern has the exact same candlestick as the hammer but has different price action before it, so it signals for a reversal downwards
func (df *DataFrame) IsHangingMan() bool {
	return df.IsHammer()
}

// IsInvertedHammer is a candlestick similar to the hammer and hanging man patterns in that it can signal a reversal.
// With an inverted hammer, a small bullish candlestick body forms with a long shadow on top, and occurs during a downtrend.
func (df *DataFrame) IsInvertedHammer() bool {
	return df.IsShootingStar()
}

// IsShootingStar similar to the inverted hammer but occurs during an uptrend and can signal a reversal downward DataSet.
// The candlestick for a shooting star is a small bearish body with a long shadow on top.
func (df *DataFrame) IsShootingStar() bool {
	var diffHighLow float64
	var diffOpenClose float64

	if df.LastIndex < 0 {
		return false
	}

	if df.IsFuzzyEqual(df.Last.Close, df.Last.Low) || df.IsFuzzyEqual(df.Last.Open, df.Last.Low) == false {
		return false
	}

	diffOpenClose = math.Abs(df.Last.Open - df.Last.Close)
	diffHighLow = math.Abs(df.Last.High - df.Last.Low)
	if diffHighLow/2.0 >= diffOpenClose {
		return true
	}

	return false
}

// IsEveningStar similar to the morning star pattern but occurs during an uptrend and signals a reversal downward DataSet.
// The evening stars' first candle is a bullish candle with a long body.
// The second candle is a doji, which signals indecision.
// The third and final candle in the chart pattern is the bearish candle that closes past at least the halfway point of the first bullish candle.
func (df *DataFrame) IsEveningStar() bool {
	var retVal bool
	var cond [11]bool

	if df.LastIndex-2 < 0 {
		return false
	}

	cond[0] = df.Last.Open > df.Last.Close
	cond[1] = df.Last.High > df.Last.Open
	cond[2] = df.Last.Low < df.Last.Close

	retVal = true
	for c := 0; c < 3; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[3] = df.Prev.Open > df.Prev.Close
	cond[4] = df.Prev.High > df.Prev.Open
	cond[5] = df.Prev.Low < df.Prev.Close

	retVal = true
	for c := 3; c < 6; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[6] = df.Series[df.LastIndex-2].Open < df.Series[df.LastIndex-2].Close
	cond[7] = df.Series[df.LastIndex-2].High > df.Series[df.LastIndex-2].Close
	cond[8] = df.Series[df.LastIndex-2].Low < df.Series[df.LastIndex-2].Open

	retVal = true
	for c := 6; c < 9; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[9] = df.Prev.Close > df.Series[df.LastIndex-2].Close
	cond[10] = df.Prev.Close > df.Last.Open

	retVal = true
	for c := 9; c < 11; c++ {

		retVal = retVal && cond[c]
	}

	return retVal
}

// IsMorningStar is a bearish candle with a long body.
// It is then followed by a doji (a small body candle with long shadows on bottom and top).
// The doji signals indecision and doesn't matter if it closes up or down.
// The third candlestick is a bullish candlestick that should at least pass the halfway point of the first bearish candle.
// The morning star is a buy indicator.
func (df *DataFrame) IsMorningStar() bool {
	var retVal bool
	var cond [11]bool

	if df.LastIndex-2 < 0 {
		return false
	}

	cond[0] = df.Last.Open < df.Last.Close
	cond[1] = df.Last.High > df.Last.Close
	cond[2] = df.Last.Low < df.Last.Open

	retVal = true
	for c := 0; c < 3; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[3] = df.Prev.Open < df.Prev.Close
	cond[4] = df.Prev.High > df.Prev.Close
	cond[5] = df.Prev.Low < df.Prev.Open

	retVal = true
	for c := 3; c < 6; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[6] = df.Series[df.LastIndex-2].Open > df.Series[df.LastIndex-2].Close
	cond[7] = df.Series[df.LastIndex-2].High > df.Series[df.LastIndex-2].Open
	cond[8] = df.Series[df.LastIndex-2].Low < df.Series[df.LastIndex-2].Close

	retVal = true
	for c := 6; c < 9; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[9] = df.Prev.Close < df.Series[df.LastIndex-2].Close
	cond[10] = df.Prev.Close < df.Last.Open

	retVal = true
	for c := 9; c < 11; c++ {
		retVal = retVal && cond[c]
	}

	return retVal
}

// IsInsideDown pattern is the opposite of the three inside up pattern.
// In this case, the pattern is an indicator for a reversal downward DataSet and must follow a recent uptrend.
// The first candlestick in the pattern is a bullish candle with a long body.
// The second is a bearish candle that passes at least the halfway point of the first bullish candle.
// The last candlestick is another bearish candle that passes at least the low of the first bullish candle.
func (df *DataFrame) IsInsideDown() bool {
	var retVal bool
	var cond [12]bool

	if df.LastIndex-2 < 0 {
		return false
	}

	cond[0] = df.Last.Open > df.Last.Close
	cond[1] = df.Last.High > df.Last.Open
	cond[2] = df.Last.Low < df.Last.Close

	retVal = true
	for c := 0; c < 3; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[3] = df.Prev.Open > df.Prev.Close
	cond[4] = df.Prev.High > df.Prev.Open
	cond[5] = df.Prev.Low < df.Prev.Close

	retVal = true
	for c := 3; c < 6; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[6] = df.Series[df.LastIndex-2].Open < df.Series[df.LastIndex-2].Close
	cond[7] = df.Series[df.LastIndex-2].High > df.Series[df.LastIndex-2].Close
	cond[8] = df.Series[df.LastIndex-2].Low < df.Series[df.LastIndex-2].Open

	retVal = true
	for c := 6; c < 9; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[9] = df.Last.Open < df.Prev.Open
	cond[10] = df.Prev.Open < df.Series[df.LastIndex-2].Close
	cond[11] = df.Prev.Close > df.Series[df.LastIndex-2].Open

	retVal = true
	for c := 9; c < 12; c++ {
		retVal = retVal && cond[c]
	}

	return retVal
}

// IsInsideUp pattern occurs after a recent downtrend and signals for a reversal to an uptrend.
// The first candle in the pattern is a bearish candle with a long body.
// The next is a bullish candle that passes at least the halfway point of the first bearish candle.
// The third and final candle is another bullish candle that passes at least the high of the first bearish candle.
func (df *DataFrame) IsInsideUp() bool {
	var retVal bool
	var cond [12]bool

	if df.LastIndex-2 < 0 {
		return false
	}

	cond[0] = df.Last.Open < df.Last.Close
	cond[1] = df.Last.High > df.Last.Close
	cond[2] = df.Last.Low < df.Last.Open

	retVal = true
	for c := 0; c < 3; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[3] = df.Prev.Open < df.Prev.Close
	cond[4] = df.Prev.High > df.Prev.Close
	cond[5] = df.Prev.Low < df.Prev.Open

	retVal = true
	for c := 3; c < 6; c++ {
		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[6] = df.Series[df.LastIndex-2].Open > df.Series[df.LastIndex-2].Close
	cond[7] = df.Series[df.LastIndex-2].High > df.Series[df.LastIndex-2].Open
	cond[8] = df.Series[df.LastIndex-2].Low < df.Series[df.LastIndex-2].Close

	retVal = true
	for c := 6; c < 9; c++ {

		retVal = retVal && cond[c]
	}

	if retVal == false {
		return false
	}

	cond[9] = df.Last.Open > df.Prev.Open
	cond[10] = df.Prev.Open > df.Series[df.LastIndex-2].Close
	cond[11] = df.Prev.Close < df.Series[df.LastIndex-2].Open

	retVal = true
	for c := 9; c < 12; c++ {
		retVal = retVal && cond[c]
	}

	return retVal
}
