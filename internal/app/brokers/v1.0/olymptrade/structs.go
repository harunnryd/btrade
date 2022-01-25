package olymptrade

import (
	"github.com/chuckpreslar/emission"
	"github.com/sacOO7/gowebsocket"
)

type Options struct {
	websocket   gowebsocket.Socket
	accountMode string
	accountID   string
	pair        string

	emitter *emission.Emitter
}

type DataFrame []struct {
	Frames []Frame `json:"d"`
}

type Frame struct {
	Pair      string   `json:"p"`
	Timeframe int      `json:"tf"`
	Candles   []Candle `json:"candles"`
}

type Candle struct {
	Time  int64   `json:"t"`
	Open  float64 `json:"open"`
	Low   float64 `json:"low"`
	High  float64 `json:"high"`
	Close float64 `json:"close"`
}
