package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
	"strings"
	"bytes"
)

type Brancher interface {
	ListBranches() ([]string, error)
	CreateBranch(name string) (string, error)
}

type branch struct {
	baseCMD    string
	baseBranch string
}

//NewBranch
//Constructor for branch struct
func NewBranch(baseCMD string) *branch {
	return &branch{baseCMD: baseCMD, baseBranch: "branch"}
}

//Branch
//Get list of all branches: equals "git branch"
func (b *branch) ListBranches() ([]string, error) {
	output, err := subprocess.SimpleExec(b.baseCMD, b.baseBranch)
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

	_, err := subprocess.SimpleExec(b.baseCMD, b.baseBranch, name)

	if err != nil {
		return "", err
	}

	return name, nil
}

func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
