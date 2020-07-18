package utils

import (
	"errors"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/config"
)

// ErrIsNotRepository signals that no git repositories have been found
var ErrIsNotRepository = errors.New("This directory and all parents are not git repositories")

// ErrFailedReadGitConfig signals failure to read .git/config
var ErrFailedReadGitConfig = errors.New("Failed to read git config")

// GitRepository implements the git-go storage.ConfigStorer interface
type GitRepository struct {
	fs billy.Filesystem
}

// ReadRepository creates a GitRepository given a fs, and the relative path to the repo in that fs.
func ReadRepository(path string, fs billy.Filesystem) (*GitRepository, error) {
	repoFs, err := fs.Chroot(path)
	if err != nil {
		panic(err)
	}
	return &GitRepository{repoFs}, nil
}

// FindGitRepository finds the .git directory in the given directory, or any parents.
func FindGitRepository(fromPath string, fs billy.Filesystem) (*GitRepository, error) {
	// if root return error not found
	if fromPath == "/" {
		return nil, ErrIsNotRepository
	}

	// if this isn't a repo, recurse up
	_, err := fs.Stat(fs.Join(fromPath, ".git"))
	if err != nil {
		return FindGitRepository(fs.Join(fromPath, ".."), fs)
	}

	// otherwise it is, so return
	gitRepo, err := ReadRepository(fromPath, fs)
	if err != nil {
		panic(err)
	}
	return gitRepo, nil
}

// Config reads a Config object from a GitRepository
func (repo *GitRepository) Config() (*config.Config, error) {
	file, err := repo.fs.Open(".git/config")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	conf, err := config.ReadConfig(file)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// SetConfig saves a config object to the GitRepository filesystem
func (repo *GitRepository) SetConfig(conf *config.Config) error {
	bytes, err := conf.Marshal()
	if err != nil {
		return err
	}

	file, err := repo.fs.Create(".git/config")
	defer file.Close()
	if err != nil {
		return err
	}

	file.Write(bytes)
	return nil
}

// Root prints the root directory of the GitRepository
func (repo *GitRepository) Root() string {
	return repo.fs.Root()
}
