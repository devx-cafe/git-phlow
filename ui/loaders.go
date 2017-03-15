package ui

import (
	"time"
	"github.com/briandowns/spinner"
)

//PhlowSpinner ...
var PhlowSpinner Loader

type Loader struct {
	spn *spinner.Spinner
}

func init() {
	PhlowSpinner = Loader{spn: spinner.New(spinner.CharSets[9], 60*time.Millisecond)}
}

func (l *Loader) Start(suffix string) {
	l.spn.Suffix = " " + suffix
	l.spn.Start()
}

func (l *Loader) Stop() {
	l.spn.Stop()
}
