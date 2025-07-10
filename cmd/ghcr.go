/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ghcrCmd represents the ghcr command
var ghcrCmd = &cobra.Command{
	Use:   "ghcr",
	Short: "Commands for interacting with GitHub Conatiner Registry",
	Long:  `A group of commands which help you manage and inspect images in GHCR`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ghcr called")
	},
}

func init() {

	rootCmd.AddCommand(ghcrCmd)

}
