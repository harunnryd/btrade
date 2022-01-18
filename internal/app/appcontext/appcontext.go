package appcontext

import (
	"context"

	"github.com/harunnryd/btrade/internal/app/config"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

// NewAppContext ...
func NewAppContext(ctx context.Context, ctxCancel context.CancelFunc, eg *errgroup.Group) *AppContext {
	return &AppContext{
		Ctx:       ctx,
		CtxCancel: ctxCancel,
		Eg:        eg,
	}
}

// InitConfig ...
func (a *AppContext) InitConfig(cfg config.Config) (err error) {
	a.Config = cfg
	a.IsLive = cfg.App.Environment == "live"

	return
}

// InitAppErrors ...
func (a *AppContext) InitAppErrors(appErr ...error) {
	a.AppErrors = append(a.AppErrors, appErr...)
}

// InitLogger ...
func (a *AppContext) InitLogger(logger zerolog.Logger) (err error) {
	a.Logger = logger
	return
}
