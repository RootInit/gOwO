//CLI owofier.
//Copyright 2021, RootInit

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	gowo "github.com/RootInit/gOwO"
)

func main() {
	args := parseArguments()
	// Set input source
	var input io.Reader
	input = os.Stdin
	if args.inFile != nil {
		input = args.inFile
	} else if args.text != "" {
		input = strings.NewReader(args.text)
	}
	// Set output source
	var output = os.Stdout
	if args.outFile != nil {
		output = args.outFile
	}
	// Initialize Owofier
	owo := gowo.DefaultOwofier()
	// Setup stats map if needed
	if args.statsEnabled {
		stats := owo.Stats(input)
		fmt.Print("Statistics:\n")
		for key, count := range stats {
			// Match  -  replacement    used Num times
			fmt.Printf("\t%3s: %4s - %10d\n",
				key, owo.Replacements[key], count)
		}
	} else {
		err := owo.Translate(input, output)
		if err != nil {
			os.Exit(1)
		}
	}
}

type cliOptions struct {
	statsEnabled bool
	inFile       *os.File
	outFile      *os.File
	text         string
}

func parseArguments() cliOptions {
	// Remove executable path argument
	args := os.Args[1:]
	params := cliOptions{}
	var err error
	var i = 0
loop:
	for ; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			// Help
			showHelp("")
			os.Exit(0)
		case "-s", "--stats":
			// Enable statitistics mode
			params.statsEnabled = true
		case "-i", "--infile":
			// Parse infile argument
			fileArg := args[i+1]
			if len(args) < i+1 || strings.HasPrefix(fileArg, "-") {
				showHelp("Expected argument for infile.")
				os.Exit(1)
			}
			// Open file
			params.inFile, err = os.Open(fileArg)
			if err != nil {
				fmt.Printf(
					"Unable to open infile \"%s\" error: %v", fileArg, err)
				os.Exit(1)
			}
			i += 1 // Already used next argument
		case "-o", "--outfile":
			// Parse outfile argument
			fileArg := args[i+1]
			if len(args) < i+1 || strings.HasPrefix(fileArg, "-") {
				showHelp("Expected argument for outfile.")
				os.Exit(1)
			}
			// Create or open file
			params.outFile, err = os.Create(fileArg)
			if err != nil {
				fmt.Printf(
					"Unable to create outfile \"%s\" error: %v",
					fileArg, err)
				os.Exit(1)
			}
			i += 1 // Already used next argument
		default:
			break loop // End of flags
		}
	}
	params.text = strings.Join(args[i:], " ")
	return params
}

func showHelp(errMsg string) {
	if errMsg != "" {
		fmt.Println("Error: ", errMsg)
	}
	fmt.Println(
		"Usage: owo [OPTION]... [TEXT|FILE]...\n\n" +

			"Owoify text passed as an argument.\n" +
			"\tExample: owo \"text to owoify\"\n\n" +

			"Or as part of a pipeline operation:\n" +
			"\tExample: owo --help | owo\n\n" +

			"Options:\n" +
			"\t-h, --help\t\tshow this message\n" +
			"\t-s, --stats\t\tdry run showing replacement statistics\n" +
			"\t-i, --infile [PATH]\tload text from file\n" +
			"\t-o, --outfile [PATH]\toutput to file" +
			"If no text is provided as an argument read from standard input.\n" +
			"Exit status is 0 if no error occurs, 1 otherwise.")
}
