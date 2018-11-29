package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/upamune/go-esa/esa"
)

var RootCmd = &cobra.Command{
	Use:           "myesa",
	Short:         "esa client",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	viper.SetConfigName("myesarc")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")

	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read config file: %v", err)
		os.Exit(1)
	}
	accessToken := viper.GetString("ACCESS_TOKEN")
	team := viper.GetString("TEAM")
	client := esa.NewClient(accessToken)
	cobra.OnInitialize()
	RootCmd.AddCommand(
		searchCmd(*client, team),
	)
}
