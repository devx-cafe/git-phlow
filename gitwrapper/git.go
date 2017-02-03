package gitwrapper


//Git constants 
const (
	GitCommand string = "git"
)


//GitInterface
//interface for git commands
type GitInterface interface {
	branch() GitBranch
	Fetch() GitFetch
}

//Git
//Struct for implementation
type Git struct {

}

func (g *Git)Branch() GitBranch {
	return NewBranch()
}

func (g *Git) Fetch() GitFetch {
	return NewFetch()
}



