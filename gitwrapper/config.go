package gitwrapper

import (
	"fmt"
)

//Configurator ...
//Interface for persistent configuration
type Configurator interface {
	SetConfig(string, string, func(name string, args ...string) (string, error)) error
	GetConfig(string, func(name string, args ...string) (string, error)) string
}

type config struct {
}

//NewConfigurator ...
//Generates a new Configurator
func NewConfigurator() Configurator {
	return &config{}
}

//GetConfig ...
//Gets configuration through executor
func (c *config) GetConfig(key string, executor func(string, ...string) (string, error)) string {
	var keys = fmt.Sprintf("%s.%s", "phlow", key)
	var baseCMD, config, scope, get = "git", "config", "--global", "--get"

	value, _ := executor(baseCMD, config, scope, get, keys)
	return value
}

//SetConfig ...
//Sets configuration through executor
func (c *config) SetConfig(key, value string, executor func(string, ...string) (string, error)) error {

	var keys = fmt.Sprintf("%s.%s", "phlow", key)
	var baseCMD, config, scope = "git", "config", "--global"

	_, err := executor(baseCMD, config, scope, keys, value)

	if err != nil {
		return fmt.Errorf("Error in keys, was: %s %s %s %s %s", baseCMD, config, scope, keys, value)
	}

	return nil
}
