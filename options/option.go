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

	//GlobalFlagVersion ...
	GlobalFlagVersion = false

	//Sha1 git commit hash
	Sha1 string

	//Version build version
	Version string

	//Date date of build
	Date string
)
