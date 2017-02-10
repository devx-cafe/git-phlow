package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
	"strings"
	"bytes"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

type Brancher interface {
	ListBranches() ([]string, error)
	CreateBranch(name string) (string, error)
}

const (
	baseBranch string = "branch"
)

type branch struct {
	cmd      string
}

//NewBranch
//Constructor for branch struct
func NewBranch(baseCMD string) *branch {
	return &branch{cmd: baseCMD}
}

//Branch
//Get list of all branches: equals "git branch"
func (b *branch) ListBranches() ([]string, error) {
	output, err := subprocess.SimpleExec(b.cmd, baseBranch)
	if err != nil {
		return nil, err
	}

	var branches []string
	for _, branch := range strings.Split(output, "\n") {
		if branch != "" {
			branches = append(branches, branch)
		}
	}
	return branches, nil
}

//CreateBranch
//Create a new branch: equals "git branch [name]"
func (b *branch) CreateBranch(name string) (string, error) {

	_, err := subprocess.SimpleExec(b.cmd, baseBranch, name)

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
