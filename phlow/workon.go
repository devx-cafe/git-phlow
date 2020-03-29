package phlow

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/code-cafe/git-phlow/context"

	"strconv"

	"strings"

	. "github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/plugins"
	"github.com/code-cafe/git-phlow/setting"
	"github.com/code-cafe/git-phlow/ui"
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

	if "zenhub" == conf.Service {
		fmt.Println("Zenhub is not supported yet")
	}

	if "github" == strings.ToLower(conf.Service) {
		_, err := strconv.Atoi(keyOrID)
		if err != nil {
			fmt.Println("Whoops \nYour argument, " + keyOrID + ", is not a valid GitHub issue number \n")
			os.Exit(1)
		}
		WorkOn(keyOrID, conf, UpdateGithubIssue)
		return
	}

	fmt.Println(conf.Service + "Is an unknown Service in you project .phlow file")
}

//WorkOn ...
//creates a new workspace from issue by given WorkOnUpdate and configuration
func WorkOn(keyOrID string, conf *setting.ProjectSetting, update WorkOnUpdate) {

	ctx := context.WorkspaceContext

	fmt.Println("Preparing workspace...")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	defer s.Stop()

	if _, err := Run("git", "fetch", "--all"); err != nil {
		fmt.Println(err)
		return
	}

	//Are we already on the branch we want to work on
	if strings.HasPrefix(ctx.CurrentBranch, keyOrID) {
		fmt.Fprintf(os.Stdout, "You are already on branch %s \n", ui.Format.Branch(ctx.CurrentBranch))
		return
	}
	//Does another workspace already exist with the key or ID
	for _, branch := range ctx.Workspaces {
		if strings.HasPrefix(branch, keyOrID) {

			if _, err := Run("git", "checkout", branch); err != nil {
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

	_, gitErr := Run("git", "checkout", "-b", name, conf.Remote+"/"+conf.IntegrationBranch)
	if gitErr != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Created workspace:  %s \n", ui.Format.Branch(name))
}

//UpdateGithubIssue ...
//Updating an issue on github and returns the branch name
func UpdateGithubIssue(issue string, conf *setting.ProjectSetting) (string, error) {

	ctx := context.WorkspaceContext

	issueOb, err := plugins.GetIssueGitHub(conf.IssueApi, ctx.Organisation, ctx.Repository, issue, ctx.Token)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err := plugins.SetAssigneeGitHub(conf.IssueApi, ctx.Organisation, ctx.Repository, ctx.Token, issue, ctx.User); err != nil {
		fmt.Println(err)
	}

	if _, err := plugins.SetLabelGitHub(conf.IssueApi, ctx.Organisation, ctx.Repository, ctx.Token, plugins.PhlowLabels["Status - in progress"].Title, issue); err != nil {
		fmt.Println(err)
	}

	is := strconv.Itoa(issueOb.Number)
	fmt.Fprintf(os.Stdout, "\n-------- Issue %s updated --------  \n", ui.Format.Issue(is))
	fmt.Fprintf(os.Stdout, "Label    => %s \n", ui.Format.Label.G4Move(plugins.PhlowLabels["Status - in progress"].Title))
	fmt.Fprintf(os.Stdout, "Assignee => %s \n", ui.Format.Assignee(ctx.User))
	fmt.Println("----------------------------------")

	return plugins.BranchNameFromIssue(is, issueOb.Title), nil

}
