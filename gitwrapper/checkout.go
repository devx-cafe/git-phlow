package gitwrapper

import "github.com/praqma/git-phlow/subprocess"



//GitCheckout interface
//Interface for all methods a git checkout command should poses
type GitCheckout interface {
	Checkout()
}

type checkout struct {
	gitCheckoutCommand string
}

func NewCheckout() *GitCheckout {
	return &checkout{gitCheckoutCommand:"checkout"}
}

func (c *checkout) Checkout(ref string) (string, error) {

	coMessage, err := subprocess.SimpleExec(GitCommand, c.gitCheckoutCommand, ref)

	if err != nil {
		return "", err
	}

	return coMessage, nil

}
