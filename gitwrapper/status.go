package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

//Statuses ...
//Interface for git status commands
type Statuses interface {
	Status() error
}

//Status ...
//Struct for git status base commands
type Status struct {
	baseCmd    string
	baseStatus string
}

//NewStatus ...
//Creates a new git status
func NewStatus(baseCMD string) *Status {
	return &Status{baseCmd: baseCMD, baseStatus: "status"}

}

//Status ...
//Runs git status, and returns an error if executed
//outside of a git repository
func (s *Status) Status() error {
	_, err := subprocess.SimpleExec(s.baseCmd, s.baseStatus)

	if err != nil {
		return err
	}

	return nil
}
