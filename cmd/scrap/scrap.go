/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/hoehwa/gopkg/scrap"
	"github.com/mindulle/mdexec/cmd"
	"github.com/spf13/cobra"
)

var scrapCmd = &cobra.Command{
	Use:   "scrap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		url := args[0]
		selector := args[1]

		fmt.Println(scrap.BySelector(url, selector))
	},
}

func init() {
	cmd.RootCmd.AddCommand(scrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
