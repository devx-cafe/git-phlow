package githandler

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	. "github.com/praqma/git-phlow/executor"
)

//CheckOut ...
func CheckOut(branch string, new bool) error {
	if new {
		_, err := RunCommand("git", "checkout", "-b", branch, "origin/master")
		return err
	}
	_, err := RunCommand("git", "checkout", branch)
	return err
}

//Status ...
func Status() error {
	_, err := RunCommand("git", "status")
	return err
}

//Add ...
func Add() error {
	_, err := RunCommand("git", "add", "--all")
	return err
}

//Commit ...
func Commit(message string) (string, error) {
	return RunCommand("git", "add", "--all")
}

//Fetch ...
func Fetch() error {
	_, err := RunCommand("git", "fetch", "--all")
	return err
}

//Pull ...
func Pull() (string, error) {
	return RunCommand("git", "pull", "--rebase")
}

//Push ...
func Push(branch string, rename bool) (string, error) {
	str := fmt.Sprintf("%s:ready/%s", branch, branch)
	if rename {
		return RunCommand("git", "push", "origin", str)
	}
	return RunCommand("git", "push")
}

//Merge ...
func Merge(branch string) error {
	_, err := RunCommand("git", "merge", branch)
	return err
}

//RemoteInfo ...
type RemoteInfo struct {
	Organisation string
	Repository   string
}

//Remote ...
func Remote() (*RemoteInfo, error) {
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)
	output, err := RunCommand("git", "remote", "-v")
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(output)
	return &RemoteInfo{match[1], match[2]}, nil
}

//Config ...
func Config(key, value string, get bool) (string, error) {
	pair := fmt.Sprintf("%s.%s", "do", key)

	if get {
		output, err := RunCommand("git", "config", "--global", "--get", pair)
		return output, err
	}
	_, err := RunCommand("git", "config", "--global", pair, value)
	return "", err
}

//BranchInfo ...
type BranchInfo struct {
	Current string
	List    []string
}

//Branch ...
func Branch(key string) (*BranchInfo, error) {
	var err error
	info := BranchInfo{}

	switch key {
	case "current":
		info.Current, err = RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
		if err == nil {
			info.Current = strings.TrimSpace(info.Current)
		}
		return &info, err
	case "list":
		if output, err := RunCommand("git", "branch"); err == nil {

			for _, branch := range strings.Split(output, "\n") {
				if branch != "" {
					info.List = append(info.List, strings.TrimSpace(branch))
				}
			}

			return &info, err
		}
		return nil, err
	default:
		return nil, errors.New("function 'Branch' input is wrong")
	}
}

//BranchRename ...
func BranchRename(name string) error {
	_, err := RunCommand("git", "branch", "-m", name, "delivered/"+name)
	return err
}
