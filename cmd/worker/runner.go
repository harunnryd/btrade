package worker

import (
	"fmt"
	"github.com/gocraft/work"
	"github.com/harunnryd/btrade/internal/app"
	"github.com/harunnryd/btrade/internal/app/config"
	uConfig "github.com/harunnryd/btrade/internal/pkg/utils/config"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

// InitConfig ...
func InitConfig() {
	_ = godotenv.Load("params/.env")
	var cfg config.Config

	if err := uConfig.LoadConfigs(ConfigPaths[:], &cfg); err != nil {
		loghelper.Panic(App.Ctx, "failed to init config", err)
	}

	_ = App.InitConfig(cfg)
}

// InitDependencies ...
func InitDependencies() {
	_ = App.InitOlymptrade()
	_ = App.InitRedisPool()
}

// CronContext defines worker context
type CronContext struct{}

// Start ...
func Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	pool := work.NewWorkerPool(CronContext{}, 1, "btrade", App.RedisPool)
	scheduler := app.WiringScheduler(App, pool)
	scheduler.Analysis.Run()

	go func() {
		pool.Start()
	}()

	for {
		select {
		case <-interrupt:
			loghelper.AddStr(App.Ctx, "interrupt")
			pool.Stop()
			return
		}
	}
}

// Shutdown ...
func Shutdown() {
	fmt.Println("\n-#-")
}
