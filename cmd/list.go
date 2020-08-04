package cmd

import (
	"fmt"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/profile"
	"github.com/spf13/cobra"
)

// addCmd handles the `committer add` command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved committers",
	Run: func(cmd *cobra.Command, args []string) {
		profileStore := profile.NewStorage(osfs.New(committerConfigDir()))
		profiles, _ := profileStore.Read()

		for name, profile := range profiles {
			fmt.Printf("%s: {\n\tuser.name: \t%s\n\tuser.email: \t%s\n}\n", name, profile.Name, profile.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
