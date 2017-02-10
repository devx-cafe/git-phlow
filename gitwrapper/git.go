package gitwrapper

const (
	gitCommand string = "git"
)

type Git interface {
	Branch() Branch
	Fetch() Fetch
	Checkout() Checkout
}

type localGit struct {
	command  string //git
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

func (l *localGit) DefaultGit() *localGit {
	return &localGit{command: "git"}
}
