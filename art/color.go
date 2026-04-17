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
	colorCode := COLORS[strings.ToLower(color)]
	result := make([]Word, len(words))

	for i, word := range words {
		if len(word) == 0 {
			result[i] = word
			continue
		}

		colored := coloredIndices(substr, lines[i])
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

/*
func Color0(input, substr, color, banner string) string {
	var result strings.Builder

	if strings.ReplaceAll(input, "\\n", "") == "" { // handle input with just '\n's
		count := len(input) / 2
		result.WriteString(strings.Repeat("\n", count))
		return result.String()
	}

	template := GenerateTemplate(banner)
	if template == nil {
		return result.String()
	}

	for word := range strings.SplitSeq(input, "\\n") {
		if word == "" {
			result.WriteRune('\n')
			continue
		}

		n := len(word)
		r := []rune(word)
		drawn := make([][]string, n)

		// ignore "\n" in substring to color
		var pattern string
		if strings.Contains(substr, "\\n") {
			pattern = strings.ReplaceAll(substr, "\\n", "")
		}

		// get locations of substr matches in input
		re := regexp.MustCompile(regexp.QuoteMeta(pattern))
		matches := re.FindAllStringIndex(word, -1)
		fmt.Println(matches)

		var colorFlag bool
		var stopColor int
		for i, ch := range r {
			start := (ch-' ')*9 + 1
			uncolored := template[start : start+8]
			colored := []string{}

			if len(matches) > 0 && i == matches[0][0] {
				colorFlag = true
				stopColor = i + len(substr)
				matches = matches[1:] // remove matches one-by-one after coloring them
			}

			// color whole substring match
			if colorFlag && i < stopColor {
				for _, line := range uncolored {
					colored = append(colored, COLORS[strings.ToLower(color)]+line+Reset)
				}
			}

			// done coloring substring match
			if colorFlag && i == stopColor {
				colorFlag = false
			}

			if len(colored) > 0 {
				drawn[i] = colored
			} else {
				drawn[i] = uncolored
			}
		}

		// build final string
		for i := range 8 {
			for j := range len(drawn) {
				result.WriteString(drawn[j][i])
			}
			result.WriteRune('\n')
		}
	}

	return result.String()
}
*/
