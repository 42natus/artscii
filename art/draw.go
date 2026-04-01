package art

import (
	"fmt"
	"os"
	"strings"
)

func Draw(input, banner string) {
	template := GenerateTemplate(banner)

	if input == "" {
		os.Exit(0)
	}

	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	userInput := strings.SplitSeq(input, "\\n")

	for word := range userInput {
		if word != "" {
			// Reject non-ASCII characters
			for _, char := range word {
				if char < 32 || char > 126 {
					fmt.Printf("Error: ASCII character '%c' is not supported\n", char)
					return
				}
			}
			// Print ASCII art row by row
			for row := range 8 {
				for _, char := range word {
					// Map char to banner lines
					indexCalc := (int(char)-32)*9 + 1 + row
					fmt.Print(template[indexCalc])
				}
				fmt.Println() // Next row
			}
		} else {
			fmt.Println()
		}
	}
}
