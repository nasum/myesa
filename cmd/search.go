package cmd

import (
	"net/url"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
	"github.com/upamune/go-esa/esa"
)

func searchCmd(client esa.Client, team string) *cobra.Command {
	displayConsole := &lib.DisplayConsole{}

	cmd := &cobra.Command{
		Use:   "search",
		Short: "search articles",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			query := url.Values{}
			query.Add("q", args[0])

			response, error := client.Post.GetPosts(team, query)

			if error != nil {
				return error
			}

			return displayConsole.ShowArticles(response.Posts)
		},
	}

	return cmd
}
