package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
)

//Checkouter ...
type Checkouter interface {
	Checkout(ref string) (string, error)
}

//Checkout ...
//Struct for checkout
type Checkout struct {
	baseCmd      string
	baseCheckout string
}

//NewCheckout ...
func NewCheckout(baseCMD string) *Checkout {
	return &Checkout{baseCmd: baseCMD, baseCheckout: "checkout"}
}

//Checkout ...
func (c *Checkout) Checkout(ref string) (string, error) {
	coMessage, err := subprocess.SimpleExec(c.baseCmd, c.baseCheckout, ref)
	if err != nil {
		return "", err
	}

	if len(coMessage) == 0 {
		return ref, nil
	}
	return coMessage, nil
}
