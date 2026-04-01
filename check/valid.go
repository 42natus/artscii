package check

import (
	"regexp"
	"strings"
)

func IsValidColorFlag(args []string) int {
	re := regexp.MustCompile(`--color=[a-zA-Z]+`)

	if strings.HasPrefix(args[0], "--") && re.MatchString(args[0]) {
		return 1
	} else if strings.HasPrefix(args[0], "--") && !re.MatchString(args[0]) {
		return -1
	}
	return 0
}

func IsValidAlignFlag(arg string) int {
	// re := regexp.MustCompile(`--align=^(center|left|right|justify)$`)
	re := regexp.MustCompile(`--align=[a-zA-Z]+`)

	if strings.HasPrefix(arg, "--") && re.MatchString(arg) {
		return 1
	} else if strings.HasPrefix(arg, "--") && !re.MatchString(arg) {
		return -1
	}
	return 0
}