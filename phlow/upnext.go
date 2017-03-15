package phlow

import (
	"fmt"
	"os"
	"sort"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
)

//UpNext ...
//Returns the next branch ready for integration based on time of creation
//Oldest branches gets integrated first.
func UpNext(remote string) {

	branches := githandler.BranchReady(remote)

	if len(branches) != 0 {
		if options.GlobalFlagHumanReadable {
			fmt.Println("'ready/' branches found on remote")
		}

		next := getNextBranch(branches)
		fmt.Fprint(os.Stdout, next)
		return
	}

	if options.GlobalFlagHumanReadable {
		fmt.Println("No 'ready/' branches found on remote")
	}

}

//getNextBranch
//Sort branches and returns the oldest ready branch
func getNextBranch(branches []string) string {
	m := make(map[int]string)
	var time int
	var err error

	//Create map with time and branch name
	for _, br := range branches {
		if time, err = githandler.BranchTime(br); err == nil {
			m[time] = br

			if options.GlobalFlagHumanReadable {
				fmt.Printf("%s : %d \n", br, time)
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
