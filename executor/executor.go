package executor

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/praqma/git-phlow/options"
)

//VerboseOutput ...
//prints the commands being run by the program
func VerboseOutput(application string, argv ...string) {
	fmt.Print("Exec: ", application)
	for _, arg := range argv {
		fmt.Print(" " + arg)
	}
	fmt.Println()
}

//ExecuteCommand ...
func ExecuteCommand(command string, argv ...string) (string, error) {
	if options.GlobalFlagVerbose {
		VerboseOutput(command, argv...)
	}

	exe := exec.Command(command, argv...)

	var stdOutBuffer, stdErrBuffer bytes.Buffer

	exe.Stderr = &stdErrBuffer
	exe.Stdout = &stdOutBuffer

	if err := exe.Start(); err != nil {
		return "", errors.New(stdErrBuffer.String())
	}

	if err := exe.Wait(); err != nil {
		return "", errors.New(stdErrBuffer.String())
	}

	return stdOutBuffer.String(), nil
}
