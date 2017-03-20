/*Package options ...
Contains all global variables
*/
package options

//Global flags for commandline inputs ...
var (
	//GlobalFlagVerbose ...
	GlobalFlagVerbose = false

	//GlobalFlagLocal ...
	GlobalFlagLocal = false

	//GlobalFlagMine ...
	GlobalFlagMine = false

	//GlobalFlagVersion ...
	GlobalFlagVersion = false

	//GlobalFlagForce ...
	GlobalFlagForce = false

	//GlobalFlagShowTestOutput ...
	GlobalFlagShowTestOutput bool

	//GlobalFlagHumanReadable ...
	GlobalFlagHumanReadable bool

	//GlobalFlagHard ...
	GlobalFlagHard bool

	//Sha1 git commit hash
	Sha1 string

	//Version build version
	Version string

	//Date date of build
	Date string
)
