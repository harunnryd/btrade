package cmd

import (
	"github.com/harunnryd/btrade/cmd/root"
	"github.com/harunnryd/btrade/cmd/version"
	"github.com/spf13/cobra"
)

// Command represents the version command
var cmdVersion = &cobra.Command{
	Use:              version.CmdUse,
	Aliases:          version.CmdAliases,
	TraverseChildren: true,
	Short:            "Print the version number of " + root.AppName,
	Long:             `All software has versions. This is ` + root.AppName,
	Run: func(cmd *cobra.Command, args []string) {
		version.Start()
	},
}

func init() {
	rootCmd.AddCommand(cmdVersion)
}
