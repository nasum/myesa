package cmd

import (
	"fmt"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
)

func showCmd() *cobra.Command {
	esaClient := &lib.EsaClient{}
	esaClient.Init()

	cmd := &cobra.Command{
		Use:   "show",
		Short: "show articles",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			articles := esaClient.GetArticlesByName(args[0])
			fmt.Println(articles[0].BodyMd)
			return nil
		},
	}

	return cmd
}
