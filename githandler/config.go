package githandler

import (
	"fmt"
	"github.com/praqma/git-phlow/executor"
	"strings"
)

//ConfigBranchRemote ...
func ConfigBranchRemote(branch string) string {
	configArg := fmt.Sprintf("branch.%s.remote", branch)
	output, _ := executor.ExecuteCommand("git", "config", configArg)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigGet ...
func ConfigGet(key, group string) string {
	pair := fmt.Sprintf("%s.%s", group, key)
	output, _ := executor.ExecuteCommand("git", "config", "--global", "--get", pair)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigSet ...
func ConfigSet(key, value, group string) error {
	pair := fmt.Sprintf("%s.%s", group, key)
	_, err := executor.ExecuteCommand("git", "config", "--global", pair, value)
	return err
}
