package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

const (
	all       string = "--all"
	baseFetch string = "fetch"
)

//GitFetch
//interface for git fetch commands
type Fetcher interface {
	Fetch(origin bool) (string, error)
}

type fetch struct {
	baseCMD string
}

//NewFetch
//Constructor for fetch struct
func NewFetch(baseCMD string) *fetch {
	return &fetch{baseCMD: baseCMD}

}

//Fetch
//Doing a normal git fetch
func (f *fetch) Fetch(fromOrigin bool) (string, error) {

	var message string
	var err error

	if fromOrigin {
		message, err = subprocess.SimpleExec(f.baseCMD, baseFetch, all)
	} else {
		message, err = subprocess.SimpleExec(f.baseCMD, baseFetch)
	}

	if err != nil {
		return "", err
	}

	return message, nil
}
