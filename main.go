package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/art"
)

type Options struct {
	color, output, align string
}

const ColorUsageMsg = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
const OutputUsageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
const AlignUsageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"
const CustomUsageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Invalid number of arguments")
		return
	}

	var options Options

	flag.StringVar(&options.color, "color", "", "Color an optional substring in input")
	flag.StringVar(&options.output, "output", "", "Send output to a .txt file")
	flag.StringVar(&options.align, "align", "default", "Align output in terminal")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(CustomUsageMsg)
		return
	}

	// extract positional arguments
	var input, substr string
	banner := "standard"

	if options.color != "" { // active color flag
		if len(args) == 1 {
			substr = ""
			input = args[0]
		} else {
			substr = args[0]
			input = args[1]
			if len(args) == 3 {
				banner = args[2]
			}
		}	
	} else { // unused color flag
		input = args[0]
		if len(args) == 2 {
			banner = args[1]
		}
	}

	// render ASCII art
	lines := strings.Split(input, "\\n")
	words := art.Draw(input, banner)

	// output ASCII art to file
	if options.output != "" {
		var allLines []string
		for _, word := range words {
			allLines = append(allLines, word.Lines()...)
		}
		art.Output(allLines, options.output)
	}

	// color ASCII art
	if options.color != "" {
		words = art.Color(words, lines, substr, options.color)
	}

	// align ASCII art
	if options.align != "default" {
		alignment, ok := art.AlignFuncs[options.align]
		if !ok {
			fmt.Println("Unknown alignment")
			return
		}
		result := art.Align(words, alignment)

		fmt.Print(strings.Join(result, "\n"))
	}

	// output words to terminal
	if strings.ReplaceAll(input, "\\n", "") == "" { // handle input with just '\n's
		count := len(input) / 2
		fmt.Print(strings.Repeat("\n", count))	
	} else {
		result := art.Display(words)
		fmt.Println(strings.Join(result, "\n"))
	}
}
