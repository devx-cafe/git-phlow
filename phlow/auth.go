package phlow

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"io"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/setting"
)

//AuthCaller
//Wraps auth and injects dependencies
func AuthCaller() {
	cf := setting.GitConfig{Run: executor.Run}
	Auth(cf)
}

//Auth ...
//Authenticates the user
func Auth(cf setting.Configurator) {
	token := cf.Get(setting.PhlowToken)
	user := cf.Get(setting.PhlowUser)

	if token != "" && user != "" {
		fmt.Println("Checking token validity...")
		isAuthenticated, err := plugins.GitHub.CheckAuth()
		if !isAuthenticated {
			fmt.Println("Token test expected HTTP code 200 but received " + err.Error())
			if ReadInput("Delete local token and reauthenticate? (y/n): ", os.Stdin) == "y" {
				fmt.Println("Deleting local token and reauthenticating...")
				cf.Unset(setting.PhlowToken)
				cf.Unset(setting.PhlowUser)
				AuthCaller()
			} else {
				fmt.Println("Aborting...")
			}
			return
		}
		fmt.Println("Token successfully validated. You are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "Enter credentials for %s \n", "GitHub")

	//Read user input username
	username := ReadInput("username: ", os.Stdin)
	//Read user input password
	password := ReadPassword("password: ")

	token, err := plugins.GitHub.Auth(username, password)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return
	}
	cf.Set(setting.PhlowUser, username)
	cf.Set(setting.PhlowToken, token)

	fmt.Println("")
	fmt.Println(fmt.Sprintf("%s Successfully authorized: 'git phlow' is now enabled", username))
}

//ReadInput ...
//Reads input from user
func ReadInput(messageToUser string, input io.Reader) string {
	fmt.Print(messageToUser)

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	return scanner.Text()
}

//ReadPassword ...
//read input from the user anonymously
func ReadPassword(messageToUser string) string {
	fmt.Print(messageToUser)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return strings.TrimSpace(string(bytePassword))
}
