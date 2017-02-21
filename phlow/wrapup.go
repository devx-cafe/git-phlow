package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/githandler"
	"strings"
	"github.com/praqma/git-phlow/printers"
)

//WrapUp ...
func WrapUp(verbose bool) {

	printers.PrintVerbose("Addind files to index", verbose)

	//Add all files to index
	if err := githandler.Add(); err != nil {
		fmt.Println("project files could not be added: " + err.Error())
		return
	}

	//Retrieve branch info - current branch
	info, _ := githandler.Branch()
	commitMessage := "close #" + strings.Replace(info.Current, "-", " ", -1)

	if _, err := githandler.Commit(commitMessage); err != nil {
		fmt.Println(err)
		return
	}
	printers.PrintVerbose(commitMessage, verbose)
}
