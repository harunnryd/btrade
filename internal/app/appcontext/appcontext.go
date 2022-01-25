package appcontext

import (
	"context"
	driverOlymptrade "github.com/harunnryd/btrade/internal/pkg/drivers/v1.0/olymptrade"
	driverRedis "github.com/harunnryd/btrade/internal/pkg/drivers/v1.0/redis"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"

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

// InitRedisPool ...
func (a *AppContext) InitRedisPool() (err error) {
	if a.Config.RedisPool.IsEnabled {
		a.RedisPool = driverRedis.NewRedis(a.Config.RedisPool.Options())
	}
	return
}

// InitOlymptrade ...
func (a *AppContext) InitOlymptrade() (err error) {
	if a.Config.Olymptrade.IsEnabled {
		a.OlymptradeOnce.Do(func() {
			if a.Olymptrade, err = driverOlymptrade.NewOlymptrade(
				a.Config.Olymptrade.Options(),
			); err != nil {
				loghelper.AddErr(a.Ctx, err)
			}
		})
	}

	return
}
