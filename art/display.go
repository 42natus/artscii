package art

// Display returns a printable string slice of multiple Words
func Display(words []Word) []string {
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			result = append(result, "")
			continue
		}
		result = append(result, word.Lines()...)
	}
	return result
}
