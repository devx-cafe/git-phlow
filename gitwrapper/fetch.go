package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

//Fetcher ...
type Fetcher interface {
	FetchFromOrigin() (string, error)
	HasRemote() bool
}

//Fetch ...
type Fetch struct {
	baseCMD    string
	baseFetch  string
	originFlag string
}

//NewFetch ...
//Constructor for fetch struct
func NewFetch(baseCMD string) *Fetch {
	return &Fetch{baseCMD: baseCMD, baseFetch: "fetch", originFlag: "origin"}

}

//FetchFromOrigin ...
//Doing a normal git fetch
func (f *Fetch) FetchFromOrigin() (string, error) {

	_, err := subprocess.SimpleExec(f.baseCMD, f.baseFetch, f.originFlag)
	if err != nil {
		return "", err
	}

	return "Fetcing from origin", nil
}

//HasRemote ...
//Verifies if repository has a remote
func (f *Fetch) HasRemote() bool {
	if _, err := subprocess.SimpleExec(f.baseCMD, "ls-remote"); err != nil {
		return false
	}
	return true
}
