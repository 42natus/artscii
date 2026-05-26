package art

import (
	"strings"
)

func Draw(input, banner string) [][]Word {
	template := GenerateTemplate(banner)
	if template == nil {
		return nil
	}
	
	var lines [][]Word
	for field := range strings.SplitSeq(input, "\\n") {
		var line []Word
		runes := []rune(field)
		for _, ch := range runes {
			line = append(line, render(string(ch), template))
		}
		lines = append(lines, line)
	}

	return lines
}

func render(s string, template []string) Word {
	if len(s) == 0 {
		return nil
	}
	ch := rune(s[0])
	word := make(Word, 1)
	start := int(ch-' ')*9 + 1
	word[0] = template[start : start+8]
	return word
}