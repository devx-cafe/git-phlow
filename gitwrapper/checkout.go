package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
)

//GitCheckout interface
//Interface for git checkout
type Checkouter interface {
	Checkout(ref string) (string, error)
}

type checkout struct {
	baseCmd      string
	baseCheckout string
}

//NewCheckout
//Creates a new checkout baseGit
func NewCheckout(baseCMD string) *checkout {
	return &checkout{baseCmd: baseCMD, baseCheckout: "checkout"}
}

func (c *checkout) Checkout(ref string) (string, error) {

	coMessage, err := subprocess.SimpleExec(c.baseCmd, c.baseCheckout, ref)

	if err != nil {
		return "", err
	}

	if len(coMessage) == 0 {
		coMessage = ref
	}
	return ref, nil

}
