package context

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/code-cafe/git-phlow/executor"
)

// WorkspaceContext ...
var WorkspaceContext Context

// Context ...
type Context struct {
	DeliveredBranches       []string
	RemoteDeliveredBranches []string
	Workspaces              []string
	Branches                []string
	CurrentBranch           string
	Remote                  string
	Organisation            string
	Repository              string
	Token                   string
	User                    string
}

func init() {

	WorkspaceContext.Branches = GetBranches(func() string {
		out, err := executor.Run("git", "branch")
		if err != nil {
			fmt.Println(out)
			os.Exit(1)
		}
		return out
	}())

	deliveredChn := make(chan []string, 1)
	workspaceChn := make(chan []string, 1)
	orgAndRepoChn := make(chan string, 2)
	tokenChn := make(chan string, 1)
	userChn := make(chan string, 1)

	go func() {
		deliveredChn <- GetDelivered(WorkspaceContext.Branches)
	}()

	go func() {
		workspaceChn <- GetWorkSpaces(WorkspaceContext.Branches)
	}()

	go GetToken(tokenChn, func() (s string, err error) {
		return executor.Run("git", "config", "--get", "phlow.token")
	})

	go GetUser(userChn, func() (s string, err error) {
		return executor.Run("git", "config", "--get", "phlow.user")
	})

	go func() {
		org, repo := GetOrganizationAndRepository(func() string {
			out, err := executor.Run("git", "ls-remote", "--get-url", "origin")
			if err != nil {
				panic(err)
			}
			return out
		}())

		orgAndRepoChn <- org
		orgAndRepoChn <- repo

	}()

	WorkspaceContext.Token = <-tokenChn
	WorkspaceContext.User = <-userChn
	WorkspaceContext.DeliveredBranches = <-deliveredChn
	WorkspaceContext.Workspaces = <-workspaceChn
	WorkspaceContext.Organisation = <-orgAndRepoChn
	WorkspaceContext.Repository = <-orgAndRepoChn
}

// GetBranches ...
// get a list of branches -
func GetBranches(out string) (branches []string) {
	unfilteredBranches := strings.Split(out, "\n")

	for _, str := range unfilteredBranches {
		str = strings.ReplaceAll(str, " ", "")

		if strings.HasPrefix(str, "*") {
			str = strings.TrimPrefix(str, "*")
		}

		if str != "" {
			branches = append(branches, str)
		}
	}

	return
}

// GetDelivered ...
// returns all branches prefixed with delivered
func GetDelivered(branches []string) (delivered []string) {
	for _, branch := range branches {
		if strings.HasPrefix(branch, "delivered/") {
			delivered = append(delivered, branch)
		}

	}

	return
}

// GetWorkSpaces ...
// return all workspace branches prefixed with an issue number
func GetWorkSpaces(branches []string) (workspaces []string) {
	reg := regexp.MustCompile(`\d+`)

	for _, branch := range branches {
		// Find index returns a slice, if the slice is not empty the first element must be a number
		if indx := reg.FindIndex([]byte(branch)); len(indx) > 0 && indx[0] == 0 {
			workspaces = append(workspaces, branch)
		}
	}

	return
}

// GetCurrent ...
// gets the current branch
func GetCurrent() string {
	out, err := executor.Run("git", "rev-parse --abbrev-ref HEAD")
	if err != nil {
		panic(err)
	}
	return out
}

// GetOrganizationAndRepository ...
// Get repository name and organization name from remotes
func GetOrganizationAndRepository(url string) (string, string) {

	//Extracts repo and org from ssh url format
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)

	//if there is no protocol prefix we the two capture groups will fetch
	//organization and repo
	if strings.HasPrefix(url, "git@") {
		match := re.FindStringSubmatch(url)
		return match[1], match[2]
	}
	//Is the protocol prefix is http we split the urls backwards and remove .git
	if strings.HasPrefix(url, "http") {
		splitURL := strings.Split(strings.TrimSuffix(url, ".git"), "/")
		org := splitURL[len(splitURL)-2]
		repo := splitURL[len(splitURL)-1]
		return org, repo
	}

	//If any protocol is given (apart from file:// that will fail) the capture groups
	if strings.Contains(url, "://") {
		protoExp := regexp.MustCompile(`\w*:\/\/[\w.@]+(?:\:\d+)?\/(\w+)\/(\w+)\.git`)
		match := protoExp.FindStringSubmatch(url)
		return match[1], match[2]
	}

	panic(`no "remote origin" found`)
}

// GetUser ...
// Get the name of the user
func GetUser(chn chan string, fn func() (string, error)) {
	str, err := fn()
	if err != nil {
		log.Fatalln(err)
	}
	chn <- strings.TrimSpace(str)
}

// GetToken ...
// get the OAuth token
func GetToken(chn chan string, fn func() (string, error)) {
	str, err := fn()
	if err != nil {
		log.Fatalln(err)
	}
	chn <- strings.TrimSpace(str)
}
