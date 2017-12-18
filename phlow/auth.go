//Package phlow contains all the business logic for the workflow commands
package phlow

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"io"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/setting"
	"golang.org/x/crypto/ssh/terminal"
)

//AuthCaller ...
//Wraps auth and injects dependencies
func AuthCaller() {
	INIBlock := options.GlobalFlagTarget
	conf := setting.NewProjectStg(INIBlock)
	if "jira" == strings.ToLower(conf.Service) {
		Auth(INIBlock, plugins.AuthorizeJIRA, plugins.AuthenticateJIRA, "phlow.jirauser", "phlow.jiratoken", conf.Service)
	} else if "github" == strings.ToLower(conf.Service) {
		Auth(INIBlock, plugins.AuthorizeGitHub, plugins.AuthenticateGitHub, "phlow.user", "phlow.token", conf.Service)
	} else {
		fmt.Println(conf.Service + "Is an unknown Service in you project .phlow file")
	}
}

//Auth ...
//Authenticates the desired service
func Auth(INIBlock string, authorization plugins.Authorization, authentication plugins.Authentication, configUser string, configToken string, service string) {
	conf := setting.NewProjectStg(INIBlock)
	git := githandler.Git{Run: executor.RunGit}

	token, err := git.Config("--get", configToken)
	user, err := git.Config("--get", configUser)

	if token != "" && user != "" {
		fmt.Printf("Checking token validity for %s... \n", service)
		isAuthenticated, err := authentication(conf.IssueApi, user, token)
		if !isAuthenticated {
			fmt.Println("Token test expected HTTP code 200 but received " + err.Error())
			if ReadInput("Delete local token and reauthenticate? (y/n): ", os.Stdin) == "y" {
				fmt.Println("Deleting local token and reauthenticating...")
				git.Config("--global", "--unset", configUser)
				git.Config("--global", "--unset", configToken)
				Auth(conf.INIBlock, authorization, authentication, configUser, configToken, service)
			} else {
				fmt.Println("Aborting...")
			}
			return
		}
		fmt.Println("Token successfully validated. You are already signed in")
		return
	}

	fmt.Fprintf(os.Stdout, "Enter credentials for %s \n", service)
	//Read user input username
	username := ReadInput("username: ", os.Stdin)
	//Read user input password
	password := ReadPassword("password: ")

	token, err = authorization(conf.IssueApi, username, password)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return
	}
	_, err = git.Config("--global", configUser, username)
	_, err = git.Config("--global", configToken, token)

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
