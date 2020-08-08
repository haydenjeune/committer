package profile

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
)

func Test_NewStorage_MakesNewStorage(t *testing.T) {
	fs := memfs.New()
	storage := NewStorage(fs)
	if storage.fs != fs {
		t.Error("Filesystem not set inside storage")
	}
}

func Test_Read_ReadsProfiles(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/homedir/", 0600)
	file, _ := fs.Create("/test/homedir/profiles")
	file.Write([]byte("[testProfile]\n"))
	file.Write([]byte("  Name = \"test\"\n"))
	file.Write([]byte("  Email = \"a@b.com\"\n"))
	file.Close()

	homeFs, _ := fs.Chroot("/test/homedir")

	profileStore := NewStorage(homeFs)
	profiles, err := profileStore.Read()
	if err != nil {
		t.Error(err)
		return
	}
	profile, ok := profiles["testProfile"]
	if !ok {
		t.Error("No profile for 'testProfile' found.")
	} else if profile.Name != "test" || profile.Email != "a@b.com" {
		t.Error("Incorrect profile for 'testProfile' found.")
	}
}

func Test_Read_ReturnsError_ConfigInvalid(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/homedir/", 0600)
	file, _ := fs.Create("/test/homedir/profiles")
	file.Write([]byte("testProfile\n"))
	file.Close()

	homeFs, _ := fs.Chroot("/test/homedir")

	profileStore := NewStorage(homeFs)
	_, err := profileStore.Read()
	if err == nil {
		t.Error("Read() should have failed but didn't")
		return
	}
}

func Test_Read_ReturnsError_ConfigNotExist(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/homedir/", 0600)

	homeFs, _ := fs.Chroot("/test/homedir")

	profileStore := NewStorage(homeFs)
	_, err := profileStore.Read()
	if err == nil {
		t.Error("Read() should have failed but didn't.")
	}
}

func Test_Read_ReturnsError_TooManyKeys(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/homedir/", 0600)
	file, _ := fs.Create("/test/homedir/profiles")
	file.Write([]byte("[testProfile]\n"))
	file.Write([]byte("  Name = \"test\"\n"))
	file.Write([]byte("  Email = \"a@b.com\"\n"))
	file.Write([]byte("  ExtraKey = \"Random data\"\n"))
	file.Close()

	homeFs, _ := fs.Chroot("/test/homedir")

	profileStore := NewStorage(homeFs)
	_, err := profileStore.Read()
	if err == nil {
		t.Error("Read() should have failed but didn't.")
	}
}

func Test_Save_SavesProfile(t *testing.T) {
	p := Profile{Name: "tester", Email: "a@b.com"}
	profiles := make(map[string]Profile)
	profiles["test"] = p
	fs := memfs.New()
	profileStore := NewStorage(fs)

	err := profileStore.Save(profiles)
	if err != nil {
		t.Error(err)
		return
	}

	file, err := fs.Open("profiles")
	if err != nil {
		t.Error(err)
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}
	str := string(bytes)
	if !strings.Contains(str, "[test]") ||
		!strings.Contains(str, "Name = \"tester\"") ||
		!strings.Contains(str, "Email = \"a@b.com\"") {
		t.Error("Incorrectly written profile file")
	}
}

// memfs does not provide sufficient control to trigger the error conditions in
// Storage.Save without sacrificing readability of the library, or implementing a
// whole billy.FileSystem for testing. These error conditions are unlikely to be
// the fault of any changes made in this library so no tests are included for
// those errors.
