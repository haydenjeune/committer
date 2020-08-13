package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/errors"
	"github.com/haydenjeune/committer/pkg/git"
	"github.com/haydenjeune/committer/pkg/profile"
	"github.com/mitchellh/go-homedir"
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
		errors.PrintAndExit(fmt.Errorf("error determining current directory: %v", err))
	}

	// Find repo in present dir or a parent
	gitRepo, err := git.FindRepository(pwd, fs)
	if err == git.ErrIsNotRepository {
		errors.PrintAndExit(err)
	}
	if err != nil {
		errors.PrintAndExit(fmt.Errorf("error finding git repository: %v", err))
	}

	conf, err := gitRepo.Config()
	if err != nil {
		errors.PrintAndExit(fmt.Errorf("error retrieving repository config: %v", err))
	}

	// Read Profiles
	profileStore := profile.NewStorage(osfs.New(committerConfigDir()))
	profiles, err := profileStore.Read()

	// Modify config
	if user, exists := profiles[profileName]; !exists {
		errors.PrintAndExit(fmt.Errorf("profile '%v' does not exist", profileName))
	} else {
		conf.User = user
	}

	// Set config
	if err = gitRepo.SetConfig(conf); err != nil {
		errors.PrintAndExit(fmt.Errorf("failed to save repo config: %v", err))
	}
}

func committerConfigDir() string {
	home, err := homedir.Dir()
	if err != nil {
		errors.PrintAndExit(fmt.Errorf("failed to find committer config: %v", err))
	}
	return path.Join(home, ".committer")
}
