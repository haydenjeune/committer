package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_FindDotGit_CantFindDotGit_IsNotRepo(t *testing.T) {
	dir, err := ioutil.TempDir("", "committer-test-*")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Error(err)
	}
	result, err := FindDotGit(dir)
	if result != "" || err != ErrIsNotRepository {
		t.Errorf("FindDotGit shouldn't have found a git repo but found %s", result)
	}
}

func Test_FindDotGit_FindsDotGit_IsRepo(t *testing.T) {
	dir, err := ioutil.TempDir("", "committer-test-*")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Error(err)
	}
	os.Mkdir(dir+"/.git", 0600)
	result, err := FindDotGit(dir)
	if result != dir+"/.git" || err != nil {
		t.Error("FindDotGit should have found a git repo but didn't")
	}
}

func Test_FindDotGit_FindsDotGit_IsChildDirOfRepo(t *testing.T) {
	dir, err := ioutil.TempDir("", "committer-test-*")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Error(err)
	}

	dotGitDir := dir + "/.git"
	os.Mkdir(dotGitDir, 0600)

	testDir := dir + "/child_dir"
	os.Mkdir(testDir, 0600)

	result, err := FindDotGit(testDir)
	if err == ErrIsNotRepository {
		t.Errorf("FindDotGit should have found a git repo at %s but didn't", dotGitDir)
	}
	if result != dotGitDir {
		t.Errorf("FindDotGit should have found a git repo at %s but found one at %s (base %s)", dotGitDir, result, dir)
	}
}
