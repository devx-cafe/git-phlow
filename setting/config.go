package setting

import (
	"strings"
	"github.com/praqma/git-phlow/executor"
)

const (
	//PhlowUser ...
	PhlowUser = "phlow.user"

	//PhlowToken ...
	PhlowToken = "phlow.token"
)

//DefaultGitConfig ...
func DefaultGitConfig() Configurator {
	return GitConfig{Run: executor.Run}
}

//Configurator ...
//Interface for system configurations
type Configurator interface {
	Set(group string, value string)
	Get(group string) string
	Unset(group string)
}

//GitConfig ...
type GitConfig struct {
	Run executor.Runner
}

//Get ...
func (pc GitConfig) Get(group string) string {
	output, err := pc.Run("git", "config", "--global", "--get", group)
	if err != nil {
		panic(err)
	}
	return strings.Replace(output, "\n", "", -1)
}

//Set ...
func (pc GitConfig) Set(group string, value string) {
	_, err := pc.Run("git", "config", "--global", group, value)
	if err != nil {
		panic(err)
	}
}

//Unset ...
func (pc GitConfig) Unset(group string) {
	_, err := pc.Run("git", "config", "--global", "--unset", group)
	if err != nil {
		panic(err)
	}
}
