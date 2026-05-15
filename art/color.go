package art

import (
	"regexp"
	"strings"
)

var COLORS = map[string]string{
	"white":   "\033[38;2;255;255;255m",
	"silver":  "\033[38;2;192;192;192m",
	"gray":    "\033[38;2;128;128;128m",
	"black":   "\033[38;2;0;0;0m",
	"red":     "\033[38;2;255;0;0m",
	"maroon":  "\033[38;2;128;0;0m",
	"orange":  "\033[38;2;255;165;0m",
	"yellow":  "\033[38;2;255;255;0m",
	"olive":   "\033[38;2;128;128;0m",
	"lime":    "\033[38;2;0;255;0m",
	"green":   "\033[38;2;0;128;0m",
	"cyan":    "\033[38;2;0;225;225m",
	"aqua":    "\033[38;2;0;225;225m",
	"teal":    "\033[38;2;0;128;128m",
	"blue":    "\033[38;2;0;0;225m",
	"navy":    "\033[38;2;0;0;128m",
	"magenta": "\033[38;2;225;0;225m",
	"fuchsia": "\033[38;2;255;0;255m",
	"purple":  "\033[38;2;128;0;128m",
	"default": "\033[0;39;0m",
}

const Reset = "\033[0m"

func coloredIndices(line, substr string) map[int]bool {
	colored := make(map[int]bool)
	if substr == "" {
		// color everything
		for i := range len(line) {
			colored[i] = true
		}
		return colored
	}

	re := regexp.MustCompile(regexp.QuoteMeta(substr))
	for _, match := range re.FindAllStringIndex(line, -1) {
		for i := match[0]; i < match[1]; i++ {
			colored[i] = true
		}
	}

	return colored
}

func Color(words []Word, lines []string, substr, color string) []Word {
	substr = strings.ReplaceAll(substr, "\\n", "")
	colorCode := COLORS[strings.ToLower(color)]
	result := make([]Word, len(words))

	for i, word := range words {
		if len(word) == 0 {
			result[i] = word
			continue
		}

		colored := coloredIndices(lines[i], substr)
		newWord := make(Word, len(word))

		for j, char := range word {
			if colored[j] {
				lines := make([]string, 8)
				for k, line := range char {
					lines[k] = colorCode + line + Reset
				}
				newWord[j] = lines
			} else {
				newWord[j] = char
			}
		}

		result[i] = newWord
	}

	return result
}
