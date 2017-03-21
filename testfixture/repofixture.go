package testfixture

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/praqma/git-phlow/executor"
)

var (
	//GoPath ...
	GoPath string
	//ProjectPath ..
	ProjectPath string
	//Script ...
	Script string
	//Repo ...
	Repo string
	//Target ...
	Target string
)

//init
//Set GoPath variable before execution
func init() {
	GoPath = os.Getenv("GOPATH")
	if len(GoPath) == 0 {
		fmt.Fprintln(os.Stdout, errors.New("GOPATH not set"))
		os.Exit(1)
	}
	ProjectPath = GoPath + "/src/github.com/praqma/git-phlow"
	Script = ProjectPath + "/testfixture/gen_test_repo.sh"
	Repo = ProjectPath + "/build/phlow-test-pkg"
	Target = ProjectPath + "/build"
}

//CreateTestRepository ...
//Runs gen_test_repo shell script
func CreateTestRepository(test *testing.T, verbose bool) {
	output, err := executor.ExecuteCommand(Script)
	if err != nil {
		test.Log(err)
		os.Exit(1)
	}
	if verbose {
		test.Log(output)
	}
	os.Chdir(Repo)
}

//RemoveTestRepository ...
//Deletes the test repository and folders
func RemoveTestRepository(test *testing.T) {
	os.Chdir(ProjectPath)
	err := os.RemoveAll(Target)
	if err != nil {
		test.Log(err.Error())
		os.Exit(1)
	}
}
