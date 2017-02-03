package main

import (
	"fmt"
	"github.com/praqma/git-phlow/gitwrapper"
)

func main() {

	//cmd.Execute()

	git := gitwrapper.Git{}

	branch := git.Branch()
	branches, err := branch.Branch()

	if err != nil {

		fmt.Println(err)
	}

	for _, name := range branches {
		fmt.Println(name)
	}

}