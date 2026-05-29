package art

import (
	"os/exec"
	"strconv"
	"strings"
)

// extractContentWords is the helper function you proposed.
// It loops through all characters across all elements, groups letters into 
// actual content words, and uses space characters as boundaries.

// func isSpaceWord(w Word) bool {
// 	lines := w.Lines()
// 	if len(lines) == 0 {
// 		return true
// 	}
// 	// If the first row contains only empty spaces, this whole Word is a spacing element
// 	return strings.TrimSpace(stripANSI(lines[0])) == ""
// }

func extractContentWords(words []Word) []Word {
	var cleanWords []Word
	var currentWord Word

	for _, word := range words {
		for _, charMatrix := range word {
			// Check if this individual character is a space block.
			// charMatrix is a []string of 8 rows representing a single character.
			if len(charMatrix) > 0 && strings.TrimSpace(stripANSI(charMatrix[2])) == "" {
				// We hit a space character! This is our word boundary.
				// If we have been building a word, save it and reset.
				if len(currentWord) > 0 {
					cleanWords = append(cleanWords, currentWord)
					currentWord = nil
				}
			} else {
				// It's a visible letter! Add this character to our current word.
				currentWord = append(currentWord, charMatrix)
			}
		}
	}

	// Don't forget to grab the very last word after the loop finishes
	if len(currentWord) > 0 {
		cleanWords = append(cleanWords, currentWord)
	}

	return cleanWords
}

func justifyLine(words []Word, terminalWidth int) []string {
	// 1. Use your helper function to get a pristine slice of content words
	cleanWords := extractContentWords(words)
	println(len(cleanWords))

	// Fallback: If there's 1 or 0 words, justification is mathematically impossible.
	// We default to a standard left-aligned render.
	if len(cleanWords) <= 1 {
		return Display(words)
	}

	// 2. Calculate the combined horizontal width of all real words
	totalWordWidth := 0
	for _, w := range cleanWords {
		totalWordWidth += w.Width()
	}

	// 3. Determine how many spaces we need to distribute across the gaps
	totalSpacesNeeded := terminalWidth - totalWordWidth
	numGaps := len(cleanWords) - 1

	// If the text is wider than the terminal screen, fallback to default display
	if totalSpacesNeeded <= 0 {
		return Display(words)
	}

	baseGap := totalSpacesNeeded / numGaps
	remainder := totalSpacesNeeded % numGaps

	// 4. Reconstruct the 8 rows by stitching our clean words and new spaces
	justifiedLines := make([]string, 8)
	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for i, w := range cleanWords {
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

func leftLine(line string, terminalWidth int) string {
	return line
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
	default: // "left" / "default"
		return Display(words)
	}
}