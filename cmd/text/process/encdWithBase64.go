/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package process

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var encdWithBase64Cmd = &cobra.Command{
	Use:   "encdWithBase64",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exeDest := "text/process/encode-with-base64.js"
		polyfillExeDest := "text/process/encode-with-base64-polyfill.js"

		rFlag, _ := cmd.Flags().GetBool("raw")
		polyfillFlag, _ := cmd.Flags().GetBool("polyfill")

		if rFlag {
			if polyfillFlag {
				beautify.RawContent(internal.BaseDir, polyfillExeDest)
				return
			}

			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		if polyfillFlag {
			internal.RunForigenScriptFrom(polyfillExeDest, args...)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	processCmd.AddCommand(encdWithBase64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encdWithBase64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	encdWithBase64Cmd.Flags().BoolP("polyfill", "p", false, "change desitination to execute to polyfill script.")
}
