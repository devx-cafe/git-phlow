package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

const baseCheckout string = "checkout"

//GitCheckout interface
//Interface for git checkout
type Checkouter interface {
	Checkout(ref string) (string, error)
}

type checkout struct {
	baseCmd string
}

//NewCheckout
//Creates a new checkout baseGit
func NewCheckout(baseCMD string) *checkout {
	return &checkout{baseCmd: baseCMD}
}

func (c *checkout) Checkout(ref string) (string, error) {

	coMessage, err := subprocess.SimpleExec(c.baseCmd, baseCheckout, ref)

	if err != nil {
		return "", err
	}

	return coMessage, nil

}
