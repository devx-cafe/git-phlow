package platform

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


func DefaultConfiguration() Configurator {
	return KeyConfiguration{Run: executor.Run}
}

//Configurator ...
//Interface for system configurations
type Configurator interface {
	Set(group string, value string)
	Get(group string) string
	Unset(group string)
}

//KeyConfiguration ...
type KeyConfiguration struct {
	Run executor.Runner
}

//Get ...
func (pc KeyConfiguration) Get(group string) string {
	output, err := pc.Run("git", "config", "--global", "--get", group)
	if err != nil {
		panic(err)
	}
	return strings.Replace(output, "\n", "", -1)
}

//Set ...
func (pc KeyConfiguration) Set(group string, value string) {
	_, err := pc.Run("git", "config", "--global", group, value)
	if err != nil {
		panic(err)
	}
}

//Unset ...
func (pc KeyConfiguration) Unset(group string) {
	_, err := pc.Run("git", "config", "--global", "--unset", group)
	if err != nil {
		panic(err)
	}
}
