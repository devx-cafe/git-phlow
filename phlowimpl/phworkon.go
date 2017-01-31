package phlowimpl

import (
	"errors"
	"strconv"
	"github.com/praqma/git-phlow/gitwrapper"
)

var (
	name string = "issue-feature-branch"
	empty string = ""
)

func WorkOn(args []string) (string, error) {

	if len(args) == 0 {
		return empty, errors.New("argument must be of type [number]")
	}

	issue, parseError := strconv.ParseInt(args[0], 0, 32)

	if parseError != nil {
		return empty, errors.New("argument must be of type [number]")
	}

	name = ConvertToBranchName(strconv.Itoa(int(issue)) + "-" + name)
	err := gitwrapper.CheckoutNewBranch(name)

	if err != nil {
		return empty, err
	}

	return "Switched to a new branch " + "'" + name + "'", nil
}




