package cmd

import (
	"fmt"
	"net/url"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
)

// SearchParams is search paramater
type SearchParams struct {
	Count       int
	Interactive bool
	Debug       bool
}

func searchCmd() *cobra.Command {
	displayConsole := &lib.DisplayConsole{}
	searchParams := &SearchParams{}
	esaClient := &lib.EsaClient{}

	esaClient.Init()

	cmd := &cobra.Command{
		Use:   "search",
		Short: "search articles",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			query := url.Values{}
			query.Add("name", args[0])
			if searchParams.Interactive {
				lib.Interactive()
			}

			if searchParams.Debug {
				fmt.Println(query.Encode())
			}

			articles := esaClient.Search(query)

			for _, value := range articles {
				displayConsole.ShowArticles(value)
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&searchParams.Interactive, "interactive", "i", false, "interactive")
	flags.BoolVarP(&searchParams.Debug, "debug", "d", false, "show request url")

	return cmd
}
