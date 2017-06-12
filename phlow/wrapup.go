package phlow

import (
	"fmt"
	"strings"

	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
)

//WrapUp ...
func WrapUp() {

	//Add all files to index
	if err := githandler.Add(); err != nil {
		fmt.Println("Project files could not be added: " + err.Error())
		return
	}

	//Retrieve branch info - current branch
	info, _ := githandler.Branch()
	var cmsg string

	if options.GlobalFlagForceMessage != "" {
		cmsg = "close #" + strings.Split(info.Current, "-")[0] + " " + options.GlobalFlagForceMessage
	} else {
		cmsg = "close #" + strings.Replace(info.Current, "-", " ", -1)
	}

	if _, err := githandler.Commit(cmsg); err != nil {
		fmt.Println("Nothing to commit!")
		return
	}
	fmt.Fprintln(os.Stdout, cmsg)
}
