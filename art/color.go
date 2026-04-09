package art

import (
	"regexp"
	"strings"
)

var template = make([]string, 0, 855)

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

func Color(input, substr, color, banner string) string {
	var result strings.Builder

	if input == "" {
		result.WriteRune('\n')
		return result.String()
	}

	r := []rune(input)
	template = GenerateTemplate(banner)
	var drawn = make([][]string, len(input))

	// get locations of substr matches in input
	re := regexp.MustCompile(substr)
	matches := re.FindAllStringIndex(input, -1)

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

		if colorFlag && i < stopColor {
			for _, line := range uncolored {
				colored = append(colored, COLORS[strings.ToLower(color)]+line+Reset)
			}
		}

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

	template = template[:0] // reset cached template
	return result.String()
}

/*
var Colors = map[string]string{
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

func Color(input, substr, color, banner string) {
	template := GenerateTemplate(banner)

	if input == "" {
		os.Exit(0)
	}

	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	userInput := strings.SplitSeq(input, "\\n")

	for input := range userInput {
		var matches [][]int

		// find <substring> to color
		if substr == input { // if no <substring> given, use input itself as pattern
			sub := regexp.MustCompile(input)
			matches = sub.FindAllStringIndex(input, -1)
		} else { // if <substring> given, check for its existence in input
			sub := regexp.MustCompile(substr)
			matches = sub.FindAllStringIndex(input, -1)
		}

		// draw normally if no matches found
		if len(matches) == 0 {
			Draw(input, banner)
			continue
		}

		if input == "" {
			fmt.Println()
		} else {
			// Reject non-ASCII characters
			for _, char := range input {
				if char < 32 || char > 126 {
					fmt.Printf("Error: ASCII character '%c' is not supported\n", char)
					return
				}
			}
			// Print ASCII art row by row
			for row := range 8 {
				count := 0
				var colorSwitch bool
				for idx, char := range input {
					// Map char to banner lines
					indexCalc := (int(char)-32)*9 + 1 + row
					if idx == matches[count][0] {
						colorSwitch = true
					}

					if idx == matches[count][1] { // end of match index
						colorSwitch = false
						count++
					}

					if count == len(matches) { // keep count within bounds
						count = 0
					}

					if colorSwitch {
						if Colors[color] == "" { // print without color if color not in map
							color = "default"
						}
						fmt.Print(Colors[color] + template[indexCalc] + Reset)
					} else {
						fmt.Print(template[indexCalc])
					}
				}
				fmt.Println() // Next row
			}
		}
	}
}
*/