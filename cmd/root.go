package cmd

import (
	"context"
	"fmt"
	"github.com/harunnryd/btrade/cmd/worker"
	"os"

	"github.com/harunnryd/btrade/cmd/root"
	"github.com/harunnryd/btrade/internal/app/appcontext"
	"github.com/harunnryd/btrade/internal/app/apperror"
	"github.com/harunnryd/btrade/internal/pkg/utils/atexit"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   root.AppName,
	Short: root.AppDescShort,
	Long:  root.AppDescLong,
	Run: func(cmd *cobra.Command, args []string) {
		root.Start()
	},
}

// Execute ...
func Execute(ctx context.Context, ctxCancel context.CancelFunc, eg *errgroup.Group, logger zerolog.Logger) {
	root.App = appcontext.NewAppContext(ctx, ctxCancel, eg)
	root.App.InitAppErrors(apperror.AppErrors...)

	_ = root.App.InitLogger(logger)

	atexit.Add(root.App.CtxCancel)

	cobra.OnInitialize(root.InitConfig)
	cobra.OnInitialize(root.InitDependencies)

	worker.App = appcontext.NewAppContext(ctx, ctxCancel, eg)
	worker.App.InitAppErrors(apperror.AppErrors...)

	_ = worker.App.InitLogger(logger)

	atexit.Add(worker.App.CtxCancel)

	cobra.OnInitialize(worker.InitConfig)
	cobra.OnInitialize(worker.InitDependencies)

	atexit.Add(root.Shutdown)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
