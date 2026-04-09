package art

import (
	"strings"
)

func Draw(input, banner string) string {
	var result strings.Builder

	if input == "" {
		result.WriteRune('\n')
		return result.String()
	}

	n := len(input)
	r := []rune(input)
	template := GenerateTemplate(banner)

	// if len(template) == 0 {
	// 	return
	// }

	var drawn = make([][]string, n)

	/*
		generate the art for each character in the input.

		the formula uses the character's distance from SPACE in unicode (effectively the ASCII table),
		multiplied by the no. of lines each character takes up*, plus 1 to account for the
		topmost empty line in the banner file.

		(*each character takes up 8 lines + 1 newline after it = 9 lines)
	*/
	for i, ch := range r {
		start := (ch-' ')*9 + 1
		drawn[i] = template[start : start+8]
	}

	// build final string
	for i := range 8 {
		for j := range n {
			result.WriteString(drawn[j][i])
		}
		result.WriteRune('\n')
	}
	template = template[:0]
	return result.String()
}
