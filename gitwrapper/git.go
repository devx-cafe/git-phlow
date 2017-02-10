package gitwrapper

type Giter interface {
	Branch() Brancher
	Fetch() Fetcher
	Checkout() Checkouter
}

type localGit struct {
	baseGit string
}

func (g *localGit) Branch() Brancher {
	return NewBranch(g.baseGit)
}

func (g *localGit) Fetch() Fetcher {
	return NewFetch(g.baseGit)
}

func (g *localGit) Checkout() Checkouter {
	return NewCheckout(g.baseGit)
}

func InitGit() *localGit {
	return &localGit{baseGit: "git"}
}
