package githandler

import (
	"strings"

	"github.com/praqma/git-phlow/executor"
	"bytes"
	"fmt"
)

//Git ...
//Git object for execution operations on the local git installation
type Git struct {
	Run executor.GitCommandRunner
}

//LSRemote ...
//Executes local git ls-remote with params
func (os *Git) LSRemote(argv ...string) (string, error) {
	return os.Run("git", "ls-remote", argv...)
}

//Branch ...
//Executes local git branch with params
func (os *Git) Branch(argv ...string) (string, error) {
	return os.Run("git", "branch", argv...)
}

//CheckOut ...
//Executes local git checkout with params
func (os *Git) CheckOut(argv ...string) (string, error) {
	return os.Run("git", "checkout", argv...)
}

//Add ...
//Executes local git add with params
func (os *Git) Add(argv ...string) (string, error) {
	return os.Run("git", "add", argv...)
}

//Commit ...
//Executes local git commit with params
func (os *Git) Commit(argv ...string) (string, error) {
	return os.Run("git", "commit", argv...)
}

//Fetch ...
//Executes local git fetch with params
func (os *Git) Fetch(argv ...string) (string, error) {
	return os.Run("git", "fetch", argv...)
}

//Pull ...
//Executes local git pull with params
func (os *Git) Pull(argv ...string) (string, error) {
	return os.Run("git", "pull", argv...)
}

//Push ...
//Executes local git push with params
func (os *Git) Push(argv ...string) (string, error) {
	return os.Run("git", "push", argv...)
}

//Status ...
//Executes local git status with params
func (os *Git) Status(argv ...string) (string, error) {
	return os.Run("git", "status", argv...)
}

//Merge ...
//Executes local git merge with params
func (os *Git) Merge(argv ...string) (string, error) {
	return os.Run("git", "merge", argv...)
}

//Config ...
//Executes local git config with params
func (os *Git) Config(argv ...string) (string, error) {
	stdOut, stdErr := os.Run("git", "config", argv...)
	if stdErr != nil {
		return "", stdErr
	}
	return strings.Replace(stdOut, "\n", "", -1), nil
}

//DEPRECATESD SECTION ---------------------------------------------------------------
//FormatPatch ...
//dry runs patch to see if we can auto merge
func FormatPatch(buf *bytes.Buffer, remoteBranch string) (err error) {
	//err = executor.ExecPipeCommand(buf,
	//	exec.Command("git", "format-patch", remoteBranch, "--stdout"),
	//	exec.Command("git", "apply", "check"),
	//)
	return
}

//StatusPorcelain ...
//generates behind and ahead status
func StatusPorcelain() (string, error) {
	out, err := executor.RunCommand("git", "status", "short", "--branch", "--porcelain")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

//PushRename ...
func PushRename(branch, defaultBranch string) (string, error) {
	remote := ConfigBranchRemote(defaultBranch)
	str := fmt.Sprintf("%s:ready/%s", branch, branch)

	return executor.RunCommand("git", "push", remote, str)
}
