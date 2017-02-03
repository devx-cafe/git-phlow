package gitwrapper

import "github.com/praqma/git-phlow/subprocess"



//GitCheckout interface
//Interface for all methods a git checkout command should poses
type GitCheckout interface {
	Checkout(ref string) (string, error)
}

type checkout struct {
	gitCheckoutCommand string
}

//NewCheckout
//Creates a new checkout command
func NewCheckout() *checkout {
	return &checkout{gitCheckoutCommand:"checkout"}
}

func (c *checkout) Checkout(ref string) (string, error) {

	coMessage, err := subprocess.SimpleExec(GitCommand, c.gitCheckoutCommand, ref)

	if err != nil {
		return "", err
	}

	return coMessage, nil

}
