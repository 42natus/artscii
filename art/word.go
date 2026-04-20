package art

import (
	"regexp"
	"strings"
)

// Word is a rendered ASCII-art word.
// Word[i] is the i-th character; Word[i][j] is the j-th line of that character.
type Word [][]string

// Lines collapses the Word into a slice of printable strings, one per row.
func (w Word) Lines() []string {
	if len(w) == 0 {
		return nil
	}

	lines := make([]string, 8)

	for i := range 8 {
		var line strings.Builder

		for _, char := range w {
			line.WriteString(char[i])
		}
		lines[i] = line.String()
	}
	return lines
}

var ansiRe = regexp.MustCompile(`\033\[[0-9;]*m`)

func stripANSI(line string) string {
	return ansiRe.ReplaceAllString(line, "")
}

func (w Word) Width() int {
	if len(w) == 0 {
		return 0
	}
	// get the entire first line of the ASCII art word
	var sb strings.Builder
	for _, char := range w {
		sb.WriteString(char[0])
	}
	return len(stripANSI(sb.String()))
}
