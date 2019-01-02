package cmd

import (
	"fmt"
	"net/url"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
)

// SearchParams is search paramater
type SearchParams struct {
	Count          int
	IncludeComment string
	Debug          bool
}

func searchCmd() *cobra.Command {
	displayConsole := &lib.DisplayConsole{}
	searchParams := &SearchParams{}
	esaClient := &lib.EsaClient{}

	esaClient.Init()

	cmd := &cobra.Command{
		Use:   "search",
		Short: "search articles",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			query := url.Values{}
			query.Add("", args[0])
			if searchParams.IncludeComment != "" {
				query.Add("comment", searchParams.IncludeComment)
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
	flags.StringVarP(&searchParams.IncludeComment, "include-comment", "i", "", "include comment")
	flags.BoolVarP(&searchParams.Debug, "debug", "d", false, "show request url")

	return cmd
}
