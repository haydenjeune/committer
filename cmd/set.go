package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/haydenjeune/committer/pkg/profile"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/haydenjeune/committer/pkg/git"
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
		panic(err)
	}

	// Find repo in present dir or a parent
	gitRepo, err := git.FindRepository(pwd, fs)
	if err == git.ErrIsNotRepository {
		fmt.Println(err)
		return
	}
	if err != nil {
		panic(err)
	}

	conf, err := gitRepo.Config()
	if err != nil {
		panic(err)
	}

	// Read Profiles
	profileStore := profile.NewStorage(osfs.New(committerConfigDir()))
	profiles, err := profileStore.Read()

	// Modify config
	if user, ok := profiles[profileName]; ok {
		conf.User = user
	} else {
		fmt.Printf("Profile '%s' does not exist\n", profileName)
		return
	}

	// Set config
	if err = gitRepo.SetConfig(conf); err != nil {
		panic(err)
	}
}

func committerConfigDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".committer")
}
