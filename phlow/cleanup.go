package phlow

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/code-cafe/git-phlow/context"
	"github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/flags"
	"github.com/code-cafe/git-phlow/setting"
	"github.com/fatih/color"
	"log"
)

//CleanCaller ...
func CleanCaller(ini string) {
	conf := setting.NewProjectStg(ini)
	Clean(conf)
}

//Clean ...
//deletes all the delivered branches
func Clean(conf *setting.ProjectSetting) {
	var wsc = context.WorkspaceContext

	if flags.CleanupDelivered {
		deleteBranches(wsc.DeliveredBranches)
		return
	}

	var selectedAnswers []string
	prompt := &survey.MultiSelect{
		Message: "Select branches delete from repository",
		Options: wsc.Branches,
	}

	if err := survey.AskOne(prompt, &selectedAnswers); err != nil {
		fmt.Println("aborting ...")
	}

	deleteBranches(selectedAnswers)

	if flags.CleanUpTidy {
		out, err := executor.Run("git", "fetch", "--prune")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	}

}

func deleteBranches(branches []string) {
	if len(branches) == 0 {
		fmt.Println("No branches to delete")
		return
	}

	for _, branch := range branches {
		force := "-d"

		if flags.CleanupForce {
			force = "-D"
		}

		if flags.CleanUpDryRun {

			fmt.Printf("Will delete branch %s \n", color.New(color.Bold).Sprint(branch))
		} else {

			out, err := executor.Run("git", "branch", force, branch)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Deleted branch %s", out)
		}
	}
}
