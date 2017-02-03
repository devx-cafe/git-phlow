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

		fmt.Println("Branchy err: " + err.Error())
	}

	for _, name := range branches {
		fmt.Println("branch message: " + name)
	}

	fetchCommand := git.Fetch()
	msg, err := fetchCommand.Fetch(true)

	if err != nil {

		fmt.Println("Fetch err: " + err.Error())
	}

	fmt.Println("Fetch messsage: " + msg)
}