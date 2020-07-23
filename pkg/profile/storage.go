package profile

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5"
	"github.com/pkg/errors"
)

// Profile stores the details of the committer.
type Profile struct {
	// Name is the personal name committer.
	Name string
	// Email is the email of the commiter.
	Email string
}

// ErrUnexpectedProfileKey is returned when an unexpected key is found in the profiles file
var ErrUnexpectedProfileKey = errors.New("unexpected key found in profile config")

// Storage writes and reads Profiles to the given filesystem
type Storage struct {
	fs billy.Filesystem
}

// NewStorage returns an instance of Storage configured to use the given filesystem
func NewStorage(fs billy.Filesystem) *Storage {
	return &Storage{fs}
}

// Read the saved profiles from the filesystem
func (s *Storage) Read() (map[string]Profile, error) {
	profiles := make(map[string]Profile)
	file, err := s.fs.Open("profiles")
	if err != nil {
		return nil, fmt.Errorf("failed to open profiles file: %v", err)
	}
	defer file.Close()
	if meta, err := toml.DecodeReader(file, &profiles); err != nil {
		return nil, fmt.Errorf("failed to parse profiles: %v", err)
	} else if len(meta.Undecoded()) > 0 {
		return nil, ErrUnexpectedProfileKey
	}
	return profiles, nil
}
