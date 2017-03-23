package phlow

import "strings"


func isBehindOrAhead(str string) (behind, ahead bool) {
	ahead = strings.Contains(str, "ahead")
	behind = strings.Contains(str, "behind")
	return
}
