package gitwrapper

import ()

const (
	gitCommand  string = "git"
	gitFromPath string = "-C"
)

type Git interface {
	Branch() Branch
	Fetch() Fetch
	Checkout() Checkout
}

type localGit struct {
	command  string
	baseArgs string
}

func (g *localGit) Branch() Branch {
	return NewBranch()
}

func (g *localGit) Fetch() Fetch {
	return NewFetch()
}

func (g *localGit) Checkout() Checkout {
	return NewCheckout()
}

func (l localGit) GitFromPath(path string) *localGit {
	args := gitFromPath + " " + path
	return &localGit{command: gitCommand, baseArgs: args}
}

func (l localGit) DefaultGit() *localGit {
	return &localGit{command: "git"}
}
