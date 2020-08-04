package cmd

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"github.com/go-git/go-billy/v5/osfs"
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
	err := inputFromDefault(&defaultProfile)
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

// inputFromDefault takes a pointer to a profile and iterates over its fields,
// setting each one based on input from stdin, with prompts sent to stdout.
// Fields written to stdin in Reader are separated by newlines. If no bytes
// are written before a newline then the default value for that field will
// be used. Values are updated in-place on the provided struct.
func inputFromDefault(defaultProfile *profile.Profile) error {
	v := reflect.ValueOf(defaultProfile).Elem()
	t := v.Type()

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s [%v]:", t.Field(i).Name, v.Field(i).String())
		scanner.Scan()
		text := scanner.Text()
		if text != "" {
			v.Field(i).SetString(text)
		}
	}
	return nil
}
