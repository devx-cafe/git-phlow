package subprocess

import (
	"os/exec"
	"bytes"
	"errors"
	"syscall"
	"log"
)

const (
	EmptyReturnString string = ""
)

type ExecError struct {
	err    error
	StdErr   string
	ExitCode int
}

func (e ExecError) Error() string {
	return e.StdErr
}


//SimpleExec
//Executes a cmd on your operating system
func SimpleExec(name string, args ...string) (string, error) {

	if err := IsInPath(name); err != nil {
		return EmptyReturnString, err
	}

	cmd := exec.Command(name, args...)

	var outBuffer, errBuffer bytes.Buffer

	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %d", err)
		return EmptyReturnString, err
	}

	if err := cmd.Wait(); err != nil {

		if exitErr, ok := err.(*exec.ExitError); ok {
			//Program exited with exit code != 0

			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				//Trying to obtain exit error from failed command

				log.Printf("Exit Status: %d", status.ExitStatus())
				return EmptyReturnString, ExecError{err, errBuffer.String(), status.ExitStatus()}
			}
		}
		return EmptyReturnString, err
	}

	return outBuffer.String(), nil
}

//IsInPath
//Checks if an app has been added to $PATH
func IsInPath(application string) (error) {
	_, err := exec.LookPath(application)

	if err != nil {
		return errors.New(application + " is not in $PATH")

	}
	return nil
}