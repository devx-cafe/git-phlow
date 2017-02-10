package gitwrapper

const (
	gitCommand string = "git"
)

type Giter interface {
	Branch() Brancher
	Fetch() Fetch
	Checkout() Checkout
}

type localGit struct {
	command  string
	baseArgs string
}

func (g *localGit) Branch() Brancher {
	return NewBranch(gitCommand)
}

func (g *localGit) Fetch() Fetch {
	return NewFetch()
}

func (g *localGit) Checkout() Checkout {
	return NewCheckout()
}

func InitGit() *localGit {
	return &localGit{command: "git"}
}
