package phlow

import (
	"bufio"
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"

	"strings"
	"syscall"

	"github.com/praqma/git-phlow/plugins"
	"golang.org/x/crypto/ssh/terminal"
)

/*
   Reauthorize should also be possible
*/

//Enable ...
func Enable() {
	token, tErr := githandler.Config("token", "", "phlow", true)
	user, uErr := githandler.Config("user", "", "phlow", true)

	if (tErr == nil && token != "") || (uErr == nil && user != "") {
		fmt.Println("you are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "enter credentials for %s \n", "github")
	//Read user input username
	username := ReadInput("username: ")
	//Read user input password

	password := ReadPassword("password: ")

	token, err := plugins.Authorize(username, password, plugins.AuthURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	githandler.Config("token", token, "phlow", false)
	githandler.Config("user", username, "phlow", false)

	fmt.Println("Success fully authorized: 'git phlow' is now enabled  ")
}

//ReadInput ...
//Reads input from user
func ReadInput(messageToUser string) string {
	fmt.Print(messageToUser)
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')

	return strings.Replace(text, "\n", "", -1)
}

//ReadPassword ...
func ReadPassword(messageToUser string) string {
	fmt.Print(messageToUser)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return strings.TrimSpace(string(bytePassword))
}
