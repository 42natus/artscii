package main

import (
	"fmt"
	"os"
	"strings"

	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/align"
	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/art"
	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/check"
)

// const ColorUsgMsg = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
const JustifyUsgMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(JustifyUsgMsg)
		return
	}

	args := os.Args[1:]

	switch check.IsValidAlignFlag(args[0]) {
	case -1: // invalid align flag
		fmt.Println(JustifyUsgMsg)
		return
	case 0: // no align flag
		if len(args) == 1 {
			art.Draw(args[0], "standard") // maybe use align.Left()
		} else if len(args) == 2 {
			art.Draw(args[0], args[1]) // maybe use align.Left()
		}
	case 1: // valid align flag
		var input, alignment, banner string
		alignment = args[0][8:]
		values := args[1:]

		input = values[0]
		if len(values) == 1 {
			banner = "standard"
		} else if len(values) == 2 {
			banner = values[1]
		} else {
			fmt.Println(JustifyUsgMsg)
			return
		}

		if strings.ToLower(alignment) == "center" {
			align.Center(input, banner)
		}
	}
}
