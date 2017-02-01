package main

import (
	"github.com/praqma/git-phlow/subprocess"
	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	cmd.Execute()
	subprocess.SimpleExec("./test.sh")
}