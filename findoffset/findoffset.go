// findoffset.go is a command line tool that finds the offset of the first occurence of a string in a file and prints it to stdout.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. parse the command line arguments

	// the operating system provides command line arguments to your program.
	// os.Args[0] is the name of the program, and the rest are the 'real' arguments.
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: findoffset <filename> <string>")
		os.Exit(1)
	}

	filepath, pattern := os.Args[1], os.Args[2]

	// 2. read the file into memory
	// we read the whole file into memory for the sake of the demonstration
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v", filepath, err)
		os.Exit(1)
	}

	for i := 0; i < len(b)-len(pattern); i++ {
		for j := range pattern {
			if b[i+j] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				fmt.Fprintf(os.Stdout, "%d\n", i)
				os.Exit(0)
			}
		}
	}
	os.Exit(1)
}
