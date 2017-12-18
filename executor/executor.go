package executor

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/praqma/git-phlow/options"
)

//verboseOutput ...
//prints the commands being run by the program
func verboseOutput(argv ...string) {
	fmt.Print("Exec:")
	for _, arg := range argv {
		fmt.Print(" " + arg)
	}
	fmt.Println()
}

//Runner ...
//Runner type for git executions
type Runner func(command string, argv ...string) (string, error)

//RunCommand ...
//implemented runner
func RunCommand(command string, argv ...string) (string, error) {
	var stdOutBuffer, stdErrBuffer bytes.Buffer
	exe := exec.Command(command, argv...)

	if options.GlobalFlagVerbose {
		verboseOutput(exe.Args...)
	}

	exe.Stderr = &stdErrBuffer
	exe.Stdout = &stdOutBuffer

	err := exe.Run()
	if err != nil {
		if out := stdOutBuffer.String(); stdErrBuffer.String() == "" {
			return "", errors.New(out)
		}
		return "", errors.New(stdErrBuffer.String())
	}
	return stdOutBuffer.String(), nil
}

//GitCommandRunner ...
type GitCommandRunner func(git string, sub string, argv ...string) (string, error)

//RunGit ...
//Executes a single command from strings
func RunGit(git string, sub string, argv ...string) (string, error) {

	argv = append([]string{sub}, argv...)
	exe := exec.Command(git, argv...)

	if options.GlobalFlagVerbose {
		verboseOutput(exe.Args...)
	}

	var stdOutBuffer, stdErrBuffer bytes.Buffer

	exe.Stderr = &stdErrBuffer
	exe.Stdout = &stdOutBuffer

	if err := exe.Start(); err != nil {
		return "", errors.New(stdErrBuffer.String())
	}

	if err := exe.Wait(); err != nil {
		if out := stdOutBuffer.String(); stdErrBuffer.String() == "" {
			return "", errors.New(out)
		}
		return "", errors.New(stdErrBuffer.String())
	}

	return stdOutBuffer.String(), nil
}
