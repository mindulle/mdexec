/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		if iFlag, _ := cmd.Flags().GetBool("toInput"); iFlag {
			internal.RunForigenScriptFrom("bulk/rename-files-to-input-recursively.sh", args...)
			if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
				beautify.RawContent(internal.BaseDir, "bulk/rename-files-to-input-recursively.sh")
			}
			return
		}
	},
}

func init() {
	bulkCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	renameCmd.Flags().BoolP("toInput", "i", false, "Change names of files into a user input recursively")
}
