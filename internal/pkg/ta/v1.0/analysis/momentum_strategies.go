package analysis

import (
	"fmt"

	talibcdl "github.com/iwat/talib-cdl-go"
)

// RsiStrategy sells above sell at, buys below buy at.
func RsiStrategy(df *DataFrame, sellAt, buyAt float64) []Action {
	actions := make([]Action, len(df.Series))

	_, rsi := df.Rsi()

	for i := 0; i < len(actions); i++ {
		if rsi[i] <= buyAt {
			actions[i] = BUY
		} else if rsi[i] >= sellAt {
			actions[i] = SELL
		} else {
			actions[i] = HOLD
		}
	}

	return actions
}

func Summaries(df *DataFrame) (actions, summaries []string) {
	highs := make([]float64, len(df.Series))
	opens := make([]float64, len(df.Series))
	closes := make([]float64, len(df.Series))
	lows := make([]float64, len(df.Series))

	for i := 0; i < len(df.Series); i++ {
		highs[i] = df.Series[i].High
		opens[i] = df.Series[i].Open
		closes[i] = df.Series[i].Close
		lows[i] = df.Series[i].Low
	}

	ndf := talibcdl.SimpleSeries{
		Highs:  highs,
		Opens:  opens,
		Closes: closes,
		Lows:   lows,
	}

	var candles []int

	if candles = talibcdl.ThreeBlackCrows(ndf); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("ThreeBlackCrows Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.ThreeInside(ndf); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("ThreeInside Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	} else if candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ThreeInside Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.EveningStar(ndf, talibcdl.DefaultFloat64); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("EveningStar: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.ThreeStarsInSouth(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ThreeStarsInSouth Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.ThreeWhiteSoldiers(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ThreeWhiteSoldiers Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.ThreeLineStrike(ndf); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("ThreeLineStrike Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	} else if candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ThreeLineStrike Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.BeltHold(ndf); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("BeltHold Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	} else if candles[len(candles)-1] == 100{
		summaries = append(summaries, fmt.Sprintf("BeltHold Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.ConcealBabySwall(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ConcealBabySwall Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.AbandonedBaby(ndf, talibcdl.DefaultFloat64); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("AbandonedBaby Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	} else if candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("AbandonedBaby Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.AdvanceBlock(ndf); candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("AdvanceBlock Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.ThreeOutside(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ThreeOutside Bullish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	} else if candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("ThreeOutside Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.StickSandwich(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("StickSandwich Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if candles = talibcdl.Piercing(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("Piercing Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.MatchingLow(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("MatchingLow Bulish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	}

	if candles = talibcdl.ClosingMarubozu(ndf); candles[len(candles)-1] == 100 {
		summaries = append(summaries, fmt.Sprintf("ClosingMarubozu Bulish: %d", candles[len(candles)-1]))
		actions = append(actions, "BUY")
	} else if candles[len(candles)-1] == -100 {
		summaries = append(summaries, fmt.Sprintf("ClosingMarubozu Bearish: %d", candles[len(candles)-1]))
		actions = append(actions, "SELL")
	}

	if df.IsSmaShowDownTrend() {
		summaries = append(summaries, "IsDownTrend")
	} else if df.IsSmaShowUpTrend() {
		summaries = append(summaries, "IsUpTrend")
	} else {
		summaries = append(summaries, "IsSideways")
	}

	return
}

func CandlestickChartStrategy(df *DataFrame) []Action {
	var actions = []Action{0}

	highs := make([]float64, len(df.Series))
	opens := make([]float64, len(df.Series))
	closes := make([]float64, len(df.Series))
	lows := make([]float64, len(df.Series))

	for i := 0; i < len(df.Series); i++ {
		highs[i] = df.Series[i].High
		opens[i] = df.Series[i].Open
		closes[i] = df.Series[i].Close
		lows[i] = df.Series[i].Low
	}

	ndf := talibcdl.SimpleSeries{
		Highs:  highs,
		Opens:  opens,
		Closes: closes,
		Lows:   lows,
	}

	var candles []int

	if candles = talibcdl.ThreeBlackCrows(ndf); candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.ThreeInside(ndf); candles[len(candles)-1] == -100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == 100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.EveningStar(ndf, talibcdl.DefaultFloat64); candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.ThreeStarsInSouth(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.ThreeWhiteSoldiers(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.ThreeLineStrike(ndf); candles[len(candles)-1] == -100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == 100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.BeltHold(ndf); candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	} else if candles[len(candles)-1] == 100 && df.IsSmaShowUpTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == -100 && df.IsSmaShowDownTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.ConcealBabySwall(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.AbandonedBaby(ndf, talibcdl.DefaultFloat64); candles[len(candles)-1] == 100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	} else if candles[len(candles)-1] == -100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.AdvanceBlock(ndf); candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.ThreeOutside(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	} else if candles[len(candles)-1] == -100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.StickSandwich(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, SELL)
	}

	if candles = talibcdl.Piercing(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.MatchingLow(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	}

	if candles = talibcdl.ClosingMarubozu(ndf); candles[len(candles)-1] == 100 && df.IsSmaShowDownTrend() {
		actions = append(actions, BUY)
	} else if candles[len(candles)-1] == -100 && df.IsSmaShowUpTrend() {
		actions = append(actions, SELL)
	}

	if len(actions) < 1 {
		actions = append(actions, 0)
	}

	return actions
}

// DefaultRsiStrategy it buys below 30 and sells above 70.
func DefaultRsiStrategy(df *DataFrame) []Action {
	return RsiStrategy(df, 70, 30)
}
