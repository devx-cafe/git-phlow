package executor

import (
	"bytes"
	"errors"
	"os/exec"
)

//RunCommand ...
func RunCommand(command string, argv ...string) (string, error) {
	cmd := exec.Command(command, argv...)

	var stdOutBuffer, stdErrBuffer bytes.Buffer

	cmd.Stderr = &stdErrBuffer
	cmd.Stdout = &stdOutBuffer

	if err := cmd.Start(); err != nil {
		return "", errors.New(stdErrBuffer.String())
	}

	if err := cmd.Wait(); err != nil {
		return "", errors.New(stdErrBuffer.String())
	}

	return stdOutBuffer.String(), nil
}
