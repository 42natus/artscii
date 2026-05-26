package art

import "strings"

// Display returns a horizontally stitched slice of 8 strings for a given line of words
func Display(words []Word) []string {
	if len(words) == 0 {
		return nil
	}

	lines := make([]string, 8)
	// Stitch row 0 of every word, then row 1 of every word, etc.
	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for _, word := range words {
			wordLines := word.Lines()
			if row < len(wordLines) {
				sb.WriteString(wordLines[row])
			}
		}
		lines[row] = sb.String()
	}
	return lines
}