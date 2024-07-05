/*
Copyright © 2024 Obakeng Mosadi <mosadiobakeng7@gmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/obakeng-develops/houdini/pkg"
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
		buildCmdOptions.validate()
		buildCmdOptions.run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&buildCmdOptions.path, "path", "p", "", "Provide a directory path to build an image from")
	buildCmd.Flags().StringVarP(&buildCmdOptions.tag, "tag", "t", "", "Provide a tag for your image")
}

// validate the command options
func (b *buildOptions) validate() error {
	if b.path == "" {
		slog.Error("You need to provide a directory path")
	}

	if b.tag == "" {
		slog.Error("You need to provide an image tag")
	}

	return nil
}

func (b *buildOptions) run() error {

	findDominantLanguage, err := pkg.DirectoryWalkthrough(b.path)
	if err != nil {
		slog.Error("could not find the directory", "err", err)
	}

	// Check output
	fmt.Println(findDominantLanguage)

	return nil
}
