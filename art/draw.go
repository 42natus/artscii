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
		var line []Word
		inputWords := strings.Split(field, " ")
		for i, w := range inputWords {
			line = append(line, render(w, template))
			if i < len(inputWords) - 1 {
				line = append(line, render(" ", template))
			}
		}
		words = append(words, line)
	}

	return words
}

// generates the ASCII art for each line its called on
func render(s string, template []string) Word {
	r := []rune(s)
	word := make(Word, len(r))

	for i, ch := range r {
		start := (ch-' ')*9 + 1
		word[i] = template[start : start+8]
	}
	return word
}
