package art

import (
	"os/exec"
	"strconv"
	"strings"
)

func isSpaceWord(lines []string) bool {
	for _, line := range lines {
		if strings.TrimSpace(stripANSI(line)) != "" {
			return false
		}
	}
	return true
}

// isolate actual content words (i.e. not whitespace) from the words set
func extractContentWords(words []Word) []Word {
	var actualWords []Word
	var currentWord Word

	for _, word := range words {
		// check if each character is a space block.
		for _, charMatrix := range word {
			if len(charMatrix) > 0 && isSpaceWord(charMatrix) { // space words act as word boundaries.
				if len(currentWord) > 0 {
					actualWords = append(actualWords, currentWord) // put "string" of char matrices in slice of actual words
					currentWord = nil
				}
			} else {
				// current character is a letter. add it to the current word.
				currentWord = append(currentWord, charMatrix)
			}
		}
	}

	// grab the very last word after the loop ends
	if len(currentWord) > 0 {
		actualWords = append(actualWords, currentWord)
	}

	return actualWords
}

func justifyLine(words []Word, terminalWidth int) []string {
	// get slice of actual content words
	actualWords := extractContentWords(words)

	// default to left alignment if there's 0 or 1 word.
	if len(actualWords) <= 1 {
		return Display(words)
	}

	// calculate the combined width of all real words
	totalWordWidth := 0
	for _, w := range actualWords {
		totalWordWidth += w.Width()
	}

	// determine how many spaces we need to distribute across the gaps
	totalSpacesNeeded := terminalWidth - totalWordWidth
	numGaps := len(actualWords) - 1

	// If the text is wider than the terminal screen, fallback to default display
	if totalSpacesNeeded <= 0 {
		return Display(words)
	}

	baseGap := totalSpacesNeeded / numGaps
	remainder := totalSpacesNeeded % numGaps

	// 4. Reconstruct the 8 rows by stitching our clean words and new spaces
	justifiedLines := make([]string, 8)
	for row := range 8 {
		var sb strings.Builder
		for i, w := range actualWords {
			wordLines := w.Lines()
			sb.WriteString(wordLines[row])

			// Inject the calculated terminal padding between words (but not after the last word)
			if i < numGaps {
				currentGap := baseGap
				if i < remainder {
					currentGap++ // Smoothly distribute leftover remainder spaces left-to-right
				}
				sb.WriteString(strings.Repeat(" ", currentGap))
			}
		}
		justifiedLines[row] = sb.String()
	}

	return justifiedLines
}

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
