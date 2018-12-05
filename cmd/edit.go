package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/upamune/go-esa/esa"
)

func editCmd(client esa.Client, team string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit articles",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			editor := os.Getenv("EDITOR")

			if editor == "" {
				fmt.Println("$EDITOR not found.")
				return nil
			}
			execEditor(editor)
			// response, error := client.Post.GetPosts(team, query)
			return nil
		},
	}

	return cmd
}

func execEditor(editor string) {
	fmt.Println(editor)
	cmd := exec.Command("/bin/bash", "-c", editor, "./__tmp__")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		fmt.Println("ERROR")
		return
	}
	fmt.Println("done")
}
