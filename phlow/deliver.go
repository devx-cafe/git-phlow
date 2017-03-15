package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"os"
	"strings"

	"github.com/praqma/git-phlow/ui"
)

//Deliver ...
func Deliver(defaultBranch string) {

	branchInfo, _ := githandler.Branch()
	githandler.Fetch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == defaultBranch) {
		fmt.Printf("Could not deliver: %s", branchInfo.Current)
		return
	}

	ui.PhlowSpinner.Start("pushing")
	_, err := githandler.PushRename(branchInfo.Current, defaultBranch)
	if err != nil {
		fmt.Println(err)
		return
	}
	ui.PhlowSpinner.Stop()
	fmt.Println("Changes pushed")

	if err := githandler.BranchRename(branchInfo.Current); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Branch %s  is now delivered \n", ui.BranchFormat(branchInfo.Current))
}

//LocalDeliver ...
func LocalDeliver(defaultBranch string) {

	branchInfo, _ := githandler.Branch()

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") || (branchInfo.Current == defaultBranch) {
		fmt.Printf("You cannot deliver: %s \n", branchInfo.Current)
		return
	}

	fmt.Fprintf(os.Stdout, "Checking out default branch %s \n", ui.BranchFormat(defaultBranch))
	//Checkout default branch: master
	if err := githandler.CheckOut(defaultBranch); err != nil {
		fmt.Println(err)
		return
	}
	//Pull rebase latest changes
	fmt.Fprintln(os.Stdout, "Trying to pull latest changes")
	output, err := githandler.Pull()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	fmt.Fprintf(os.Stdout, "Merging changes from branch %s into branch %s \n", ui.BranchFormat(branchInfo.Current), ui.BranchFormat(defaultBranch))
	//Merge feature branch into default
	if err = githandler.Merge(branchInfo.Current); err != nil {
		fmt.Println(err)
	}
	//Rename default branch to delivered
	githandler.BranchRename(branchInfo.Current)

	//Push changes to github
	fmt.Fprintf(os.Stdout, "Pushing changes to remote %s \n", ui.BranchFormat(defaultBranch))
	ui.PhlowSpinner.Start("Pushing changes")
	output, err = githandler.Push()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ui.PhlowSpinner.Stop()
	fmt.Printf("Changes from %s delivered to %s \n", ui.BranchFormat(branchInfo.Current), ui.BranchFormat(defaultBranch))

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
