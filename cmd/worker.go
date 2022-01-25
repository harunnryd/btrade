package cmd

import (
	"github.com/harunnryd/btrade/cmd/root"
	"github.com/harunnryd/btrade/cmd/worker"
	"github.com/harunnryd/btrade/internal/pkg/utils/atexit"
	"github.com/spf13/cobra"
)

// Command represents the worker command
var cmdWorker = &cobra.Command{
	Use:              worker.CmdUse,
	Aliases:          worker.CmdAliases,
	TraverseChildren: true,
	Short:            "Worker of " + root.AppName,
	Long:             `Worker of ` + root.AppName,
	Run: func(cmd *cobra.Command, args []string) {
		atexit.Add(worker.Shutdown)
		worker.Start()
	},
}

func init() {
	rootCmd.AddCommand(cmdWorker)
}
