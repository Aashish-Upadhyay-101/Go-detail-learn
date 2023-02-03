package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(printSpaceSeparatedCommandLineArgument(os.Args))
	// fmt.Println(printCommandLineArgumentFromAIndex(os.Args, 2))
	// printCountAndTextFromFile()
	countFromStandardInput()
}

func printSpaceSeparatedCommandLineArgument(args []string) string {
	var res string = strings.Join(os.Args, " ")
	return res
}

func printCommandLineArgumentFromAIndex(args []string, index int) string {
	var res, sep string
	for _, args := range os.Args[index:] {
		res += sep + args
		sep = " "
	}
	return res
}

/*
Dup1 prints the text of each line that appears more than
once in the standard input, preceded by its count.
*/
func countFromStandardInput() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++

		if input.Text() == "" {
			break
		}
	}
	for line, n := range count {
		fmt.Println(line, n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
	prints the count and text of lines that appear more than once

in the input.  It reads from stdin or from a list of named files.
*/
func printCountAndTextFromFile() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lines)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {

}
