package executor

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/praqma/git-phlow/options"
	"io"
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

//Commander ...
//interface for os executions
type Commander interface {
	Run() error
}

//ExecuteCommander ...
//Run a function with control over stdout and stdin
func ExecuteCommander(c Commander) error {
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}


//Runner ...
//Runner type for git executions
type Runner func(command string, argv ...string) (string, error)

//Run ...
//implemented runner
func Run(command string, argv ...string) (string, error) {
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

//ExecuteCommand ...
//Executes a single command from strings
func ExecuteCommand(command string, argv ...string) (string, error) {
	exe := exec.Command(command, argv...)

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

//ExecPipeCommand ...
//Executes a series of commands
func ExecPipeCommand(out *bytes.Buffer, execStack ...*exec.Cmd) (err error) {

	var errBuf bytes.Buffer

	pipes := make([]*io.PipeWriter, len(execStack)-1)
	execStack[len(execStack)-1].Stdout = out

	for i := 0; i < len(execStack)-1; i++ {
		r, w := io.Pipe()
		execStack[i].Stdout = w
		execStack[i].Stderr = &errBuf
		execStack[i+1].Stdin = r
		pipes[i] = w

	}

	for i := 0; i < len(execStack); i++ {
		if options.GlobalFlagVerbose {
			verboseOutput(execStack[i].Args...)
		}

		if err = execStack[i].Start(); err != nil {
			return
		}
	}

	if err = closer(pipes, execStack); err != nil {
		return
	}

	return nil
}

//closer ...
//helper function to ExecutePipeCommand
func closer(pipes []*io.PipeWriter, execStack []*exec.Cmd) (err error) {
	//Return if the command-stack is empty
	if len(execStack) <= 0 {
		return nil
	}

	defer func() {
		//Close the pipe if more exists
		if len(pipes) > 0 {
			if err = pipes[0].Close(); err != nil {
				return
			}
			pipes = pipes[1:]
		}
		if err = closer(pipes, execStack[1:]); err != nil {
			return
		}
	}()
	//Wait for the command to return, defer is the
	//last function to be executed
	return execStack[0].Wait()
}
