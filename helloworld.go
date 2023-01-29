package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(printSpaceSeparatedCommandLineArgument(os.Args))
}

func printSpaceSeparatedCommandLineArgument (args []string) string {
	var res, sep string
	for i := 0; i < len(os.Args); i++ {
		res += sep + os.Args[i]
		sep = " "
	}

	return res
}

