package gitwrapper

import (
	"bytes"
	"errors"
	"strings"

	"github.com/praqma/git-phlow/subprocess"
)

//Brancher ...
//interface for branch methods
type Brancher interface {
	ListBranches() ([]string, error)
	CreateBranch(name string) (string, error)
	CurrentBranch() (string, error)
}

//Branch ...
type Branch struct {
	baseCMD    string
	baseBranch string
}

//NewBranch ...
//Constructor for branch struct
func NewBranch(baseCMD string) *Branch {
	return &Branch{baseCMD: baseCMD, baseBranch: "branch"}
}

//ListBranches ...
//Get list of all branches: equals "git branch"
func (b *Branch) ListBranches() ([]string, error) {
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

//CreateBranch ...
//Create a new branch: equals "git branch [name]"
func (b *Branch) CreateBranch(name string) (string, error) {

	_, err := subprocess.SimpleExec(b.baseCMD, b.baseBranch, name)

	if err != nil {
		return "", err
	}

	return name, nil
}

//CurrentBranch ...
//Get the currently selected branch
func (b *Branch) CurrentBranch() (string, error) {
	var symbolic, short, head string = "symbolic-ref", "--short", "HEAD"

	branch, err := subprocess.SimpleExec(b.baseCMD, symbolic, short, head)

	if err != nil {
		return "", err
	}

	if len(branch) == 0 {
		return "", errors.New("error disconnected from branch")
	}

	return strings.TrimSpace(branch), nil
}

func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
