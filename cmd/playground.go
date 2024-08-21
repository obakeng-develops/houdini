/*
Copyright © 2024 Obakeng Mosadi <mosadiobakeng7@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

type playgroundOptions struct {
	path string
}

// playgroundCmd represents the playground command
var (

	playgroundCmdOptions = &playgroundOptions{}

	playgroundCmd = &cobra.Command{
		Use:   "playground",
		Short: "Playground is a test command to try out cobra",
		Long: `
		houdini playground --path=.`,
		Run: func(cmd *cobra.Command, args []string) {
			playgroundCmdOptions.run()
	},
}
)

func init() {
	rootCmd.AddCommand(playgroundCmd)

	playgroundCmd.Flags().StringVarP(&playgroundCmdOptions.path, "path", "p", ".", "Path to the folder to read")
}

func (options *playgroundOptions) run() {
	dir, err := os.Lstat(options.path)
	if err != nil {
		fmt.Println(err)
		return
	}

	mode := dir.Mode()

	if mode.IsDir() {
		fmt.Println("It is a directory")
	} else if mode.IsRegular() {
		fmt.Println("It is a regular file.")
	} else {
		fmt.Println("It is something else")
	}
}
