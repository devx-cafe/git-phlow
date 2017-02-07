package gitwrapper


//Git constants 
const (
	GitCommand string = "git"
)


//GitInterface
//interface for git commands
type Git interface {
	branch() Branch
	Fetch() Fetch
	Checkout() Checkout
}

//Git
//Struct for implementation
type gitImpl struct {

}

func (g *gitImpl)Branch() Branch {
	return NewBranch()
}

func (g *gitImpl) Fetch() Fetch {
	return NewFetch()
}

func (g *gitImpl) Checkout() Checkout {
	return NewCheckout()
}

