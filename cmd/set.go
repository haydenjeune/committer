package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd handles the `committer set` command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set committer info on current repo.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
