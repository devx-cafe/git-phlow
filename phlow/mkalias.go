package phlow

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/printers"
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
			printers.PrintVerbose(fmt.Sprintf("Creating alias '%s'", key), verbose)
			githandler.Config(key, value, group, false)
		} else {
			printers.PrintVerbose(fmt.Sprintf("Alias '%s' already exists", key), verbose)
		}
	}

	fmt.Fprintln(os.Stdout, "Alias's updated")
}
