package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:           "myesa",
	Short:         "esa client",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		searchCmd(),
		editCmd(),
	)
}
