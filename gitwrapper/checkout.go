package gitwrapper

import (
	"github.com/libgit2/git2go"
)

type WrapperCommand interface {
	ExecutePhlowImplementation()
}


//CheckoutNewBranch
func CheckoutNewBranch(name string) (err error) {

	err = nil

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
	_, err = repo.CreateBranch(name, headCommit, false)

	if err != nil {
		return
	}

	_, err = repo.References.CreateSymbolic("HEAD", "refs/heads/" + name, true, "headOne")

	if err != nil {
		return

	}
	opts := &git.CheckoutOpts{
		Strategy: git.CheckoutSafe | git.CheckoutRecreateMissing,
	}

	if err = repo.CheckoutHead(opts); err != nil {
		return
	}
	return
}