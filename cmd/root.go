package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "tfupgrade13",
	Version:            "1.0.0",
	Short:              "A Text Replacement Tool",
	Long:               `Text Replacer Tool version 1.0.0`,
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
