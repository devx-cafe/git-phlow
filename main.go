package main

import (
	"github.com/praqma/git-phlow/gitwrapper"
	"fmt"
	"strconv"
	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	cmd.Execute()

	var branch *gitwrapper.Branch
	branch = &gitwrapper.Branch{}
	_, err := branch.Branch()

	if err != nil{
		fmt.Println(err)
	}


	for i, val := range branch.Branches {
		fmt.Println(strconv.Itoa(i) + " bbbb " + val)
	}

}