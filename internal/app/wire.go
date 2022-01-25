package app

import (
	"github.com/gocraft/work"
	"github.com/harunnryd/btrade/internal/app/appcontext"
	"github.com/harunnryd/btrade/internal/app/scheduler"
	"github.com/harunnryd/btrade/internal/app/scheduler/analysis"
)

// WiringScheduler ...
func WiringScheduler(appCtx *appcontext.AppContext, pool *work.WorkerPool) scheduler.Schedulers {
	return scheduler.Schedulers{
		Analysis: analysis.NewScheduler(appCtx, pool),
	}
}
