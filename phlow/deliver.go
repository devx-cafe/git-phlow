package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"

	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/setting"
)

//Deliver ...
func Deliver(defaultBranch string) {

	git := githandler.Git{Run: executor.RunGit}
	conf := setting.NewProjectStg("default")

	ui.PhlowSpinner.Start("delivering")
	defer ui.PhlowSpinner.Stop()

	out, err := git.Branch("-a")
	branchInfo := githandler.AsList(out)

	_, err = git.Pull("--rebase")
	if err != nil {
		fmt.Println(err)
	}

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == conf.IntegrationBranch) {
		fmt.Printf("Could not deliver: %s", branchInfo.Current)
		return
	}

	_, err = githandler.PushRename(branchInfo.Current, conf.IntegrationBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := githandler.BranchRename(branchInfo.Current); err != nil {
		//Conflicting branch name
		fmt.Printf("branch delivered/%s already exists in you local workspace \n", branchInfo.Current)
		return
	}

	out, err = git.CheckOut(conf.IntegrationBranch)
	if err != nil {
		fmt.Printf("There are still changes in your workspace %s \n", conf.IntegrationBranch)
		fmt.Println("try: 'git status' to see the changes")
		return
	}

	ui.PhlowSpinner.Stop()
	fmt.Printf("Delivered branch %s \n", ui.Format.Branch(branchInfo.Current))
}

//LocalDeliver ...
func LocalDeliver(defaultBranch string) {
	git := githandler.Git{Run: executor.RunGit}
	conf := setting.NewProjectStg("default")

	out, err := git.Branch("-a")
	branchInfo := githandler.AsList(out)

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == conf.IntegrationBranch) {
		fmt.Printf("You cannot deliver: %s \n", branchInfo.Current)
		return
	}

	//Checkout default branch: master
	_, err = git.CheckOut(conf.IntegrationBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Pull rebase latest changes
	ui.PhlowSpinner.Start("delivering")
	defer ui.PhlowSpinner.Stop()

	_, err = git.Pull("--rebase")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Merge feature branch into default
	_, err = git.Merge(branchInfo.Current)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Rename default branch to delivered
	_, err = git.Branch("-m", branchInfo.Current, "delivered/"+branchInfo.Current)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Push changes to GitHub
	_, err = git.Push()
	if err != nil {
		fmt.Println(err)
		return
	}
	ui.PhlowSpinner.Stop()
	fmt.Printf("Delivered changes from %s to %s \n", ui.Format.Branch(branchInfo.Current), ui.Format.Branch(defaultBranch))
}

//TestDeliver ...
//Run tests and returns
func TestDeliver(args []string) error {

	cmd, argv := convertCommand(args)
	output, err := executor.RunCommand(cmd, argv...)

	if err != nil {
		return err
	}

	if options.GlobalFlagShowTestOutput {
		fmt.Println(output)
	}

	return nil
}

//ConvertCommand ...
//Formats the command to ExecutorCommand
func convertCommand(args []string) (string, []string) {

	//Command with extra arguments
	if len(args) > 1 {
		return args[0], args[1:]
	}
	return args[0], []string{}
}
