/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package regexp

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var exeDest string
		scopeFlag, _ := cmd.Flags().GetString("scope")

		switch scopeFlag {
		case "capturing-group":
			exeDest = "text/regexp/capturing-groups.ex.js"
		case "capturing-group-backreferences":
			exeDest = "text/regexp/capturing-group-backreferences.ex.js"
		case "lookaheds":
			exeDest = "text/regexp/lookaheds.ex.js"
		case "named-capturing-groups":
			exeDest = "text/regexp/named-capturing-groups.ex.js"
		case "non-capturing-groups":
			exeDest = "text/regexp/non-capturing-groups.ex.js"
		case "unicode-characters":
			exeDest = "text/regexp/unicode-characters.ex.js"
		default:
			exeDest = "text/regexp/all.ex.js"
		}

		if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	regexpCmd.AddCommand(exampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	exampleCmd.Flags().StringP("scope", "s", "all", `Set a scope to see usage. followings are	available scopes:
	capturing-groups, capturing-group-backreferences, lookaheds, 
	named-capturing-groups, non-capturing-groups, unicode-characters`)
}
