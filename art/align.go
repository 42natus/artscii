package art

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// define function type that takes a line and width then returns the padded string
type AlignFunc func(line string, terminalWidth int) string

func centerLine(line string, terminalWidth int) string {
	width := (terminalWidth + len(line)) / 2
	return fmt.Sprintf("%*s", width, line)
}

func rightLine(line string, terminalWidth int) string {
	return fmt.Sprintf("%*s", terminalWidth, line)
}

func leftLine(line string, terminalWidth int) string {
	return fmt.Sprintf("%*s", -terminalWidth, line)
}

func getTerminalWidth() int {
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

func Align(words []Word, alignment AlignFunc) []string {
	// for _, word := range words {
	// 	word.Width()
	// }
}
