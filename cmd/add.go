package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/input"
	"github.com/haydenjeune/committer/pkg/profile"
	"github.com/spf13/cobra"
)

// addCmd handles the `committer add` command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add committer profile to profile store",
	Run: func(cmd *cobra.Command, args []string) {
		runAdd(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(profileName string) {
	// Get existing profiles
	profileStore := profile.NewStorage(osfs.New(committerConfigDir()))
	profiles, _ := profileStore.Read()

	// Get new profile
	defaultProfile := profiles[profileName]
	err := input.StructFromDefault(&defaultProfile, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Printf("error taking user input: %v\n", err)
		return
	}
	profiles[profileName] = defaultProfile

	err = profileStore.Save(profiles)
	if err != nil {
		fmt.Printf("error saving profile: %v\n", err)
	}
}
