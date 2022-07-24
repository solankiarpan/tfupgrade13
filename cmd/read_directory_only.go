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

var DirectoryOnly string

var directoryOnlyCmd = &cobra.Command{
	Use:                   "o <directory_path>",
	Short:                 "Upgrade tf syntax for the dir non recursively",
	Long:                  `Upgrade syntax to terraform 0.13.0`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if DirectoryOnly == "" && len(args) < 1 {
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Example: `tfupgrade13 o /Downloads`,
	Run: func(cmd *cobra.Command, args []string) {
		var dirname string
		var err error
		var argument string
		if DirectoryOnly != "" {
			argument = DirectoryOnly
		} else {
			argument = args[0]
		}
		directoryExists, err := helper.DirectoryExists(argument)
		if err != nil {
			fmt.Println("Directory does not exist or invalid argument: ", argument, " Recheck the directory path")
		}
		if directoryExists {
			dirname, err = filepath.Abs(argument)
			if err != nil {
				log.Fatal("Invalid directory name: ", dirname, " Please enter a valid directory name")
			}
		} else {
			fmt.Println("File does not exist")
		}

		utils.ReplaceLineInDirOnly(dirname)
		fmt.Println("Successfully replaced specific strings in the files..!!")
		fmt.Println("Process Completed")
	},
}

func init() {
	rootCmd.AddCommand(directoryOnlyCmd)
	fileCmd.Flags().StringVarP(&File, "dironly", "o", "", "directory path")
}
