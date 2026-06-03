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

/*
const ColorUsageMsg = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
const OutputUsageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
const AlignUsageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"
*/
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

	var input, substr string
	banner := "standard"

	if options.color != "" {
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
	} else {
		input = args[0]
		if len(args) == 2 {
			banner = args[1]
		}
	}

	if strings.ReplaceAll(input, "\\n", "") == "" {
		fmt.Print(strings.Repeat("\n", len(input)/2))
		return
	}

	// render ASCII art
	wordsMatrix := art.Draw(input, banner)
	
	// color the render
	inputLines := strings.Split(input, "\\n")
	if options.color != "" {
		for i, line := range wordsMatrix {
			wordsMatrix[i] = art.Color(line, inputLines[i], substr, options.color)
		}
	}

	// align the render
	var outputRows []string
	for _, line := range wordsMatrix {
		var renderedLine []string
		if options.align != "default" {
			renderedLine = art.Align(line, options.align)
		} else {
			renderedLine = art.Display(line)
		}
		outputRows = append(outputRows, renderedLine...)
	}

	// send outputRows to file specified
	if options.output != "" {
		art.Output(outputRows, options.output)
	} else {
		fmt.Print(art.FinalOutput(outputRows, inputLines))
		// fmt.Println(strings.Join(outputRows, "\n"))
	}
}
