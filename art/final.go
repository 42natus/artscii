package art

import "strings"

func FinalOutput(outputRows, inputLines []string) string {
	var sb strings.Builder

	for _, field := range inputLines {
		if field == "" {
			sb.WriteString("\n")
			continue
		}

		for i := range 8 {
			sb.WriteString(outputRows[i])
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
		outputRows = outputRows[8:]
	}

	final := sb.String()

	return final[:len(final)-1]
}
