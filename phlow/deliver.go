package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"

	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/setting"
	"os"
)

//DeliverCaller ...
//Top level deliver call, called from cmd
func DeliverCaller() {
	INIBlock := options.GlobalFlagTarget
	conf := setting.NewProjectStg(INIBlock)

	//If Run if local deliver
	if options.GlobalFlagLocal {
		LocalDeliver(conf)
		return
	}
	//Deliver with ready branch
	Deliver(conf)
}

//Deliver ...
//Push a ready branch to the remote repository
func Deliver(conf *setting.ProjectSetting) {
	git := githandler.Git{Run: executor.RunGit}

	fmt.Println("delivering...")

	out, err := git.Branch("-a")
	branchInfo := githandler.AsList(out)

	_, err = git.Pull("--rebase")
	if err != nil {
		fmt.Println(err)
	}

	//Is branch master or is branch delivered
	if strings.HasPrefix(branchInfo.Current, "delivered/") {
		fmt.Printf("It is not possible to deliver a branch already delivered: %s \n", branchInfo.Current)
		return
	}

	//Delivering master
	if branchInfo.Current == conf.IntegrationBranch {
		fmt.Printf("It is not possible to deliver the integration branch: %s \n", branchInfo.Current)
		return
	}

	//git push origin name:ready/name
	_, err = git.Push(conf.Remote, fmt.Sprintf("%s:%s/%s", branchInfo.Current, conf.DeliveryBranchPrefix, branchInfo.Current))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = git.Branch("-m", branchInfo.Current, "delivered/"+branchInfo.Current)
	if err != nil {
		fmt.Println("The branch have been pushed successfully to your remote, but there is a local name conflict")
		fmt.Printf("CONFLICT: your already have a branch named %s in your workspace \n", "delivered/"+branchInfo.Current)
		fmt.Printf("to mark it delivered run: git branch -m %s %s \n", branchInfo.Current, "delivered/"+branchInfo.Current)
		os.Exit(1)
		return
	}

	out, err = git.CheckOut(conf.IntegrationBranch)
	if err != nil {
		fmt.Printf("There are still changes in your workspace %s \n", conf.IntegrationBranch)
		fmt.Println("try: 'git status' to see the changes")
		return
	}

	fmt.Printf("Delivered branch %s \n", ui.Format.Branch(branchInfo.Current))
}

//LocalDeliver ...
//Delivers locally and pushes the changes to the remote
func LocalDeliver(conf *setting.ProjectSetting) {
	git := githandler.Git{Run: executor.RunGit}

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

	fmt.Printf("Delivered changes from %s to %s \n", ui.Format.Branch(branchInfo.Current), ui.Format.Branch(conf.IntegrationBranch))
}
