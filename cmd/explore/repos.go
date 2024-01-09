/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	github "github.com/hoehwa/gopkg/github/search"
	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		argLen := len(args)
		if argLen > 3 || argLen < 1 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		var query github.Query

		switch argLen {
		case 1:
			query = github.NewQuery(args[0])
		case 2:
			query = github.NewQuery(args[0]).WithLanguage(args[1])
		case 3:
			query = github.NewQuery(args[0]).WithLanguage(args[1]).WithTopic(args[2])
		}

		github.SearchTopRepos(query)
	},
}

func init() {
	exploreCmd.AddCommand(reposCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reposCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reposCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
