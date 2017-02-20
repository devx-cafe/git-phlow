package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/githandler"
	"strings"
)

//WrapUp ...
func WrapUp() {
	//Before check - status

	if err := githandler.Add(); err != nil {
		fmt.Println("project files could not be added: " + err.Error())
		return
	}

	info, _ := githandler.Branch("current")


	commitMessage := strings.Replace(info.Current, "-", " ", -1)
	if output, err := githandler.Commit("close #" + commitMessage); err == nil {
		fmt.Println(output)
		return
	}
}
