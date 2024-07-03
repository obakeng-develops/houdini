/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type buildOptions struct {
	path string
	tag  string
}

var buildCmdOptions = &buildOptions{}

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build will read the given file path and build an image",
	Long: `
	# This command builds the image based on the given file path
	houdini build -i --path=path/to/app -t my-app
	
	# This command builds the file path/repo given a tag by the user as well as provision containers
	houdini build --path=path/to/app -t my-app`,
	Run: func(cmd *cobra.Command, args []string) {
		buildCmdOptions.run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&buildCmdOptions.path, "path", "p", "", "Provide a directory path to build an image from")
}

func (b *buildOptions) run() {
	fmt.Println("Say hello")
}
