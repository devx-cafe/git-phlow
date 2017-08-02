package platform

import (
	"strings"
	"github.com/praqma/git-phlow/executor"
	"runtime"
)

type ProjectSetting int

const (
	_ ProjectSetting = iota

	BITBUCKET
	GITHUB
	UNRECOGNIZEDPLATFORM
	WINDOWS
	UNIX
)

//NewDefaultPlatform ...
func NewDefaultPlatform() *Platform {
	return &Platform{executor.Run}
}

//Platform ...
type Platform struct {
	Run executor.Runner
}

//System ...
func (p Platform) System() ProjectSetting {
	if runtime.GOOS == "windows" {
		return WINDOWS
	}
	return UNIX
}

//Service ...
func (p Platform) Service() ProjectSetting {
	url, err := p.Run("git", "ls-remote", "--get-url", "origin")
	if err != nil {
		panic(err)
	}
	//url format: git@github.com:Praqma/git-phlow.git
	//url format: git@bitbucket.com:Praqma/git-phlow.git
	if strings.Contains(url, "bitbucket") {
		return BITBUCKET
	}
	if strings.Contains(url, "github") {
		return GITHUB
	}
	return UNRECOGNIZEDPLATFORM
}
