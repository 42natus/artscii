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

func coloredIndices(inputLine, substr string) map[int]bool {
	colored := make(map[int]bool)
	if substr == "" {
		for i := range len(inputLine) {
			colored[i] = true
		}
		return colored
	}

	re := regexp.MustCompile(regexp.QuoteMeta(substr))
	for _, match := range re.FindAllStringIndex(inputLine, -1) {
		for i := match[0]; i < match[1]; i++ {
			colored[i] = true
		}
	}
	return colored
}

func Color(line []Word, currentInputLine string, substr, color string) []Word {
	colorCode, exists := COLORS[strings.ToLower(color)]
	if !exists {
		colorCode = COLORS["default"]
	}

	coloredMap := coloredIndices(currentInputLine, substr)
	result := make([]Word, len(line))
	globalCharIdx := 0 // Tracks our exact underlying location inside currentInputLine

	for i, word := range line {
		newWord := make(Word, len(word))
		for j, charMatrix := range word {
			if coloredMap[globalCharIdx] {
				coloredChar := make([]string, 8)
				for k, row := range charMatrix {
					coloredChar[k] = colorCode + row + Reset
				}
				newWord[j] = coloredChar
			} else {
				newWord[j] = charMatrix
			}
			globalCharIdx++
		}
		result[i] = newWord
	}
	return result
}
