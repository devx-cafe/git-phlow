package subprocess

import (
	"bytes"
	"errors"
	"os/exec"
	"syscall"
)

const (
	emptyReturnString string = ""
)

//ExecError ...
type ExecError struct {
	error
	StdErr   string
	ExitCode int
}

func (e ExecError) Error() string {
	return e.StdErr
}

//SimpleExec ...
//Executes a cmd on your operating system
func SimpleExec(name string, args ...string) (string, error) {

	cmd := exec.Command(name, args...)

	var stdOutBuffer, stdErrBuffer bytes.Buffer

	cmd.Stderr = &stdErrBuffer
	cmd.Stdout = &stdOutBuffer

	if err := cmd.Start(); err != nil {

		//Return Error with stderr, error - and exit status 1
		return emptyReturnString, ExecError{err, stdErrBuffer.String(), 1}
	}

	if err := cmd.Wait(); err != nil {

		if exitErr, ok := err.(*exec.ExitError); ok {
			//Program exited with exit code != 0

			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				//Trying to obtain exit error from failed command

				return emptyReturnString, ExecError{err, stdErrBuffer.String(), status.ExitStatus()}
			}
		}

		//Return Error with stderr, error - and exit status 1
		return emptyReturnString, ExecError{err, stdErrBuffer.String(), 1}
	}

	//If no errors are returned, return stdout
	return stdOutBuffer.String(), nil
}

//IsInPath ...
//Checks if an app has been added to $PATH
func IsInPath(application string) error {
	_, err := exec.LookPath(application)

	if err != nil {
		return errors.New(application + " is not in $PATH")

	}
	return nil
}
