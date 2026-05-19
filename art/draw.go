package art

import (
	"strings"
)

func Draw(input, banner string) []Word {
	template := GenerateTemplate(banner)
	if template == nil {
		return nil
	}

	var words []Word
	for line := range strings.SplitSeq(input, "\\n") {
		words = append(words, render(line, template))
	}

	return words
}

// generates the ASCII art for each line its called on
func render(line string, template []string) Word {
	r := []rune(line)
	word := make(Word, len(r))

	for i, ch := range line {
		start := (ch-' ')*9 + 1
		word[i] = template[start : start+8]
	}
	return word
}
