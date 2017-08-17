package testfixture

import (
	"os"
	"testing"

	"github.com/praqma/git-phlow/executor"
	"log"
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
		log.Panicln("GOPATH NOT SET")
	}
	ProjectPath = GoPath + "/src/github.com/praqma/git-phlow"
	Script = ProjectPath + "/testfixture/gen_test_repo.sh"
	Repo = ProjectPath + "/build/phlow-test-pkg"
	Target = ProjectPath + "/build"
}

//CreateTestRepository ...
//Runs gen_test_repo shell script
func CreateTestRepositoryNoLog(verbose bool) {
	output, err := executor.RunCommand(Script)
	if err != nil {
		log.Panicln(output, err)
	}

	if err = os.Chdir(Repo); err != nil {
		log.Panicln(err)
	}
}

//RemoveTestRepository ...
//Deletes the test repository and folders
func RemoveTestRepositoryNoLog() {
	if err := os.Chdir(ProjectPath); err != nil {
		log.Panicln(err)
	}

	err := os.RemoveAll(Target)
	if err != nil {
		log.Panicln(err)
	}
}

//CreateTestRepository ...
//Runs gen_test_repo shell script
func CreateTestRepository(test *testing.T, verbose bool) {
	output, err := executor.RunCommand(Script)
	if err != nil {
		test.Log(err)
		os.Exit(1)
	}
	if verbose {
		test.Log(output)
	}
	if err = os.Chdir(Repo); err != nil {
		log.Panicln(err)
	}
}

//RemoveTestRepository ...
//Deletes the test repository and folders
func RemoveTestRepository(test *testing.T) {
	os.Chdir(ProjectPath)
	err := os.RemoveAll(Target)
	if err != nil {
		log.Panicln(err)
	}
}
