package phlow

import (
	"strings"
	"github.com/praqma/git-phlow/githandler"
	"fmt"
	"os"
	"bytes"
)

//Sync ...
func Sync() {

	remote := githandler.RemoteBranch()
	status, err := githandler.StatusPorcelain()
	if err != nil {
		fmt.Println(err)
		return
	}

	behind, ahead := isBehindOrAhead(status)

	if behind {
		fmt.Fprintf(os.Stdout, "Your branch is behind %s\n", remote)

		if ahead {
			fmt.Fprintf(os.Stdout, "Your branch is behind %s \n", remote)

			var buf bytes.Buffer
			ms := githandler.FormatPatch(&buf, remote)
			conflict := strings.Contains(ms.Error(), "Patch failed")

			if conflict {
				fmt.Println("You have a merge conflict")
				fmt.Println(ms.Error())

			} else {
				fmt.Println("No merge conflicts indentified")
			}
		}

	} else {
		//up to date
	}
}

//isBehindOrAhead
func isBehindOrAhead(str string) (behind, ahead bool) {
	ahead = strings.Contains(str, "ahead")
	behind = strings.Contains(str, "behind")
	return
}
