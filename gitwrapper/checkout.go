package gitwrapper

import (
	"os"

	"github.com/libgit2/git2go"
	"fmt"
)

type WrapperCommand interface {
	ExecutePhlowImplementation()
}



//CheckoutNewBranch creates a new branch and checks it out
func CheckoutNewBranch(name string) {

	var err error

	path, err := GetCurrentDirectory()

	if err != nil {
		return
	}

	repo, err := git.OpenRepository(path)

	if err != nil {
		return
	}

	head, err := repo.Head()
	headCommit, err := repo.LookupCommit(head.Target())

	if err != nil {
		return
	}

	//Creates branch no-force
	 br, err := repo.CreateBranch(name, headCommit, false)

	fmt.Println(br)
	if err != nil {
		os.Exit(1)
	}

}