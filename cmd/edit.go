package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
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

			body := execEditor(editor)
			post := esa.Post{
				Name:   args[0],
				BodyMd: body,
			}
			response, err := client.Post.Create(team, post)

			if err != nil {
				log.Fatal(err)
				return nil
			}

			fmt.Println(response.URL)
			// response, error := client.Post.GetPosts(team, query)
			return nil
		},
	}

	return cmd
}

func execEditor(editor string) string {
	cmdstr := fmt.Sprintf("tpich __tmp__ & %s __tmp__", editor)
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("./__tmp__")

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
