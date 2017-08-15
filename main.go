package main

import (
	"github.com/praqma/git-phlow/setting"
	"fmt"
)

func main() {
	//cmd.Execute()

	proj := setting.NewProjectStg("default")

	err := setting.ValidateLoadedSetting(proj)
	fmt.Println(err)

}
