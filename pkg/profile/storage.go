package profile

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5"
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

const storageFile string = "profiles"

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
	file, err := s.fs.Open(storageFile)
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

// Save the given profiles to the filesystem
func (s *Storage) Save(profiles map[string]Profile) error {
	file, err := s.fs.Create(storageFile)
	if err != nil {
		// TODO: But what does %w do
		return fmt.Errorf("failed to open or create profiles file: %v", err)
	}
	defer file.Close()

	enc := toml.NewEncoder(file)
	if err := enc.Encode(profiles); err != nil {
		return fmt.Errorf("failed to write profiles to file: %v", err)
	}
	return nil
}
