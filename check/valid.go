package check

import (
	"regexp"
	"strings"
)

func HasValidColorOption(args []string) int {
	re := regexp.MustCompile(`--color=[a-zA-Z]+`)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !re.MatchString(arg) {
			return -1
		}
	}
	return 1
}

func HasValidOutputOption(args []string) int {
	re := regexp.MustCompile(`--output=[a-zA-Z]+.txt`)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !re.MatchString(arg) {
			return -1
		}
	}
	return 1
}

func HasValidAlignOption(args []string) int {
	// re := regexp.MustCompile(`--align=^(center|left|right|justify)$`)
	re := regexp.MustCompile(`--align=[a-zA-Z]+`)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !re.MatchString(arg) {
			return -1
		}
	}
	return 1
}

func HasValidOptions(args []string) (color, output, align int) {
	color = HasValidColorOption(args)
	output = HasValidOutputOption(args)
	align = HasValidAlignOption(args)

	return color, output, align
}
