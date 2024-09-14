package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "A simple Random password Generator",
	Long:  `A Bare simple random password generator`,
}

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

// Execute executes the root command.
func Execute() {
	completion := completionCommand()

	// mark completion hidden
	completion.Hidden = true
	rootCmd.AddCommand(completion)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
