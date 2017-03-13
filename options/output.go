package options

import "fmt"

//Bold ...
func Bold(message string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", message)
}
