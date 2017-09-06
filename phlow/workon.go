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
)

//WorkOnUpdate ...
//Type for updating issues
//returns a name or an error
type WorkOnUpdate func(key string, projectSetting *setting.ProjectSetting) (string, error)

//WorkOnCaller ...
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
			fmt.Fprintf(os.Stdout, "Whoops \nYour argument, %s, is not a number! I only accept numbers \n", keyOrID)
			os.Exit(0)
		}
		WorkOn(keyOrID, conf, UpdateGithubIssue)
		return
	}

	fmt.Println(conf.Service + "Is an unknown Service in you project .phlow file")
}

func WorkOn(keyOrID string, conf *setting.ProjectSetting, update WorkOnUpdate) {

	git := githandler.Git{Run: executor.RunGit}

	ui.PhlowSpinner.Start("Setting up workspace")
	defer ui.PhlowSpinner.Stop()

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

	fmt.Println(plugins.IssueFromBranchName(branchInfo.Current))
	//Are we already on the branch we want to work on
	if plugins.IssueFromBranchName(branchInfo.Current) == keyOrID {
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.Format.Branch(branchInfo.Current))
		return
	}
	//Does another workspace already exist with the key or ID
	for _, branch := range branchInfo.List {
		if plugins.IssueFromBranchName(branch) == keyOrID {

			if _, err = git.CheckOut(branch); err != nil {
				fmt.Println(err)
			}
			ui.PhlowSpinner.Stop()
			fmt.Fprintf(os.Stdout, "Resuming to workspace:  %s \n", ui.Format.Branch(branch))
			return
		}
	}

	//Get issue and do the transition

	name, err := update(keyOrID, conf)

	//Proceed to create the new workspace

	_, err = git.CheckOut("-b", name, conf.Remote+"/"+conf.IntegrationBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.PhlowSpinner.Stop()
	fmt.Fprintf(os.Stdout, "Created workspace:  %s \n", ui.Format.Branch(name))

}

//UpdateJIRAIssue ...
//Updates the issue on jira and returns the name of the branch
func UpdateJIRAIssue(key string, conf *setting.ProjectSetting) (string, error) {
	git := githandler.Git{Run: executor.RunGit}

	user, _ := git.Config("--get", "phlow.jirauser")
	token, _ := git.Config("--get", "phlow.jiratoken")

	//Get jira issue or fail
	issue, err := plugins.GetJiraIssue(conf.IssueURL, key, user, token)
	if err != nil {
		return "", err
	}

	var transitionErr error
	var assignErr error

	//Get transition
	transition, err := plugins.GetTransitions(conf.IssueURL, key, user, token)
	if err == nil {
		for _, tran := range transition.Transitions {
			if tran.To.Name == "In Progress" {
				transitionErr = plugins.DoTransition(conf.IssueURL, key, user, token, tran.ID)
				break
			}
		}
	}

	assignErr = plugins.AssignUser(conf.IssueURL, key, user, token)

	ui.PhlowSpinner.Stop()
	if transitionErr != nil || assignErr != nil {
		fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.Format.Issue(issue.Key))
	} else {
		fmt.Printf("\nIssue %s could not be moved to 'In Progress'\n", ui.Format.Issue(issue.Key))
		fmt.Printf("Assingee '%s' could not be assigned to issue %s\n", ui.Format.Assignee(user), ui.Format.Issue(issue.Key))
		fmt.Println(ui.Format.Bold("Go to Jira and manually set the assignee and state"))

	}

	if transitionErr != nil {
		fmt.Fprintf(os.Stdout, "Moved to => %s \n", ui.Format.Label.G4Move(plugins.PhlowLabels["In Progress"].Title))
	}

	if assignErr != nil {
		fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.Format.Assignee(user))
	}

	if transitionErr != nil || assignErr != nil {
		fmt.Println("----------------------------------")
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


	orgAndRepo := githandler.OrgAndRepo(remote)
	fmt.Println(orgAndRepo.Repository)

	issueOb, err := plugins.GetIssueGitHub(conf.IssueURL, orgAndRepo.Organisation, orgAndRepo.Repository, issue, token)
	if err != nil {
		fmt.Println("No matching issues")
		os.Exit(0)
	}

	if _, err := plugins.GitHub.SetLabel(plugins.PhlowLabels["Status - in progress"].Title, issueOb.Number); err != nil {
		fmt.Println(err)
	}

	if err := plugins.GitHub.SetAssignee(user, issueOb.Number); err != nil {
		fmt.Println(err)
	}

	is := strconv.Itoa(issueOb.Number)
	fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.Format.Issue(is))
	fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.Format.Label.G4Move(plugins.PhlowLabels["Status - in progress"].Title))
	fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.Format.Assignee(user))
	fmt.Println("----------------------------------")

	return plugins.BranchNameFromIssue(is, issueOb.Title), nil

}
