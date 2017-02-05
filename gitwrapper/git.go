package gitwrapper


//Git constants 
const (
	GitCommand string = "git"
)


//GitInterface
//interface for git commands
type GitInterface interface {
	branch() Branch
	Fetch() Fetch
	Checkout() Checkout
}

//Git
//Struct for implementation
type Git struct {

}

func (g *Git)Branch() Branch {
	return NewBranch()
}

func (g *Git) Fetch() Fetch {
	return NewFetch()
}

func (g *Git) Checkout() Checkout {
	return NewCheckout()
}

