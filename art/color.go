package art

import (
	"fmt"
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
		// color everything
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

func Color(line []Word, inputLines []string, substr, color string) []Word {
	substr = strings.ReplaceAll(substr, "\\n", "")
	colorCode := COLORS[strings.ToLower(color)]
	result := make([]Word, len(line))

	for i := range len(line) {
		fmt.Printf("%v: %v\n", i, line[i])
	}
	fmt.Println()

	fmt.Printf("len(line): %v\n", len(line))

	for i, word := range line {
		fmt.Printf("current i: %v\n", i)
		if len(word) == 0 {
			result[i] = word
			continue
		}

		// // since the line itself 
		// cleanedInputLine := strings.ReplaceAll(inputLines[i], " ", "")
		// colored := coloredIndices(cleanedInputLine, substr)
		colored := coloredIndices(inputLines[i], substr)
		fmt.Printf("inputLine: %v\n", inputLines)
		newWord := make(Word, len(word))

		fmt.Printf("colored map:%v\n", colored)
		
		for j, char := range word {
			fmt.Printf("colored char?:%v\n", colored[j])
			if colored[j] {
				inputLines := make([]string, 8)
				for k, line := range char {
					inputLines[k] = colorCode + line + Reset
				}
				newWord[j] = inputLines
			} else {
				newWord[j] = char
			}
		}

		fmt.Printf("newWord: %v\n", newWord)

		result[i] = newWord
	}

	return result
}
