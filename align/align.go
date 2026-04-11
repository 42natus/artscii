package align

import (
	"fmt"
	"strings"

	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/art"
	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/check"
)

func Center(input, banner string) {
	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	template := art.GenerateTemplate(banner)
	if template == nil {
		return
	}

	terminalWidth := check.GetTerminalWidth()

	for word := range strings.SplitSeq(input, "\\n") {
		if word == "" {
			fmt.Println()
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
			var line strings.Builder
			for j := range n {
				line.WriteString(drawn[j][i])
			}
			width := (terminalWidth + len(line.String())) / 2
			fmt.Printf("%*v\n", width, line.String())
		}
	}
}

func Right(input, banner string) {
	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	template := art.GenerateTemplate(banner)
	if template == nil {
		return
	}

	terminalWidth := check.GetTerminalWidth()

	for word := range strings.SplitSeq(input, "\\n") {
		if word == "" {
			fmt.Println()
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
			var line strings.Builder
			for j := range n {
				line.WriteString(drawn[j][i])
			}
			fmt.Printf("%*v\n", terminalWidth, line.String())
		}
	}
}

func Justify(input, banner string) {
	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))
		return
	}

	template := art.GenerateTemplate(banner)
	if template == nil {
		return
	}

	terminalWidth := GetTerminalWidth()

	for word := range strings.SplitSeq(input, "\\n") {
		if word == "" {
			fmt.Println()
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
			var line strings.Builder
			for j := range n {
				line.WriteString(drawn[j][i])
			}
			width := (terminalWidth + len(line.String())) / 2
			fmt.Printf("%*v\n", width, line.String())
		}
	}
}
