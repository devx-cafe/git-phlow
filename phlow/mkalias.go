package phlow

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/githandler"
)

//MkAlias ...
func MkAlias() {
	group := "alias"
	aliases := make(map[string]string)
	aliases["wrapup"] = "phlow wrapup"
	aliases["workon"] = "phlow workon"
	aliases["deliver"] = "phlow deliver"
	aliases["clean"] = "phlow clean"

	for key, value := range aliases {
		str := githandler.ConfigGet(key, group)
		if str == "" {
			fmt.Printf("Creating alias '%s' \n", key)
			githandler.ConfigSet(key, value, group)
		} else {
			fmt.Printf("Alias '%s' already exists \n", key)
		}
	}

	fmt.Fprintln(os.Stdout, "Alias's updated")
}
