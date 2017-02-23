package phlow

import (
	"fmt"
	"strings"

	"os"

	"github.com/praqma/git-phlow/githandler"
)

//WrapUp ...
func WrapUp() {

	fmt.Fprintln(os.Stdout, "Addind files to index")

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
	fmt.Fprintln(os.Stdout, commitMessage)
}
