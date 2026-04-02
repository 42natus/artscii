package align

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/art"
)

func GetTerminalWidth() int {
	cmd := exec.Command("sh", "-c", "tput cols 2>/dev/tty")
	cols, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not determine terminal width using tput, defaulting to 80")
		cols = []byte("80")
	}

	colsStr := strings.TrimSpace(string(cols))
	width, err := strconv.Atoi(colsStr)
	if err != nil {
		width = 80
	}

	return width
}

func Center(input, banner string) {
	template := art.GenerateTemplate(banner)

	if len(template) == 0 {
		return
	}

	if input == "" {
		os.Exit(0)
	}

	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	userInput := strings.SplitSeq(input, "\\n")
	terminalWidth := GetTerminalWidth()

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
				var line strings.Builder
				for _, char := range word {
					// Map char to banner lines
					indexCalc := (int(char)-32)*9 + 1 + row
					line.WriteString(template[indexCalc])
				}
				// to center, halve the sum of the terminal width and string length
				width := (terminalWidth + len(line.String())) / 2
				fmt.Printf("%*v\n", width, line.String())
			}
		} else {
			fmt.Println()
		}
	}
}
