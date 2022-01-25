package analysis

import (
	"fmt"
	"github.com/gocraft/work"
	"github.com/harunnryd/btrade/internal/app/appcontext"
	"github.com/harunnryd/btrade/internal/app/engine"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"
	"github.com/robfig/cron"
	"time"
)

type scheduler struct {
	CronName string
	App      *appcontext.AppContext
	Pool     *work.WorkerPool
}

// NewScheduler ...
func NewScheduler(app *appcontext.AppContext, pool *work.WorkerPool) Scheduler {
	return &scheduler{
		CronName: "btrade-analysis",
		App:      app,
		Pool:     pool,
	}
}

// StartJob ...
func (sch *scheduler) StartJob(_ *work.Job) error {
	e := engine.NewEngine(sch.App)
	e.Olymptrade.StartCandleStream(sch.App.Ctx)

	return nil
}

// Run ...
func (sch *scheduler) Run() {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := parser.Parse("0 * * * * *")
	if err != nil {
		loghelper.AddErr(sch.App.Ctx, fmt.Errorf("[%s] %v", sch.CronName, err))
		panic(err)
	}

	loghelper.Msg(sch.App.Ctx, fmt.Sprintf("RUN SCHEDULER %s : [%s]", sch.CronName, schedule.Next(time.Now())))

	sch.Pool.PeriodicallyEnqueue("0 * * * * *", sch.CronName)
	sch.Pool.JobWithOptions(sch.CronName, work.JobOptions{MaxConcurrency: 10}, sch.StartJob)
}
