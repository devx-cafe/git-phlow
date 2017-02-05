package subprocess

import (
	"os/exec"
	"bytes"
	"errors"
	"fmt"
	"syscall"
	"log"
)

const (
	EmptyReturnString string = ""
)

type execError struct {
	Error    error
	StdErr   string
	ExitCode int
}

//SimpleExec
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
				log.Printf("Exit Status: %d", status.ExitStatus())
				return execError{err, errBuffer.String(), status.ExitStatus()}
			}
		}
		return EmptyReturnString, err
	}

	return outBuffer.String(), nil
}

func IsInPath(application string) (error) {
	_, err := exec.LookPath(application)

	if err != nil {
		return errors.New(application + " is not in $PATH")

	}
	return nil
}