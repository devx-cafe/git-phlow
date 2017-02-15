package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
)

//Checkouter ...
type Checkouter interface {
	Checkout(ref string) (string, error)
	CheckoutNewBranchFromOrigin(string, string) (string, error)
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
	outPut, err := subprocess.SimpleExec(c.baseCmd, c.baseCheckout, ref)
	if err != nil {
		return "", err
	}

	if len(outPut) == 0 {
		//checks out from tracking origin branch
		return ref, nil
	}
	return outPut, nil
}

//CheckoutNewBranchFromOrigin
func (c *Checkout) CheckoutNewBranchFromOrigin(branchName, defaultOrigin string) (string, error) {

	message, err := subprocess.SimpleExec(c.baseCmd, c.baseCheckout, "-b", branchName, "origin/"+defaultOrigin)

	if err != nil {
		return "", err
	}

	return message, nil
}
