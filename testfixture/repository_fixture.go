package main

import "fmt"
import (
	"os/user"
	"os"
)

var ProjectDirectory string = os.Getenv("GOPATH") + "/src/github.com/praqma/git-phlow"

func InitializeEmptyRepository() {

	currentUser, _ := user.Current()

	usr, _ := user.Lookup(currentUser.Name)

	fmt.Fprintln(os.Stdout, usr.HomeDir)

}

func main() {
	InitializeEmptyRepository()
}
