package main

import (
	"fmt"
	"github.com/praqma/git-phlow/subprocess"
)

func main() {

	//cmd.Execute()


	str, err := subprocess.SimpleExec("git", "sts")

	fmt.Println(str)
	if err != nil {

	fmt.Println(err)

	}

	//git := gitwrapper.Git{}
	//
	//branch := git.Branch()
	//branches, err := branch.Branch()
	//
	//if err != nil {
	//
	//	fmt.Println("Branchy err: " + err.Error())
	//}
	//
	//for _, name := range branches {
	//	fmt.Println("branch message: " + name)
	//}
	//
	//fetchCommand := git.Checkout()
	//msg, err := fetchCommand.Checkout("master")
	//
	//if err != nil {
	//
	//	fmt.Println("Fetch err: " + err.Error())
	//}
	//
	//fmt.Println("Fetch messsage: " + msg)
}