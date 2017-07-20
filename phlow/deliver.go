package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"

	"github.com/praqma/git-phlow/ui"
)

//Deliver ...
func Deliver(defaultBranch string) {

	ui.PhlowSpinner.Start("delivering")
	defer ui.PhlowSpinner.Stop()

	branchInfo, _ := githandler.Branch()
	githandler.Fetch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == defaultBranch) {
		fmt.Printf("Could not deliver: %s", branchInfo.Current)
		return
	}

	_, err := githandler.PushRename(branchInfo.Current, defaultBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := githandler.BranchRename(branchInfo.Current); err != nil {
		//Conflicting branch name
		fmt.Printf("branch delivered/%s already exists in you local workspace \n", branchInfo.Current)
		return
	}

	githandler.CheckOut(defaultBranch)
	if err != nil {
		fmt.Printf("There are still changes in your workspace %s \n", defaultBranch)
		fmt.Println("try: 'git status' to see the changes")
		return
	}

	fmt.Printf("Delivered branch %s \n", ui.Format.Branch(branchInfo.Current))
}

//LocalDeliver ...
func LocalDeliver(defaultBranch string) {

	branchInfo, _ := githandler.Branch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == defaultBranch) {
		fmt.Printf("You cannot deliver: %s \n", branchInfo.Current)
		return
	}

	//Checkout default branch: master
	if err := githandler.CheckOut(defaultBranch); err != nil {
		fmt.Println(err)
		return
	}

	//Pull rebase latest changes
	ui.PhlowSpinner.Start("delivering")
	defer ui.PhlowSpinner.Stop()

	_, err := githandler.Pull()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Merge feature branch into default
	if err = githandler.Merge(branchInfo.Current); err != nil {
		fmt.Println(err)
		return
	}
	//Rename default branch to delivered
	githandler.BranchRename(branchInfo.Current)

	//Push changes to GitHub
	_, err = githandler.Push()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Delivered changes from %s to %s \n", ui.Format.Branch(branchInfo.Current), ui.Format.Branch(defaultBranch))
}

//TestDeliver ...
//Run tests and returns
func TestDeliver(args []string) error {

	cmd, argv := convertCommand(args)
	output, err := executor.ExecuteCommand(cmd, argv...)

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
