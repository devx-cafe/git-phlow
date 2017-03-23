package executor

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/praqma/git-phlow/options"
	"io"
	"os"
	"strconv"
)

//verboseOutput ...
//prints the commands being run by the program
func verboseOutput(application string, argv ...string) {
	fmt.Print("Exec: ", application)
	for _, arg := range argv {
		fmt.Print(" " + arg)
	}
	fmt.Println()
}

//ExecuteCommand ...
func ExecuteCommand(command string, argv ...string) (string, error) {
	if options.GlobalFlagVerbose {
		verboseOutput(command, argv...)
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

func ExecutePipe(outBuf *bytes.Buffer, callStack ...*exec.Cmd) (err error) {

	var errBuf bytes.Buffer

	pipeStack := make([]*io.PipeWriter, len(callStack)-1)

	i := 0
	for ; i < len(pipeStack)-1; i++ {
		inp, out := io.Pipe()
		callStack[i].Stdout = out
		callStack[i].Stderr = &errBuf
		callStack[i+1].Stdin = inp
		pipeStack[i] = out
	}

	callStack[i].Stdout = outBuf
	callStack[i].Stderr = &errBuf

	if err := call(callStack, pipeStack); err != nil {
		return err
	}
	return nil
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		fmt.Println("jkhjkhjkhjkhjkhkhjk")
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) >= 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				fmt.Println("JKSHDJKASHDJKSAHD" + strconv.Itoa(len(pipes)))

				pipes[0].Close()

				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}

func sd() {

	c1 := exec.Command("ls", "-lah")
	c2 := exec.Command("wc", "-l")

	r, w := io.Pipe()

	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer

	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	io.Copy(os.Stdout, &b2)
}
