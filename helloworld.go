package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(printSpaceSeparatedCommandLineArgument(os.Args))
	fmt.Println(printCommandLineArgumentFromAIndex(os.Args, 3))
}

func printSpaceSeparatedCommandLineArgument (args []string) string {
	var res, sep string
	for i := 1; i < len(os.Args); i++ {
		res += sep + os.Args[i]
		sep = " "
	}

	return res
}

func printCommandLineArgumentFromAIndex (args []string, index int) string {
	var res, sep string
	for _, args := range os.Args[index:] {
		res += sep + args
		sep = " "
	}
	return res
}

