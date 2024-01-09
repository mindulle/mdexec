/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var replWhtspCmd = &cobra.Command{
	Use:   "replWhtsp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exeDest := "bulk/replace-whitespace-with-delimiter.sh"

		if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	bulkCmd.AddCommand(replWhtspCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replWhtspCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replWhtspCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
