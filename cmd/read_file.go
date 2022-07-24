package cmd

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"tfupgrade13/helper"
	"tfupgrade13/utils"

	"github.com/spf13/cobra"
)

var File string

var fileCmd = &cobra.Command{
	Use:                   "f <file_path>",
	Short:                 "Upgrade tf syntax for the provided file",
	Long:                  `Upgrade syntax to terraform 0.13.0`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Example: `tfupgrade13 f /Downloads/file.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		var err error
		var argument string
		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}
		fileExists, err := helper.FileExists(argument)
		if err != nil {
			fmt.Println("File does not exist or invalid argument: ", argument, " Recheck the file path")
		}
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				log.Fatal("Invalid file name: ", filename, " Please enter a valid file name")
			}
		} else {
			fmt.Println("File does not exist")
		}

		utils.ReplaceLine(filename)
		fmt.Println("Successfully updated syntax to terraform 0.13.0")
		fmt.Println("Process Completed")

	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().StringVarP(&File, "file", "f", "", "file path")
}
