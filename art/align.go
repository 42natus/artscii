package art

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// define function type that takes a line and width then returns the padded string
type AlignFunc func(line string, terminalWidth int) string

var AlignFuncs = map[string]AlignFunc{
    "center": centerLine,
    "right":  rightLine,
    "left":   leftLine,
	// "justify": justifyLine,
}

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

func justifyLine0(line string, terminalWidth int) string {
	return fmt.Sprintf("%*s", -terminalWidth, line)
}

func justifyLine(words []Word, terminalWidth int) string {
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

	// problem: will print gap after final word on line too... reconsider whole approach
	// Display(content)
	var texts [][]string
	for _, w := range content {
		texts = append(texts, w.Lines())
	}

	var final []string
	for i, text := range texts {
		extra := 0
		if i < remainder {
			extra = 1
		}
		// this gap gets baseGap + extra spaces
		gap := baseGap + extra

		final[i] += text[i] + strings.Repeat(" ", gap)

	}

	// for i := range numGaps {
	// 	extra := 0, guy
	// 	if i < remainder {
	// 		extra = 1
	// 	}
	// 	// this gap gets baseGap + extra spaces
	// 	gap := baseGap + extra
	// }
	// The first remainder gaps each get one extra column.

	return fmt.Sprintf("%*s", -terminalWidth, line)
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

func Align(words []Word, alignment AlignFunc) []string {
	var result []string
	terminalWidth := getTerminalWidth()
	printable := Display(words)
	
	for _, line := range printable {
		if alignment != justifyLine {
			result = append(result, alignment(line, terminalWidth))
		} else {
			result = append(result, alignment(words, ))
		}
	}

	return result
}


/*
1. For each `[]Word` line, separate content words from space words — content at even indices, spaces at odd indices
2. Calculate `totalContentWidth` — sum of `word.Width()` for even-indexed words
3. Calculate `totalSpace = terminalWidth - totalContentWidth`
4. Calculate `numGaps = len(line) / 2` — number of space slots
5. Distribute space using `baseGap` and `remainder`
6. Build each output row by interleaving content lines with gap strings
*/