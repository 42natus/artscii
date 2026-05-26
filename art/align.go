package art

import (
	"os/exec"
	"strconv"
	"strings"
)

func centerLine(line string, terminalWidth int) string {
	plainText := stripANSI(line)
	if len(plainText) >= terminalWidth {
		return line
	}
	spaces := (terminalWidth - len(plainText)) / 2
	return strings.Repeat(" ", spaces) + line
}

func rightLine(line string, terminalWidth int) string {
	plainText := stripANSI(line)
	if len(plainText) >= terminalWidth {
		return line
	}
	spaces := terminalWidth - len(plainText)
	return strings.Repeat(" ", spaces) + line
}

func justifyLine(words []Word, terminalWidth int) []string {
	var cleanWords []Word
	var totalWordWidth int

	// 1. Filter out placeholder space words to find real content words
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		// If a word's plain text conversion contains only spaces, skip it
		if strings.TrimSpace(stripANSI(strings.Join(w.Lines(), ""))) == "" {
			continue
		}
		cleanWords = append(cleanWords, w)
		totalWordWidth += w.Width()
	}

	// Fallback: If there's 1 or 0 words, justification is impossible. Default to Left.
	if len(cleanWords) <= 1 {
		return Display(words)
	}

	// 2. Distribute the padding space mathematically
	totalSpacesNeeded := terminalWidth - totalWordWidth
	numGaps := len(cleanWords) - 1

	if totalSpacesNeeded <= 0 {
		return Display(words) // Content overflows screen boundary; default render
	}

	baseGap := totalSpacesNeeded / numGaps
	remainder := totalSpacesNeeded % numGaps

	// 3. Assemble the 8 horizontal rows matching the gaps
	justifiedLines := make([]string, 8)
	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for i, w := range cleanWords {
			wordLines := w.Lines()
			sb.WriteString(wordLines[row])

			// Add space padding after every word except the last one
			if i < numGaps {
				currentGap := baseGap
				if i < remainder {
					currentGap++ // Distribute remainders smoothly left-to-right
				}
				sb.WriteString(strings.Repeat(" ", currentGap))
			}
		}
		justifiedLines[row] = sb.String()
	}

	return justifiedLines
}

func getTerminalWidth() int {
	cmd := exec.Command("sh", "-c", "tput cols 2>/dev/tty")
	cols, err := cmd.Output()
	if err != nil {
		cols = []byte("80")
	}
	width, err := strconv.Atoi(strings.TrimSpace(string(cols)))
	if err != nil {
		width = 80
	}
	return width
}

func Align(words []Word, alignment string) []string {
	terminalWidth := getTerminalWidth()

	switch alignment {
	case "center":
		var result []string
		for _, line := range Display(words) {
			result = append(result, centerLine(line, terminalWidth))
		}
		return result
	case "right":
		var result []string
		for _, line := range Display(words) {
			result = append(result, rightLine(line, terminalWidth))
		}
		return result
	case "justify":
		return justifyLine(words, terminalWidth)
	default: // "left"
		return Display(words)
	}
}