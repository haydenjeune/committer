package utils

import (
	"errors"
	"os"
	"path"

	"github.com/go-git/go-git/v5/config"
)

// ErrIsNotRepository signals that no git repositories have been found
var ErrIsNotRepository = errors.New("This directory and all parents are not git repositories")

// FindDotGit finds the .git directory in the given directory, or any parents.
func FindDotGit(basePath string) (string, error) {
	if basePath == "." || basePath == "/" {
		return "", ErrIsNotRepository
	}
	dotGitPath := path.Join(basePath, ".git")
	_, err := os.Stat(dotGitPath)
	if os.IsNotExist(err) {
		return FindDotGit(path.Dir(basePath)) // Dir removes the last part of the path
	}
	return dotGitPath, nil
}

// GetLocalConfig reads a Config object given a .git path
func GetLocalConfig(dotGitPath string) (*config.Config, error) {
	file, err := os.Open(dotGitPath + "/config")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	conf, err := config.ReadConfig(file)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// SetLocalConfig write a config file given a .git path and a Config object
func SetLocalConfig(dotGitPath string, conf *config.Config) error {
	bytes, err := conf.Marshal()
	if err != nil {
		return err
	}

	file, err := os.Create(dotGitPath + "/config")
	defer file.Close()
	if err != nil {
		return err
	}
	file.Write(bytes)
	return nil
}
