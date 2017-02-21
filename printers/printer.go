package printers

import (
	"fmt"
	"os"
)

//PrintVerbose ...
func PrintVerbose(message string, verbose bool) {
	if verbose {
		fmt.Fprintln(os.Stdout, message)
	}
}
