package engine

import (
	"github.com/chuckpreslar/emission"
	"github.com/harunnryd/btrade/internal/app/appcontext"
	"github.com/harunnryd/btrade/internal/app/brokers/v1.0/olymptrade"
)

// Engine ...
type Engine struct {
	App        *appcontext.AppContext
	Olymptrade olymptrade.WebsocketServiceClient
}

// NewEngine ...
func NewEngine(app *appcontext.AppContext) *Engine {
	builder := olymptrade.Builder{}
	svc := &olymptrade.Service{}
	builder.SetEmitter(
		emission.NewEmitter().
			On("df", svc.Analysis).
			On("act", svc.Buy),
	)
	builder.SetBuilder(svc)
	builder.SetPair(app.Config.Olymptrade.Pair)
	builder.SetAccountID(app.Config.Olymptrade.AccountID)
	builder.SetAccountMode(app.Config.Olymptrade.AccountMode)
	builder.SetWebsocket(app.Olymptrade)

	return &Engine{
		Olymptrade: builder.Build(),
		App:        app,
	}
}
