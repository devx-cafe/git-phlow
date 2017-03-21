package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

//PhlowSpinner ...
var PhlowSpinner Loader

//Loader ...
//loader struct for long running tasks
type Loader struct {
	spn *spinner.Spinner
}

func init() {
	PhlowSpinner = Loader{spn: spinner.New(spinner.CharSets[9], 60*time.Millisecond)}
}

//Start ...
//Start the spinner with a prefix
func (l *Loader) Start(suffix string) {
	l.spn.Suffix = " " + suffix
	l.spn.Start()
}

//Stop ...
//Stops the current spinner
func (l *Loader) Stop() {
	l.spn.Stop()
}
