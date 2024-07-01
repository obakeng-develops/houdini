/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// playgroundCmd represents the playground command
var playgroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Playground is a test command to try out cobra",
	Long: `
	houdini playground --path=.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("playground called")
	},
}

func init() {
	rootCmd.AddCommand(playgroundCmd)

}
