package olymptrade

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/google/uuid"
	"github.com/harunnryd/btrade/internal/pkg/ta/v1.0/analysis"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"
	"github.com/sacOO7/gowebsocket"
)

// Service ...
type Service struct {
	options Options
}

// SetOptions ...
func (s *Service) SetOptions(options Options) {
	s.options = options
}

// StartCandleStream ...
func (s *Service) StartCandleStream(ctx context.Context) {
	s.options.websocket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		loghelper.AddErrAndStr(ctx, "Received connect error", err)
		os.Exit(1)
	}

	s.options.websocket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		var df DataFrame
		if err := json.Unmarshal(
			[]byte(message),
			&df,
		); err != nil {
			return
		}

		s.options.emitter.Emit("df", ctx, df)

		s.options.websocket.Close()
	}

	s.options.websocket.Connect()

	s.options.websocket.SendText(s.buildMessageCandleStream())
}

// Analysis ...
func (s *Service) Analysis(ctx context.Context, df DataFrame) {
	frames := df[0].Frames[0].Candles

	ndf := analysis.NewDataFrame()

	for i := 1; i < len(frames); i++ {
		ndf.Add(analysis.Asset{
			Time:  time.Unix(frames[i].Time, 0),
			High:  frames[i].High,
			Low:   frames[i].Low,
			Open:  frames[i].Open,
			Close: frames[i].Close,
		})
	}

	ndf.Sort()

	directions := analysis.CandlestickChartStrategy(ndf)

	_, summaries := analysis.Summaries(ndf)

	dir := directions[len(directions)-1]

	s.consoleLog(s.options.pair, strings.ToUpper(s.convertDirection(dir)), directions, summaries)

	s.options.emitter.Emit("act", ctx, s.convertDirection(dir))
}

func (s *Service) consoleLog(pair string, dir string, directions []analysis.Action, summaries []string) {
	table := simpletable.New()

	r := []*simpletable.Cell{
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", blue(time.Now().Format("2006-01-02 15:04:05")))},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", green(pair))},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", red(dir))},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", directions)},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", strings.Join(summaries, ","))},
	}

	table.Body.Cells = append(table.Body.Cells, r)

	table.SetStyle(simpletable.StyleDefault)
	table.Println()
}

// Buy ...
func (s *Service) Buy(ctx context.Context, dir string) {
	s.options.websocket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		loghelper.AddErrAndStr(ctx, "Received connect error", err)
		os.Exit(1)
	}

	s.options.websocket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		s.options.websocket.Close()
	}

	s.options.websocket.Connect()

	s.options.websocket.SendText(s.buildMessageBuy(dir))
}

func (s *Service) buildMessageCandleStream() string {
	return fmt.Sprintf(messageCandleStream, s.generateUID(), s.options.pair, time.Now().Add(1*time.Minute).Unix())
}

func (s *Service) buildMessageBuy(dir string) string {
	return fmt.Sprintf(messageBuy, s.generateUID(), 100, dir, s.options.pair, s.options.accountID, s.options.accountMode, time.Now().Add(1*time.Minute).Unix())
}

func (s *Service) generateUID() string {
	return strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1))[:19]
}

func (s *Service) convertDirection(dir analysis.Action) string {
	switch dir {
	case analysis.BUY:
		return "up"
	case analysis.SELL:
		return "down"
	default:
		return "hold"
	}
}
