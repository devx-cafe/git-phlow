package gitwrapper

import (
	"github.com/praqma/git-phlow/subprocess"
	"fmt"
)

const (
	gitCommand string = ""
	flag_C     string = ""
)

//GitInterface
//interface for git commands
type Git interface {
	Branch() Branch
	Fetch() Fetch
	Checkout() Checkout
	NewConfig(git gitConfig) *gitImpl
}

//Git
//Struct for implementation
type gitImpl struct {
}

type gitConfig struct {
	directory string
}

func (g *gitImpl) Branch() Branch {
	return NewBranch()
}

func (g *gitImpl) Fetch() Fetch {
	return NewFetch()
}

func (g *gitImpl) Checkout() Checkout {
	return NewCheckout()
}

func (g *gitImpl) NewConfig(config gitConfig) *gitImpl {
	return g
}


func someone(git Git) {

}
