package check

import (
	"regexp"
	"strings"
)

var ColorRe = regexp.MustCompile(`--color=[a-zA-Z]+`)
var OutputRe = regexp.MustCompile(`--output=[a-zA-Z]+.txt`)
var AlignRe = regexp.MustCompile(`--align=[a-zA-Z]+`)

func HasValidColorOption(args []string) int {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !ColorRe.MatchString(arg) {
			if !OutputRe.MatchString(arg) && !AlignRe.MatchString(arg) {
				return -1
			}
		}
	}
	return 1
}

func HasValidOutputOption(args []string) int {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !OutputRe.MatchString(arg) {
			if !ColorRe.MatchString(arg) && !AlignRe.MatchString(arg) {
				return -1
			}
		}
	}
	return 1
}

func HasValidAlignOption(args []string) int {
	// re := regexp.MustCompile(`--align=^(center|left|right|justify)$`)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && !AlignRe.MatchString(arg) {
			if !ColorRe.MatchString(arg) && !OutputRe.MatchString(arg) {
				return -1
			}
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
