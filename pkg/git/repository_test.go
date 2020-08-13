package git

import (
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/config"
)

func Test_ReadRepository_ReadsRepository(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/folder/.git", 0600)

	repo, err := ReadRepository("/test/folder", fs)
	if err != nil {
		t.Error(err)
		return
	}

	expected := "/test/folder"
	if repo.Root() != expected {
		t.Errorf("Wrong repository found. Expected '%s', found '%s'", expected, repo.Root())
		return
	}
}

func Test_FindRepository_CantFindRepo_IsNotRepo(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/folder", 0600)

	result, err := FindRepository("/test/folder", fs)

	if err == nil {
		t.Errorf("FindDotGit shouldn't have found a git repo but found repo with root '%s'", result.Root())
		return
	}
	if err != ErrIsNotRepository {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func Test_FindRepository_FindsRepo_IsRepo(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/folder/.git", 0600)

	repo, err := FindRepository("/test/folder", fs)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}
	expected := "/test/folder"
	if repo.Root() != expected {
		t.Errorf("Wrong repository found. Expected '%s', found '%s'", expected, repo.Root())
	}
}

func Test_FindRepository_FindsRepo_IsChildDirOfRepo(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/folder/.git", 0600)
	fs.MkdirAll("/test/folder/child", 0600)

	repo, err := FindRepository("/test/folder/child", fs)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}
	expected := "/test/folder"
	if repo.Root() != expected {
		t.Errorf("Wrong repository found. Expected '%s', found '%s'", expected, repo.Root())
	}
}

func Test_Config_GetsCorrectConfig_ReadingConfig(t *testing.T) {
	fs := memfs.New()
	file, _ := fs.Create("/test/folder/.git/config")
	file.Write([]byte("[user]\n"))
	file.Write([]byte("\tname = test\n"))
	file.Write([]byte("\temail = a@b.com\n"))
	file.Close()
	repo, _ := ReadRepository("/test/folder/", fs)

	conf, err := repo.Config()
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if conf.User.Name != "test" || conf.User.Email != "a@b.com" {
		t.Errorf("Incorrect config read")
	}
}

func Test_SetConfig_WritesCorrectFile_FromConfig(t *testing.T) {
	conf := config.NewConfig()
	conf.User.Name = "test"
	conf.User.Email = "a@b.com"

	fs := memfs.New()
	file, _ := fs.Create("/test/folder/.git/config")
	file.Close()
	repo, _ := ReadRepository("/test/folder/", fs)

	err := repo.SetConfig(conf)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	file, err = fs.Open("/test/folder/.git/config")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	defer file.Close()
	buf := make([]byte, 100)
	file.Read(buf)
	bufString := string(buf)
	if !strings.Contains(bufString, "[user]") ||
		!strings.Contains(bufString, "name = test") ||
		!strings.Contains(bufString, "email = a@b.com") {
		t.Error("Incorrectly written config file")
	}
}

func Test_Root_ReturnsCorrectDir(t *testing.T) {
	fs := memfs.New()
	fs.MkdirAll("/test/folder/.git", 0600)
	repo, _ := ReadRepository("/test/folder/", fs)

	root := repo.Root()
	expected := "/test/folder"
	if root != expected {
		t.Errorf("Incorrect repository root. Expected '%s', got '%s'", expected, root)
	}
}
