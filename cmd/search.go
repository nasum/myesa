package cmd

import (
	"fmt"
	"net/url"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
	"github.com/upamune/go-esa/esa"
)

type SearchParams struct {
	Count          int
	IncludeComment string
	Debug          bool
}

func searchCmd(client esa.Client, team string) *cobra.Command {
	displayConsole := &lib.DisplayConsole{}
	searchParams := &SearchParams{}

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
			if searchParams.IncludeComment != "" {
				query.Add("comment", searchParams.IncludeComment)
			}
			response, error := client.Post.GetPosts(team, query)

			if searchParams.Debug {
				fmt.Println(query.Encode())
			}

			if error != nil {
				return error
			}

			for _, value := range response.Posts {
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
