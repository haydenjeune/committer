package cmd

import (
	"fmt"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/errors"
	"github.com/haydenjeune/committer/pkg/profile"
	"github.com/spf13/cobra"
)

// rmCmd handles the `committer rm` command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a saved committer profile",
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		profileStore := profile.NewStorage(osfs.New(committerConfigDir()))
		profiles, err := profileStore.Read()
		if err != nil {
			errors.PrintAndExit(fmt.Errorf("error reading saved profiles: %v", err))
		}

		if _, exists := profiles[profileName]; !exists {
			return
		}

		delete(profiles, profileName)
		if err := profileStore.Save(profiles); err != nil {
			fmt.Printf("error deleting profile: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
