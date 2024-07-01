/*
Copyright © 2024 Obakeng Mosadi <mosadiobakeng7@gmail.com>

*/
package cmd

import (
	"fmt"

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
	fmt.Println(options.path)
}
