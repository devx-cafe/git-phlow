package phlowimpl

import (
	"errors"
	"strconv"
	"github.com/praqma/git-phlow/gitwrapper"
	"fmt"
)

func Workon(args []string) (message string, err error) {

	var (
		parseError error
		issue int64
		name string = "sd"
	)


	if len(args) != 0  {

		issue, parseError = strconv.ParseInt(args[0], 0, 64)

		if parseError != nil {
			return "", errors.New("argument must be of type [number]")
		}


	fmt.Println(issue)

		gitwrapper.CheckoutNewBranch(name)

	}

	return "", errors.New("missing workon [number]")
}




