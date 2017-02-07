package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
	"strings"
	"bytes"
)

type Branch interface {
	Branch() ([]string, error)
}

//Branch
type branch struct {
	gitBranchCommand string
	Branches         []string
}

//NewBranch
//Constructor for branch struct
func NewBranch() *branch {
	return &branch{gitBranchCommand:"branch"}
}


//Branch
//Get list of all branches: equals "git branch"
func (b *branch) Branch() ([]string, error) {
	output, err := subprocess.SimpleExec(GitCommand, b.gitBranchCommand)
	if err != nil {
		return nil, err
	}

	for _, branch := range strings.Split(output, "\n") {
		if branch != "" {
			b.Branches = append(b.Branches, branch)
		}
	}
	return b.Branches, nil
}

//CreateBranch
//Create a new branch: equals "git branch [name]"
func (b *branch)CreateBranch(name string) (string, error) {

	_, err := subprocess.SimpleExec(GitCommand, b.gitBranchCommand, name)

	if err != nil {
		return "", err
	}

	return efficientConcatString("branch '", name, "' created"), nil
}

func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}