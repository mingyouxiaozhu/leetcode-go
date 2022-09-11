package offer2

import "strings"

func replaceSpace(str string) string {
	return strings.ReplaceAll(str, " ", "%20")
}
