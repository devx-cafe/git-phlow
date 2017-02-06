package main

import (
	"fmt"
	"github.com/praqma/git-phlow/subprocess"
	"strconv"
	"os"
)

func main() {

	//cmd.Execute()


	_, err := subprocess.SimpleExec("git", "sts")


	if err != nil {

		str := err.(subprocess.ExecError)

		fmt.Println("ARGH" + str.StdErr)
		fmt.Println("exit code: " + strconv.Itoa(str.ExitCode))
		os.Exit(str.ExitCode)
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