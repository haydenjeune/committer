package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/utils"
	"github.com/spf13/cobra"
)

// setCmd handles the `committer set` command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set committer info on current repo.",
	Run: func(cmd *cobra.Command, args []string) {
		runSet(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func runSet(profileName string) {
	// Setup filesystem
	fs := osfs.New("/")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Find repo in present dir or a parent
	gitRepo, err := utils.FindGitRepository(pwd, fs)
	if err == utils.ErrIsNotRepository {
		fmt.Println(err)
		return
	}
	if err != nil {
		panic(err)
	}

	// Read config
	conf, err := gitRepo.Config()
	if err != nil {
		panic(err)
	}

	// Modify config
	conf.User.Name = "haydenjeune"
	conf.User.Email = "33794706+haydenjeune@users.noreply.github.com"

	// Set config
	err = gitRepo.SetConfig(conf)
	if err != nil {
		panic(err)
	}
}
