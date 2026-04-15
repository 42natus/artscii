package art

func Draw(input, banner string) Word {
	template := GenerateTemplate(banner)
	if template == nil {
		return nil
	}

	r := []rune(input)
	word := make(Word, len(r))

	for i, ch := range input {
		start := (ch-' ')*9 + 1
		word[i] = template[start : start+8]
	}

	return word
}

/*
func Draw0(input, banner string) string {
	var result strings.Builder

	if strings.ReplaceAll(input, "\\n", "") == "" { // handle input with just '\n's
		count := len(input) / 2
		result.WriteString(strings.Repeat("\n", count))
		return result.String()
	}

	template := GenerateTemplate(banner)
	if template == nil {
		return result.String()
	}

	for word := range strings.SplitSeq(input, "\\n") {
		if word == "" {
			result.WriteRune('\n')
			continue
		}
		n := len(word)
		r := []rune(word)

		drawn := make([][]string, n)

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
	}

	return result.String()
}
*/
