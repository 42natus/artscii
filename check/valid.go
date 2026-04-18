package check

import (
	"regexp"
	"strings"
)

var validFlags = []*regexp.Regexp{
	regexp.MustCompile(`^--color=[a-zA-Z]+$`),
	regexp.MustCompile(`^--output=[a-zA-Z]+\.txt$`),
	regexp.MustCompile(`^--align=(left|center|right|justify)$`),
}

func HasInvalidFlags(args []string) bool {
	for _, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			continue
		}
		recognised := false
		for _, re := range validFlags {
			if re.MatchString(arg) {
				recognised = true
				break
			}
		}
		if !recognised {
			return true
		}
	}
	return false
}
