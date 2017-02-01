package subprocess

import (
	"os"
	"fmt"
	"os/exec"
	"syscall"
)

func SimpleExec(name string, args ...string) {

	var waitStatus syscall.WaitStatus
	cmd := exec.Command(name, args...)

	//Up for change
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {

		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Sprintf("%d", waitStatus.ExitStatus())
		}
	} else {

		//Command executed successfully
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Sprintf("%d", waitStatus.ExitStatus())
	}

}

func ValidateGitInstallation() (string, error) {
	path, err := exec.LookPath("git")

	if err != nil {
		return "", err

	}
	return path, nil
}