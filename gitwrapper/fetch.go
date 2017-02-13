package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

//GitFetch
//interface for git fetch commands
type Fetcher interface {
	FetchFromOrigin() (string, error)
	HasRemote() bool
}

type fetch struct {
	baseCMD    string
	baseFetch  string
	originFlag string
}

//NewFetch
//Constructor for fetch struct
func NewFetch(baseCMD string) *fetch {
	return &fetch{baseCMD: baseCMD, baseFetch: "fetch", originFlag: "origin"}

}

//Fetch
//Doing a normal git fetch
func (f *fetch) FetchFromOrigin() (string, error) {

	_, err := subprocess.SimpleExec(f.baseCMD, f.baseFetch, f.originFlag)
	if err != nil {
		return "", err
	}

	return "Fetcing from origin", nil
}

//HasRemote
//Verifies if repository has a remote 
func (f *fetch) HasRemote() bool {
	if _, err := subprocess.SimpleExec(f.baseCMD, "ls-remote"); err != nil {
		return false
	}
	return true
}
