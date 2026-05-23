package art

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// // define function type that takes a line and width then returns the padded string
// type AlignFunc func(line string, terminalWidth int) string

// var AlignFuncs = map[string]AlignFunc{
//     "center": centerLine,
//     "right":  rightLine,
//     "left":   leftLine,
// 	// "justify": justifyLine,
// }

func centerLine(line string, terminalWidth int) string {
	width := (terminalWidth + len(line)) / 2
	return fmt.Sprintf("%*s", width, line)
}

func rightLine(line string, terminalWidth int) string {
	return fmt.Sprintf("%*s", terminalWidth, line)
}

func leftLine(line string, terminalWidth int) string {
	return fmt.Sprintf("%*s", -terminalWidth, line)
}

func justifyLine(words []Word, terminalWidth int) []string {
	var totalContentWidth, totalSpace, numGaps int
	var content []Word

	for i, word := range words {
		if i % 2 == 0 {
			totalContentWidth += word.Width()
			content = append(content, word)
		}

		if i != len(words) - 1 {
			numGaps++
		}
	}

	totalSpace = terminalWidth - totalContentWidth

	baseGap := totalSpace / numGaps
	remainder := totalSpace % numGaps

	// PROBABLE LOGICAL ERROR: will print gap after final word on line too... reconsider whole approach
	var texts [][]string

	for _, w := range content {
		texts = append(texts, w.Lines())
	}
	
	var justified = make([]string, 0, 8)

	for i := 0; i < len(texts); i++ {
		for j := 0; j < len(texts); j++ {
			justified[i] += texts[j][i] 
			if j != 0 && i != len(texts) - 1 {
				justified[i] += strings.Repeat(" ", baseGap)
			// maybe one more condition for if i == len(texts) - 1 && len(texts) == 1??
			} else {
				justified[i] += strings.Repeat(" ", baseGap + remainder)
			}
		}
	}
	return justified

	// var result string
	// for _, line := range justified {
	// 	result += line + "\n"
	// } 

	// return result
}

func getTerminalWidth() int {
	cmd := exec.Command("sh", "-c", "tput cols 2>/dev/tty")
	cols, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not determine terminal width using tput, defaulting to 80")
		cols = []byte("80")
	}

	colsStr := strings.TrimSpace(string(cols))
	width, err := strconv.Atoi(colsStr)
	if err != nil {
		width = 80
	}

	return width
}

func Align(words []Word, alignment string) []string {
	var result []string
	terminalWidth := getTerminalWidth()
	printable := Display(words)
	
	for _, line := range printable {
		switch alignment {
		case "center":
			result = append(result, centerLine(line, terminalWidth))
		case "right":
			result = append(result, rightLine(line, terminalWidth))
		case "left":
			result = append(result, leftLine(line, terminalWidth))
		case "justify":
			result = justifyLine(words, terminalWidth)
		}
	}
	// for _, line := range printable {
	// 	if alignment != "justify" {
	// 		result = append(result, AlignFuncs[alignment](line, terminalWidth))
	// 	} else {
	// 		result = justifyLine(words, terminalWidth)
	// 	}
	// }

	return result
}
