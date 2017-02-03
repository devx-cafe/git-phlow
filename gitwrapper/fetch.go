package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

const (
	origin string = "origin"
	all string = "--all"
)

//GitFetch
//wrapper interface for gitcommands
type GitFetch interface {
	Fetch(origin bool) (string, error)
}

type fetch struct {
	gitFetchCommand string
}


//NewFetch
//Gives a new fetch command
func NewFetch() *fetch {
	return &fetch{gitFetchCommand:"fetch"}

}

//Fetch
//Doing a normal git fetch
func (f *fetch) Fetch(fromOrigin bool) (string, error) {

	var message string
	var err error

	if fromOrigin {
		message, err = subprocess.SimpleExec(GitCommand, f.gitFetchCommand, origin)
	} else {
		message, err = subprocess.SimpleExec(GitCommand, f.gitFetchCommand)
	}

	if err != nil {
		return "", err
	}

	return message, nil
}
