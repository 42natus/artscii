package art

import (
	"strings"
)

func Draw(input, banner string) [][]Word {
	template := GenerateTemplate(banner)
	if template == nil {
		return nil
	}
	
	var words [][]Word
	
	for field := range strings.SplitSeq(input, "\\n") {
		if strings.Contains(field, " ") { // generate each whitespace-separated word
			var line []Word
			for subWord := range strings.SplitSeq(field, " ") {
				line = append(line, render(subWord, template))
				// line = append(line, render(" ", template))
			}
		} else {
			var line []Word
			line = append(line, render(field, template))
			words = append(words, line)
		}
	}

	return words
}

// generates the ASCII art for each line its called on
func render(s string, template []string) Word {
	r := []rune(s)
	word := make(Word, len(r))

	for i, ch := range s {
		start := (ch-' ')*9 + 1
		word[i] = template[start : start+8]
	}
	return word
}
