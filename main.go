package main

import (
	"context"
	"os"
	"time"

	"github.com/harunnryd/btrade/cmd"
	"github.com/harunnryd/btrade/internal/pkg/utils/atexit"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"
	"github.com/harunnryd/btrade/internal/pkg/utils/shutdown"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"golang.org/x/sync/errgroup"
)

const (
	// AppName ...
	AppName string = "btrade"

	// AppDescShort ...
	AppDescShort string = "bot auto trade"

	// AppDescLong ...
	AppDescLong string = `bot auto trade in binary option`
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	var gracefulShutdown bool

	defer func() {
		if !gracefulShutdown {
			atexit.AtExit()
		}
	}()

	startTime := "start_time"
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, startTime, time.Now())

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Stack().
		Str("service", AppName).
		Logger()

	ctx = logger.WithContext(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	sigTrap := shutdown.TermSignalTrap()
	cmd.Execute(ctx, cancel, eg, logger)

	eg.Go(func() error {
		return sigTrap.Wait(ctx)
	})

	loghelper.Msg(ctx, "system started")

	if err := eg.Wait(); err != nil && err != shutdown.ErrTermSig {
		loghelper.Panic(ctx, "failed to wait goroutine group", err)
	}

	atexit.AtExit()
	gracefulShutdown = true
	loghelper.Msg(context.Background(), "graceful shutdown successfully finished")
}
