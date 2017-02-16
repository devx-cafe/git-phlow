package gitwrapper

//Giter ...
// interface for git
type Giter interface {
	Branch() Brancher
	Fetch() Fetcher
	Checkout() Checkouter
	Status() Statuses
}

//LocalGit ...
type LocalGit struct {
	baseGit string
}

//Branch ...
func (g *LocalGit) Branch() Brancher {
	return NewBranch(g.baseGit)
}

//Fetch ...
func (g *LocalGit) Fetch() Fetcher {
	return NewFetch(g.baseGit)
}

// Checkout ...
func (g *LocalGit) Checkout() Checkouter {
	return NewCheckout(g.baseGit)
}

func (g *LocalGit) Status() Statuses {
	return NewStatus(g.baseGit)
}

//InitGit ...
func InitGit() *LocalGit {
	return &LocalGit{baseGit: "git"}
}
