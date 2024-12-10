package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args

	readFromStdin := false
	var options []string
	var fileName string
	if len(args) == 1 {
		readFromStdin = true
		options = []string{"-c", "-l", "-w"}
	} else if len(args) == 2 && args[1] == "--help" || args[1] == "-h" {
		printHelp()
		return
	} else if len(args) == 2 && strings.HasSuffix(args[1], ".txt") {
		fileName = args[1]
		options = []string{"-c", "-l", "-w"}
	} else if len(args) > 2 && strings.HasSuffix(args[len(args)-1], ".txt") {
		options = args[1 : len(args)-1]
		fileName = args[len(args)-1]
	} else {
		return
	}

	var content string
	if readFromStdin {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error trying to read from standard input: ", err)
			return
		}
		content = string(input)
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("error: cannot find file %s\n", fileName)
			return
		}
		defer file.Close()

		fileContent, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("error: cannot read file")
			return
		}
		content = string(fileContent)
	}

	numLines := len(strings.Split(content, "\n"))
	numWords := len(strings.Fields(content))
	numBytes := len(content)
	numChars := len([]rune(content))

	for _, option := range options {
		switch option {
		case "-c":
			fmt.Printf("%4d ", numBytes)
		case "-l":
			fmt.Printf("%4d ", numLines)
		case "-w":
			fmt.Printf("%4d ", numWords)
		case "-m":
			fmt.Printf("%4d ", numChars)
		default:
			fmt.Printf("error: cannot find option: %s\n find more with --help, or -h", option)
			return
		}
	}

	if !readFromStdin {
		fmt.Printf("%s\n", fileName)
	} else {
		fmt.Println()
	}
}

func printHelp() {
	fmt.Println(`ccwc - A simple command-line word count tool

Usage: 
  ccwc [options] <filename>
  ccwc [options] - (reads from standard input)

Command options:
  -c          print the number of bytes in the file;
  -l          print the number of lines in the file;
  -w          print the number of words in the file;
  -m          print the number of characters in the file;
  -h, --help  show the help;

Examples:
  ccwc -l -w test.txt
    Prints the number of lines and words in 'test.txt'

  cat test.txt | ccwc -c
    Reads from standard input and prints the byte count`)
}
