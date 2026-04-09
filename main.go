package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/art"
	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/check"
	"acad.learn2earn.ng/git/foloruns/ascii-art-justify/align"
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

	// validate options, if present
	validColor, validOutput, validAlign := check.HasValidOptions(os.Args[1:])

	// if validColor == 
	if validColor == -1 {
		fmt.Println(ColorUsageMsg)
		return
	}

	if validOutput == -1 {
		fmt.Println(OutputUsageMsg)
		return 
	}

	if validAlign == -1 {
		fmt.Println(AlignUsageMsg)
		return
	}

	var options Options

	flag.StringVar(&options.color, "color", "default", "Color an optional substring in input")
	flag.StringVar(&options.output, "output", "nil", "Send output to a .txt file")
	flag.StringVar(&options.align, "align", "default", "Align output in terminal")

	flag.Parse()

	args := flag.Args()
	var input, substr, banner string

	switch len(args) {
	case 1:
		substr = args[0]
		input = args[0]
		banner = "standard"
	case 2:
		substr = args[0]
		input = args[0]
		banner = args[1]
	case 3:
		if options.color == "default" && options.output != "nil" {
			fmt.Println(OutputUsageMsg)
			return
		} else if options.color == "default" && options.align != "default" {
			fmt.Println(AlignUsageMsg)
			return
		}
		substr = args[0]
		input = args[1]
		banner = args[2]
	default:
		fmt.Println(CustomUsageMsg)
		return
	}

	var output string

	if options.color != "default" {
		output = art.Color(input, substr, options.color, banner)
	} else {
		output = art.Draw(input, banner)
	}

	if options.output != "nil" {
		art.Output(options.output, input, banner)
	}

	if options.align != "default" {
		// art.Align(alignType, input, banner) --> is this cleaner?
		switch strings.ToLower(options.align) {
		case "center":
			align.Center(input, banner)
		case "right":
			align.Right(input, banner)
		case "justify":
			align.Justify(input, banner)
		default:
			fmt.Print(art.Draw(input, banner))
		}
	} else {
		fmt.Print(output)
	}
}
