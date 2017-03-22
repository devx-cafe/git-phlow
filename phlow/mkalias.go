package phlow

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/ui"
)

//MkAlias ...
func MkAlias() {
	group := "alias"
	aliases := make(map[string]string)
	aliases["wrapup"] = "phlow wrapup"
	aliases["workon"] = "phlow workon"
	aliases["deliver"] = "phlow deliver"
	aliases["cleanup"] = "phlow cleanup"
	aliases["agent"] = "phlow agent"
	aliases["web"] = "phlow web"
	aliases["issues"] = "phlow issues"
	aliases["auth"] = "phlow auth"

	for key, value := range aliases {
		str := githandler.ConfigGet(key, group)
		if str == "" {
			fmt.Printf("Creating alias %s \n", ui.Format(key).Alias)
			githandler.ConfigSet(key, value, group)
		} else {
			fmt.Printf("Alias %s already exists \n", ui.Format(key).Alias)
		}
	}

	fmt.Fprintln(os.Stdout, "Alias's updated")
}
