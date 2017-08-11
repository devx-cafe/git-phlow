package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/ui"
	"github.com/praqma/git-phlow/setting"
)

//MakeAliasCaller ...
func MakeAliasCaller() {
	MakeAlias(setting.GitConfig{})
}

//MakeAlias ...
func MakeAlias(conf setting.Configurator) {
	aliases := make(map[string]string)
	aliases["alias.wrapup"] = "phlow wrapup"
	aliases["alias.workon"] = "phlow workon"
	aliases["alias.deliver"] = "phlow deliver"
	aliases["alias.cleanup"] = "phlow cleanup"
	aliases["alias.web"] = "phlow web"
	aliases["alias.issues"] = "phlow issues"

	for group, value := range aliases {

		str := conf.Get(group)
		if str == "" {
			fmt.Printf("Creating alias %s \n", ui.Format.Alias(group))
			conf.Set(group, value)
		} else {
			fmt.Printf("Alias %s already exists \n", ui.Format.Alias(group))
		}
	}
}
