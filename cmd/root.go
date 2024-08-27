/*
Copyright © 2024 Obakeng Mosadi mosadiobakeng7@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "houdini",
	Short: "Tool that reads your folder and auto-provisions containers",
	Long:  `Houdini is a CLI tool that auto-provisions containers by reading a code folder.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
