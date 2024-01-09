/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var deepGetCmd = &cobra.Command{
	Use:   "deepGet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exeDest := "json/deep-get-object-value.js"

		if exampleFlag, _ := cmd.Flags().GetBool("example"); exampleFlag {
			exeDest = "json/deep-get-object-value.ex.js"
		}

		if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	jsonCmd.AddCommand(deepGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deepGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deepGetCmd.Flags().BoolP("example", "x", false, "See an example output. run with --raw option to see source code.")
}
