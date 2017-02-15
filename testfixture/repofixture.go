package testfixture

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var (
	errGoPathNotSet = errors.New("GOPATH is empty")
	goPath          string
	phlowPath       string
)

//init
//Runs before functions to setup variable gopath
func init() {
	goPath = os.Getenv("GOPATH")
	if len(goPath) == 0 {
		fmt.Fprintln(os.Stdout, errGoPathNotSet)
		os.Exit(1)
	}

	phlowPath = goPath + "/src/github.com/praqma/git-phlow"
}

//SetupTestRepo ...
//Creates git test repository from a zip file in /testfixture
func SetupTestRepo() {

	script := phlowPath + "/testfixture/gen_test_repo.sh"
	repo := phlowPath + "/build/phlow-test-pkg"

	cmd := exec.Command(script)
	cmd.Start()

	err := cmd.Wait()

	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
	os.Chdir(repo)
}

//TearDownTestRepo ...
//removes the unzipped test repository is it exists
func TearDownTestRepo() {

	var target = phlowPath + "/build"
	os.Chdir(phlowPath)
	err := os.RemoveAll(target)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

}
