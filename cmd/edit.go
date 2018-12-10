package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/nasum/myesa/lib"
	"github.com/spf13/cobra"
)

func editCmd() *cobra.Command {
	esaClient := &lib.EsaClient{}

	esaClient.Init()

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

			articles := esaClient.GetArticlesByName(args[0])

			if len(articles) == 1 {
				body := execEditor(editor, articles[0].BodyMd)

				url := esaClient.Update(articles[0].Number, args[0], body)

				fmt.Println(url)
			} else {
				body := execEditor(editor, "")

				url := esaClient.Create(args[0], body)

				fmt.Println(url)
			}
			return nil
		},
	}

	return cmd
}

func execEditor(editor string, body string) string {
	file, err := os.OpenFile("__tmp__", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	if body != "" {
		fmt.Fprintln(file, body)
	}

	defer os.Remove(file.Name())

	cmdstr := fmt.Sprintf("%s __tmp__", editor)
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	data, err := ioutil.ReadFile("./__tmp__")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(data)
}
