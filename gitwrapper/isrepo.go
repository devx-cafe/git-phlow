package gitwrapper

import (
	"os"
	"path/filepath"
	"github.com/libgit2/git2go"
)

//GetCurrentDirectory gets the absolute path of the current directory
func GetCurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return "", err
	}
	return dir, nil
}

//IsRepository checks if current dir is a git repository
func IsRepository(path string) bool {
	_, error := git.OpenRepository(path)

	return error == nil
}
