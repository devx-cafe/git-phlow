package phlow

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
)

//MkAlias ...
func MkAlias(verbose bool) {
	group := "alias"
	aliases := make(map[string]string)
	aliases["wrapup"] = "phlow wrapup"
	aliases["workon"] = "phlow workon"
	aliases["deliver"] = "phlow deliver"
	aliases["purge"] = "phlow purge"

	for key, value := range aliases {
		str, _ := githandler.Config(key, "", group, true)
		if str == "" {
			fmt.Printf("Creating alias '%s' \n", key)
			githandler.Config(key, value, group, false)
		} else {
			fmt.Printf("Alias '%s' already exists \n", key)
		}
	}

	fmt.Fprintln(os.Stdout, "Alias's updated")
}
