//Package phlow contains all the business logic for the workflow commands
package phlow

import (
	"bufio"
	"fmt"
	"strings"
	"syscall"

	"io"

	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/plugins"
	"github.com/code-cafe/git-phlow/setting"
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
	AuthCaller()
	fmt.Println("")
	fmt.Println(fmt.Sprintf("%s Successfully authorized: 'git phlow' is now enabled", "username"))
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
