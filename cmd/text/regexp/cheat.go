/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package regexp

import (
	"github.com/hoehwa/gopkg/beautify"
	"github.com/mindulle/mdexec/internal"
	"github.com/spf13/cobra"
)

var cheatCmd = &cobra.Command{
	Use:   "cheat",
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
		case "anchor":
			exeDest = "text/regexp/anchors.cheat.js"
		case "charSeq":
			exeDest = "text/regexp/character-sequences.cheat.js"
		case "flag":
			exeDest = "text/regexp/flag.cheat.js"
		case "group":
			exeDest = "text/regexp/group.cheat.js"
		case "quantifier":
			exeDest = "text/regexp/quantifiers.cheat.js"
		default:
			exeDest = "text/regexp/all.cheat.js"
		}

		if rFlag, _ := cmd.Flags().GetBool("raw"); rFlag {
			beautify.RawContent(internal.BaseDir, exeDest)
			return
		}

		internal.RunForigenScriptFrom(exeDest, args...)
	},
}

func init() {
	regexpCmd.AddCommand(cheatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cheatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cheatCmd.Flags().BoolP("anchor", "a", false, "Help message for toggle")
	cheatCmd.Flags().StringP("scope", "s", "all", `Set a scope to see usage. followings are	available scopes: 
	anchor, charSeq, flag, group, quantifiers`)
}
