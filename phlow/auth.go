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

//Auth ...
//Authenticates the user with gh
func Auth() {
	token := githandler.ConfigGet("token", "phlow")
	user := githandler.ConfigGet("user", "phlow")

	if token != "" && user != "" {
		fmt.Println("you are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "enter credentials for %s \n", "github")

	//Read user input username
	username := ReadInput("username: ")
	//Read user input password
	password := ReadPassword("password: ")

	token, err := plugins.GitHub.Auth(username, password)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return
	}

	githandler.ConfigSet("token", token, "phlow")
	githandler.ConfigSet("user", username, "phlow")

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
//read input from the user anonymously
func ReadPassword(messageToUser string) string {
	fmt.Print(messageToUser)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return strings.TrimSpace(string(bytePassword))
}
