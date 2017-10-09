package phlow

import (
	"fmt"
	"os"

	"strconv"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/setting"
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/options"
	"strings"
	"errors"
)

//WorkOnUpdate ...
//Type for updating issues
//returns a name or an error
type WorkOnUpdate func(key string, projectSetting *setting.ProjectSetting) (string, error)

//WorkOnCaller ...
//Toplevel workon function called from cmd
func WorkOnCaller(keyOrID string) {

	INIBlock := options.GlobalFlagTarget
	conf := setting.NewProjectStg(INIBlock)

	if "jira" == strings.ToLower(conf.Service) {
		WorkOn(keyOrID, conf, UpdateJIRAIssue)
		return
	}

	if "github" == strings.ToLower(conf.Service) {
		_, err := strconv.Atoi(keyOrID)
		if err != nil {
			fmt.Println("Whoops \nYour argument, " + keyOrID + ", is not a valid GitHub issue number \n")
			os.Exit(0)
		}
		WorkOn(keyOrID, conf, UpdateGithubIssue)
		return
	}

	fmt.Println(conf.Service + "Is an unknown Service in you project .phlow file")
}

//WorkOn ...
//creates a new workspace from issue by given WorkOnUpdate and configuration
func WorkOn(keyOrID string, conf *setting.ProjectSetting, update WorkOnUpdate) {

	git := githandler.Git{Run: executor.RunGit}

	fmt.Println("Preparing workspace...")

	if _, err := git.Fetch("--all"); err != nil {
		fmt.Println(err)
		return
	}

	out, err := git.Branch("-a")
	if err != nil {
		fmt.Println("Could not get branches")
		return
	}

	branchInfo := githandler.AsList(out)

	//Are we already on the branch we want to work on
	if strings.HasPrefix(branchInfo.Current, keyOrID) {
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.Format.Branch(branchInfo.Current))
		return
	}
	//Does another workspace already exist with the key or ID
	for _, branch := range branchInfo.List {
		if strings.HasPrefix(branch, keyOrID) {

			if _, err = git.CheckOut(branch); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintf(os.Stdout, "Resuming to workspace:  %s \n", ui.Format.Branch(branch))
			return
		}
	}

	name, err := update(keyOrID, conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//ERROR CHECK HERE
	_, err = git.CheckOut("-b", name, conf.Remote+"/"+conf.IntegrationBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Created workspace:  %s \n", ui.Format.Branch(name))
}

//UpdateJIRAIssue ...
//Updates the issue on jira and returns the name of the branch
func UpdateJIRAIssue(key string, conf *setting.ProjectSetting) (string, error) {
	git := githandler.Git{Run: executor.RunGit}

	user, _ := git.Config("--get", "phlow.jirauser")
	token, _ := git.Config("--get", "phlow.jiratoken")

	issue, err := plugins.GetJiraIssue(conf.IssueApi, key, user, token)
	if err != nil {
		return "", err
	}

	var transitionErr error
	assignErr := plugins.AssignUser(conf.IssueApi, key, user, token)

	//Get transition
	transition, err := plugins.GetTransitions(conf.IssueApi, key, user, token)
	if err == nil {
		for _, tran := range transition.Transitions {
			if tran.To.StatusCategory.Name == "In Progress" {
				transitionErr = plugins.DoTransition(conf.IssueApi, key, user, token, tran.ID)

				break
			}
		}
		transitionErr = errors.New("No 'In Progress' transition ")
	}

	fmt.Printf("\n-------- Issue %s--------  \n", ui.Format.Issue(issue.Key))

	if transitionErr != nil {
		fmt.Printf("Issue %s could not be moved to 'In Progress'\n", ui.Format.Issue(issue.Key))
	} else {
		fmt.Printf("Moved to => %s \n", ui.Format.Label.G4Move("In Progress"))
	}

	if assignErr != nil {
		fmt.Printf("Assingee '%s' could not be assigned to issue %s\n", ui.Format.Assignee(user), ui.Format.Issue(issue.Key))
	} else {
		fmt.Printf("Assignee => %s \n", ui.Format.Assignee(user))
	}

	return plugins.BranchNameFromIssue(issue.Key, issue.Fields.Summary), nil
}

//UpdateGithubIssue ...
//Updating an issue on github and returns the branch name
func UpdateGithubIssue(issue string, conf *setting.ProjectSetting) (string, error) {
	git := githandler.Git{Run: executor.RunGit}

	token, _ := git.Config("--get", "phlow.token")
	user, _ := git.Config("--get", "phlow.user")

	//Get organisation and repository for repository
	remote, err := git.LSRemote("--get-url", conf.Remote)
	if err != nil {
		panic(err)
	}

	oap := githandler.OrgAndRepo(remote)

	issueOb, err := plugins.GetIssueGitHub(conf.IssueApi, oap.Organisation, oap.Repository, issue, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err := plugins.SetAssigneeGitHub(conf.IssueApi, oap.Organisation, oap.Repository, token, issue, user); err != nil {
		fmt.Println(err)
	}

	if _, err := plugins.SetLabelGitHub(conf.IssueApi, oap.Organisation, oap.Repository, token, plugins.PhlowLabels["Status - in progress"].Title, issue); err != nil {
		fmt.Println(err)
	}

	is := strconv.Itoa(issueOb.Number)
	fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.Format.Issue(is))
	fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.Format.Label.G4Move(plugins.PhlowLabels["Status - in progress"].Title))
	fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.Format.Assignee(user))
	fmt.Println("----------------------------------")

	return plugins.BranchNameFromIssue(is, issueOb.Title), nil

}
