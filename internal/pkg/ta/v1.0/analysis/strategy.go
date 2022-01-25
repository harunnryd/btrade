package analysis

import (
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
