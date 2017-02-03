package subprocess

import (
	"os/exec"
	"bytes"
	"errors"
)

const (
	EmptyReturnString string = ""
)


//SimpleExec
//needs to also take care of stdin
//needs to take care og std error
func SimpleExec(name string, args ...string) (string, error) {

	if err := IsInPath(name); err != nil {
		return EmptyReturnString, err
	}

	cmd := exec.Command(name, args...)

	var outBuffer, errBuffer bytes.Buffer

	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer

	if err := cmd.Start(); err != nil {
		return EmptyReturnString, err
	}

	if err := cmd.Wait(); err != nil {
		return EmptyReturnString, err
	}

	

	return outBuffer.String(), nil
}

/*
func SimpleExec(name string, args ...string) (error, string) {

	if _, e := Discover(name); e != nil {
		return e, ""
	}

	command := exec.Command(name, args...)
	var buffer = &bytes.Buffer{}

	_, err := command.StdoutPipe()
	command.Stdout = buffer

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}


		fmt.Printf("==> Output: %s\n", string(buffer))

	/*
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = command.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
	}

	err = command.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err, )
	}

	return nil, buffer.String()
}

*/

func IsInPath(application string) (error) {
	_, err := exec.LookPath(application)

	if err != nil {
		return errors.New(application + " is not in $PATH")

	}
	return nil
}