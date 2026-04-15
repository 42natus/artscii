package art

import "strings"

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
