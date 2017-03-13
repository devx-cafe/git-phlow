package phlow

import (
	"github.com/praqma/git-phlow/githandler"
	"sort"
	"fmt"
	"github.com/praqma/git-phlow/options"
)

func UpNext(remote string) {

	branches := githandler.BranchReady(remote)

	if len(branches) != 0 {
		if options.GlobalFlagHumanReadable {
			fmt.Println("'ready/' branches found on remote")
		}

		next := GetNextBranch(branches)
		fmt.Print(next)
		return
	}

	if options.GlobalFlagHumanReadable {
		fmt.Println("No 'ready/' branches found on remote")
	}

}

func GetNextBranch(branches []string) string {
	m := make(map[int]string)
	var time int
	var err error

	//Create map with time and branch name
	for _, br := range branches {
		if time, err = githandler.BranchTime(br); err == nil {
			m[time] = br
			
			if options.GlobalFlagHumanReadable {
				fmt.Printf("%s : %d", br, time)
			}
		}
	}

	//Order the keys in a separate list
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	if len(keys) > 0 {
		return m[keys[0]]
	}
	return ""
}
