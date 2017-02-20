package phlow

import (
	"bufio"
	"fmt"
	"os"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
)

/*
   Reauthorize should also be possible
*/

//Enable ...
func Enable() {
	//Run status before check

	token, tErr := githandler.Config("token", "", true)
	user, uErr := githandler.Config("user", "", true)

	if (tErr == nil && token != "") || (uErr == nil && user != "") {
		fmt.Println("you are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "enter credentials for %s \n", "github")
	//Read user input username
	username := ReadInput("username: ")
	//Read user input password
	password := ReadInput("password: ")

	if token, err := plugins.Authorize(username, password, plugins.AuthURL); err != nil {
		githandler.Config("token", token, false)
		githandler.Config("user", username, false)
		fmt.Println("Success fully authorized: 'git do' is now enabled  ")
	}
}

//ReadInput ...
//Reads input from user
func ReadInput(messageToUser string) string {
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')

	return text
}
