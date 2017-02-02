package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
	"strings"
	"bytes"
)


//Branch
type Branch struct {
	gitBranchCommand string
	Branches         []string
}

//NewBranch
//initializes a branch struct
func NewBranch() *Branch {
	return &Branch{gitBranchCommand:"branch"}
}


//Branch
//Get list of all branches
func (b *Branch) Branch() ([]string, error) {
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
//Create a new branch
func (b *Branch)CreateBranch(name string) (string, error) {

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