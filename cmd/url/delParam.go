/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package url

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var delParamCmd = &cobra.Command{
	Use:   "delParam",
	Short: "Delete a url parameter.",
	Long: `
	provide a url with parameter as a first command parameter like this:
	https://web.dev/articles/webauthn-discoverable-credentials?hl=en

	and pass a key of url parameter as a second command parameter.

	e.g. mdexec url delParam https://web.dev/articles/webauthn-discoverable-credentials?hl=en hl`,
	Run: func(cmd *cobra.Command, args []string) {
		exeDest := "url/del-a-url-param.js"

		if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	urlCmd.AddCommand(delParamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delParamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delParamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
