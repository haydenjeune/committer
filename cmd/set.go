package cmd

import (
	"fmt"
	"os"

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
	// get config
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dotGitPath, err := utils.FindDotGit(pwd)
	if err == utils.ErrIsNotRepository {
		fmt.Println(err)
		return
	}
	if err != nil {
		panic(err)
	}

	conf, err := utils.GetLocalConfig(dotGitPath)
	if err != nil {
		panic(err)
	}

	// modify config
	conf.User.Name = "haydenjeune"
	conf.User.Email = "33794706+haydenjeune@users.noreply.github.com"

	// set config
	err = utils.SetLocalConfig(dotGitPath, conf)
	if err != nil {
		panic(err)
	}
}
