/*
Package cmd contains all of the logic for parsing input.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "committer",
	Short: "Set and maintain repo specific committer info",
	Long: `Committer is a tool that makes it easy to set the details that you 
	use to commit to a git repo. You might find it useful if you have different
	github accounts for uni, work, and personal things.`,
}

// Execute parses input and runs the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
